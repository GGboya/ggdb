# GG_DB
GG_boy的数据库项目



如何构建的整个数据库架构？

我们想一下平时工作中是如何使用mysql的

1、初始化一个mysql实例，用户可以填写信息。比如db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname")

2、连接mysql

3、进行更新，查询等基本操作

4、关闭服务器



那对应的，我们这个基于bitcask模型的kv数据库，使用流程应该也是如此。

1、用户传入一个配置项，我们返回给用户一个数据库实例

2、用户通过这个数据库实例，进行更新，查询等操作



那我们现在挨个的实现上面的流程试一试

**1、用户传入一个配置项，我们返回给用户一个数据库实例**

用户需要配置什么呢？暂时没想到，没关系，先写个空的结构体

数据库实例应该包含哪些信息呢？

按照bitcask模型，我们的数据库实例应当包含一个活跃的日志文件，其他的日志文件都应当是old文件。activeFile可以用来读写，olderFiles只能用来读。在这个基础上，还需要记录用户的配置项。

Ok, 我们先写出这样的结构体

```go
type DB struct {
	options    Options
	activeFile *data.DataFile            // 当前活跃数据文件，可以用于写入
	olderFiles map[uint32]*data.DataFile // 旧的数据文件，只能用于读
}
```



用户需要传入一个配置项，自然我们需要一个函数去接受这个参数。然后返回给用户一个结构体实例对吧。

我们也定义一个Open方法，来返回

```go
func Open(options Options) (*DB, error) {
    // 初始化DB实例结构体
	db := &DB{
		options:    options,
        activeFile: nil, 
		olderFiles: nil,
	}

	return db, nil
}
```

好啦，我们就完成第一个需求了，返回一个DB实例



**2、用户通过这个数据库实例，进行更新，查询等操作**

用户的查询操作：

用户传入一个参数key，我们返回给他一个value。

bitcask模型里面写入文件的数据格式是固定的

![image-20231122200150791](C:\Users\宋宇航\AppData\Roaming\Typora\typora-user-images\image-20231122200150791.png)

依次是数据校验，时间戳，key大小，value大小，key，value

bitcask的读取数据也很简单：

- 根据key找到内存中对应的记录，这个记录存储数据在磁盘的具体位置
- 根据位置找到磁盘的日志文件，再通过具体偏移找到key对应实际value。



知道了用户的key要经历的数据流转，我们就根据key的流动，把代码给写出来。

完全依照bitcask模型，key找到内存中对应的记录，我们用B+树来实现（论文用的哈希表）。用B+树可以支持数据有序遍历。

```go
package ggdb
import "github.com/google/btree"

// BTree 主要封装了 google 的 btree 库
type BTree struct {
	tree *btree.BTree  // 基于B树的索引
}

type Item struct {
	key []byte
	pos *LogRecordPos
}

// LogRecordPos 数据内存索引，主要是描述数据在磁盘上的位置
type LogRecordPos struct {
	Fid    uint32 // 文件 id, 表示将数据存储到了哪个文件当中
	Offset int64  // 偏移，表示将数据存储到了数据文件中的哪个位置
}

func (bt *BTree) Get(key []byte) *data.LogRecordPos {
	it := &Item{key: key}
	btreeItem := bt.tree.Get(it)
	if btreeItem == nil {
		return nil
	}
	return btreeItem.(*Item).pos
}
```

使用谷歌开源的B+树，我们先自定义B+树的结点结构。这里我们只定义了file_id和value_pos

然后B+树就找到这个结点，然后返回key对应的日志文件ID和偏移量

现在的问题就是，我应该怎么根据这个日志文件ID和偏移量，去把实际的value取出来呢？

没有磁盘的路径啊，存不了一点，这个存放路径就需要我们默认，或者交给用户配置。