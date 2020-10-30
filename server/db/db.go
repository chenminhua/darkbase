package db

var	rdbMap map[string]*rocksDB

func init() {
	rdbMap = make(map[string]*rocksDB)
}

type DB interface {
	Put(string, []byte) error
	Get(string) ([]byte, error)
	Del(string) error
}
