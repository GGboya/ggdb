package index

import (
	"ggdb/data"
	"testing"
)

func TestBTree_PutAndGet(t *testing.T) {
	bt := NewBTree() // 创建一个新的BTree实例

	// 测试数据
	key1 := []byte("key1")
	pos1 := &data.LogRecordPos{Fid: 1, Offset: 1234}
	key2 := []byte("key2")
	pos2 := &data.LogRecordPos{Fid: 2, Offset: 5678}

	// 测试Put方法
	if !bt.Put(key1, pos1) {
		t.Error("将项目放入BTree失败")
	}
	if !bt.Put(key2, pos2) {
		t.Error("将项目放入BTree失败")
	}

	// 测试Get方法
	resultPos := bt.Get(key1)
	if resultPos == nil || *resultPos != *pos1 {
		t.Errorf("预期位置：%v，但得到：%v", pos1, resultPos)
	}

	resultPos = bt.Get(key2)
	if resultPos == nil || *resultPos != *pos2 {
		t.Errorf("预期位置：%v，但得到：%v", pos2, resultPos)
	}
}

func TestBTree_PutAndDelete(t *testing.T) {
	bt := NewBTree() // 创建一个新的BTree实例

	// 测试数据
	key1 := []byte("key1")
	pos1 := &data.LogRecordPos{Fid: 1, Offset: 1234}

	// 测试Put方法
	if !bt.Put(key1, pos1) {
		t.Error("将项目放入BTree失败")
	}

	// 测试Delete方法
	if !bt.Delete(key1) {
		t.Error("从BTree删除项目失败")
	}

	// 确保项目已被删除
	resultPos := bt.Get(key1)
	if resultPos != nil {
		t.Errorf("预期删除键的位置为nil，但得到：%v", resultPos)
	}
}

func TestBTree_DeleteNonExistentKey(t *testing.T) {
	bt := NewBTree() // 创建一个新的BTree实例

	// 测试数据
	key1 := []byte("key1")

	// 使用不存在的键测试Delete方法
	if bt.Delete(key1) {
		t.Error("从BTree删除不存在的键")
	}
}

func TestBTree_OverwriteKey(t *testing.T) {
	bt := NewBTree() // 创建一个新的BTree实例

	// 测试数据
	key1 := []byte("key1")
	pos1 := &data.LogRecordPos{Fid: 1, Offset: 1234}
	pos2 := &data.LogRecordPos{Fid: 2, Offset: 5678}

	// 测试Put方法
	if !bt.Put(key1, pos1) {
		t.Error("将项目放入BTree失败")
	}

	// 再次用不同的位置插入相同的键
	if !bt.Put(key1, pos2) {
		t.Error("将项目放入BTree失败")
	}

	// 获取键的值应该是最后一次插入的值
	resultPos := bt.Get(key1)
	if resultPos == nil || *resultPos != *pos2 {
		t.Errorf("预期位置：%v，但得到：%v", pos2, resultPos)
	}
}
