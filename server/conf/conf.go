package conf

var batch = false

func SetBatchConfig(b bool) {
	batch = b
}

func GetBatchConfig() bool {
	return batch
}


