package data

import "encoding/binary"

type LogRecordType = byte

const (
	LogRecordNormal LogRecordType = iota
	LogRecordDeleted
)

// crc type keySize valueSize
// 4 + 1     + 5     + 5 = 15
const maxLogRecordHeaderSize = binary.MaxVarintLen32*2 + 5

// LogRecord 写入到数据文件的记录
// 之所以叫日志，是因为数据文件中的数据是追加写入的，类似日志的格式
type LogRecord struct {
	Key   []byte
	Value []byte
	Type  LogRecordType
}

// LogRecord 头部信息
type LogRecordHeader struct {
	crc        uint32        // 校验值
	recordType LogRecordType // 标识类型
	keySize    uint32        // key 长度
	valueSize  uint32        // value 长度
}

// LogRecordPos 数据内存索引，主要是描述数据在磁盘上的位置
type LogRecordPos struct {
	Fid    uint32 // 文件 id, 表示将数据存储到了哪个文件当中
	Offset int64  // 偏移，表示将数据存储到了数据文件中的哪个位置
}

// EncodeLogRecord 对LogRecord进行编码，返回字节数组及长度
func EncodeLogRecord(record *LogRecord) ([]byte, int64) {
	return nil, 0
}

// 对字节数组中的Header信息进行解码
func decodeLogRecordHeader(buf []byte) (*LogRecordHeader, int64) {
	return nil, 0
}

func getLogRecordCRC(lr *LogRecord, header []byte) uint32 {
	return 0
}
