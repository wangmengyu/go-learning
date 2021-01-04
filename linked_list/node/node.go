package node

/**
  链表结构:
  优点: 不需要处理固定容量问题
  确定: 不能随机访问某个节点
*/
type Node struct {
	e    Elem
	next *Node // 最后一个节点的next 是nil 值
}

type Elem struct {
	val int
}
