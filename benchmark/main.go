package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"

	client "github.com/chenminhua/darkbase/client"
)
const (
	tablename = "benchmark"
)
var server, operation string
var total, valueSize, threads, pipelen int

func init() {
	flag.StringVar(&server, "h", "localhost:12345", "cache server address")
	flag.IntVar(&total, "n", 1000, "total number of requests")
	flag.IntVar(&valueSize, "d", 1000, "data size of SET/GET value in bytes")
	flag.IntVar(&threads, "c", 1, "number of parallel connections")
	flag.StringVar(&operation, "t", "set", "test set, could be get/set/mixed")
	flag.IntVar(&pipelen, "P", 1, "pipeline length")
	flag.Parse()
	fmt.Println("server is", server)
	fmt.Println("total", total, "requests")
	fmt.Println("data size is", valueSize)
	fmt.Println("we have", threads, "connections")
	fmt.Println("operation is", operation)
	fmt.Println("pipeline len is", pipelen)
	rand.Seed(time.Now().UnixNano())
}

func pipeline(clt *client.DarkbaseClient) {
}

func operate(tid, count int, ch chan *result) {
	clt, err := client.NewClient(server)
	if clt == nil || err != nil {
		panic("clt is nil")
	}

	valuePrefix := strings.Repeat("a", valueSize)
	r := &result{0, 0, 0, make([]statistic, 0)}
	for i := 0; i < count; i++ {
		tmp := tid*count + i
		key := fmt.Sprintf("%d", tmp)
		value := []byte(fmt.Sprintf("%s%d", valuePrefix, tmp))
		start := time.Now()
		resultType := operation
		if operation == "set" {
			err := clt.Put(tablename, key, value)
			if err != nil {
				resultType = "err"
			}
		}
		if operation == "get" {
			r, err := clt.Get(tablename, key)
			if err != nil {
				resultType = "err"
			}
			if len(r) == 0 {
				resultType = "miss"
			}
		}
		d := time.Now().Sub(start)
		r.addDuration(d, resultType)
	}

	ch <- r
}

func main() {
	ch := make(chan *result, threads)
	res := &result{0, 0, 0, make([]statistic, 0)}
	start := time.Now()
	for i := 0; i < threads; i++ {
		go operate(i, total/threads, ch)
	}

	for i := 0; i < threads; i++ {
		res.addResult(<-ch)
	}

	d := time.Now().Sub(start)
	totalCount := res.getCount + res.missCount + res.setCount
	fmt.Printf("%d records get\n", res.getCount)
	fmt.Printf("%d records miss\n", res.missCount)
	fmt.Printf("%d records set\n", res.setCount)
	fmt.Printf("%f seconds total\n", d.Seconds())
	statCountSum := 0
	statTimeSum := time.Duration(0)
	for b, s := range res.statBuckets {
		if s.count == 0 {
			continue
		}
		statCountSum += s.count
		statTimeSum += s.time
		fmt.Printf("%d%% requests < %d ms\n", statCountSum*100/totalCount, b+1)
	}
	fmt.Printf("%d usec average for each request\n", int64(statTimeSum/time.Microsecond)/int64(statCountSum))
	fmt.Printf("throughput is %f MB/s\n", float64((res.getCount+res.setCount)*valueSize)/1e6/d.Seconds())
	fmt.Printf("rps is %f\n", float64(totalCount)/float64(d.Seconds()))


}
