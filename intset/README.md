# IntSet简介

## 前言
这是[@zhangyunhao](!https://github.com/zhangyunhao116) 大佬在分享goland并发结构时留下的homework。
目标是创建一个并发安全的有序链表，其中数据严格有序且不重复。

## IntSet结构
IntSet是一个并发读写安全的链表。
节点结构体中增加一个锁来实现并发读写，其中节点的读操作通过atomic原子操作实现，节点的写操作通过Lock+atomic实现。
```
type IntSet struct {
	head   *intNode
	length int64
}

type intNode struct {
	value int64
	next  *intNode
	isDel uint32
	mu    sync.Mutex
}
```

## 实现接口
IntSet实现的接口如下。
```
// 检查一个元素是否存在，如果存在则返回 true，否则返回 false
Contains(value int) bool

// 插入一个元素，如果此操作成功插入一个元素，则返回 true，否则返回 false
Insert(value int) bool

// 删除一个元素，如果此操作成功删除一个元素，则返回 true，否则返回 false
Delete(value int) bool

// 遍历此有序链表的所有元素，如果 f 返回 false，则停止遍历
Range(f func(value int) bool)

// 返回有序链表的元素个数
Len() int
```

## 并发测试结果
```bash
go test
go test -race  //-race选项用于检测数据竞争
```
在同一机器上结果对比如下。

| | intset耗时| 单线程有序链表耗时|
|---|:---:|:---:|
|go test| 0.270s|0.739s|
|go test -race|19.458s |9.182s|

## 参考资料
单线程的有序链表实现：https://gist.github.com/zhangyunhao116/833c3113db343a660a2adb1e4c21951d
测试代码：https://gist.github.com/zhangyunhao116/833c3113db343a660a2adb1e4c21951d
内网文档：https://bytedance.feishu.cn/docs/doccnWKnOgzM6fijdkwgaypSgjf
