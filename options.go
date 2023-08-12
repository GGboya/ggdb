package ggdb

type Options struct {
	// 数据库数据目录
	DirPath string

	// 数据文件的大小
	DataFileSize int64

	// 每次写数据是否持久化
	SyncWrites bool
}
