GO语言没有Class，只有Struct

结构体有自己的数据结构以及调用这些数据结构的方法，可以实现类所具有的功能

struct 类型的定义 方法比较特殊
格式为. 参数可以为值，也可以为指针
func (varName StructType) funcName {

}

要改变值必须用指针
结果过大也考虑用指针
一致性：如果有指针接受者，最好都是指针接收者

封装：

名字一般用CamelCase, 驼峰式

包名：
    首字母大写：代表public
    首字母小写：代表private
包的定义：
    每个目录只能有一个包
    main包包含可执行入口
    为结构定义的方法放在同一个保内
    可以是不同文件
    
    
名词注释：
package 
   包，与文件夹名字一致
   
struct
     结构体（相当于类）， 放置在每个package文件夹内任意文件内
     设定语法
     Type 类名 struct {//字段列表}
     
struct func
     结构体的方法
     1. 特有的工厂方法
        func Create类名(参数列表){//初始化语句}    
     2. 其他应用型方法
        a. 读类型
        func (s 结构体类名) 方法名 {//具体实现} 
        b. 写类型（通常都是写类型）
        func (s *结构体类名) 方法名 {//具体实现} 
     3. 方法的调用
         包名.方法名(结构体变量名)   
        

组装方式扩展原有的结构体方法
创建新的结构体，内部包含一个需要扩展的结构体变量指针
例如：
type myTreeNode struct {
	node *tree.Node
}

扩展之前的方法：
/**
 扩展了之前遍历方法
 左，右，中 顺序遍历树, 后续遍历
*/
func (myNode *myTreeNode) postOrder(){
	if myNode == nil || myNode.node== nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()

}
         



    
    
        
    



