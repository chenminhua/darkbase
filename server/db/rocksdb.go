package db

// #include <stdlib.h>
// #include "rocksdb/c.h"
// #cgo LDFLAGS: -lrocksdb -lz -lpthread -lsnappy -lstdc++ -lm -O3
import "C"
import (
	"errors"
	"github.com/chenminhua/darkbase/server/conf"
	"runtime"
	"time"
	"unsafe"
)

type rocksDB struct {
	db *C.rocksdb_t
	ro *C.rocksdb_readoptions_t
	wo *C.rocksdb_writeoptions_t
	e  *C.char
	ch chan *pair
}

type pair struct {
	k string
	v []byte
}

const BATCH_SIZE = 100

func flush_batch(db *C.rocksdb_t, b *C.rocksdb_writebatch_t, o *C.rocksdb_writeoptions_t) {
	var e *C.char
	C.rocksdb_write(db, o, b, &e)
	if e != nil {
		panic(C.GoString(e))
	}
	C.rocksdb_writebatch_clear(b)
}

func write_func(db *C.rocksdb_t, c chan *pair, o *C.rocksdb_writeoptions_t) {
	count := 0
	t := time.NewTimer(time.Second)
	b := C.rocksdb_writebatch_create()
	for {
		select {
		case p := <-c:
			count++
			key := C.CString(p.k)
			value := C.CBytes(p.v)
			C.rocksdb_writebatch_put(b, key, C.size_t(len(p.k)), (*C.char)(value), C.size_t(len(p.v)))
			C.free(unsafe.Pointer(key))
			C.free(value)
			if count == BATCH_SIZE {
				flush_batch(db, b, o)
				count = 0
			}
			if !t.Stop() {
				<-t.C
			}
			t.Reset(time.Second)
		case <-t.C:
			if count != 0 {
				flush_batch(db, b, o)
				count = 0
			}
			t.Reset(time.Second)
		}
	}
}

func OpenRocksdb(tableName string) *rocksDB {
	if db, exists := rdbMap[tableName]; exists {
		return db;
	}
	options := C.rocksdb_options_create()
	C.rocksdb_options_increase_parallelism(options, C.int(runtime.NumCPU()))
	C.rocksdb_options_set_create_if_missing(options, 1)
	var e *C.char
	db := C.rocksdb_open(options, C.CString("/Users/admin/tmp/rocksdb/" + tableName), &e)
	if e != nil {
		panic(C.GoString(e))
	}
	C.rocksdb_options_destroy(options)
	var c chan *pair = nil
	wo := C.rocksdb_writeoptions_create()
	if (conf.GetBatchConfig()) {
		c = make(chan *pair, 5000)
		go write_func(db, c, wo)
	}
	rdb := &rocksDB{db, C.rocksdb_readoptions_create(), wo, e, c}
	rdbMap[tableName] = rdb
	return rdb;
}

func (db *rocksDB) Put(key string, value []byte) error {
	if (conf.GetBatchConfig()) {
		db.ch <-  &pair{key, value}
		return nil
	}
	k := C.CString(key)
	defer C.free(unsafe.Pointer(k))
	v := C.CBytes(value)
	defer C.free(v)
	C.rocksdb_put(db.db, db.wo, k, C.size_t(len(key)), (*C.char)(v), C.size_t(len(value)), &db.e)
	if db.e != nil {
		return errors.New(C.GoString(db.e))
	}
	return nil
}


func (db *rocksDB) Get(key string) ([]byte, error) {
	k := C.CString(key)
	defer C.free(unsafe.Pointer(k))
	var length C.size_t
	v := C.rocksdb_get(db.db, db.ro, k, C.size_t(len(key)), &length, &db.e)
	if db.e != nil {
		return nil, errors.New(C.GoString(db.e))
	}
	defer C.free(unsafe.Pointer(v))
	return C.GoBytes(unsafe.Pointer(v), C.int(length)), nil
}


func (db *rocksDB) Del(key string) error {
	k := C.CString(key)
	defer C.free(unsafe.Pointer(k))
	C.rocksdb_delete(db.db, db.wo, k, C.size_t(len(key)), &db.e)
	if db.e != nil {
		return errors.New(C.GoString(db.e))
	}
	return nil
}

