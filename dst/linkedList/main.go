package main

import "fmt"

type node struct{
	value int
	next *node
}
type slist struct{
	size int
	head *node
}

func createList() slist{
	list:=slist{size:0}
	return list
}
func (sl *slist) insertAtHead(val int){
	newNode:=node{value:val}
	if sl.size==0{
		sl.head=&newNode
	}else{
		newNode.next=sl.head
		sl.head=&newNode
	}
	sl.size++
}
func (sl *slist) removeAtHead(){
	if(sl.size<=1){
		sl.size=0
		sl.head=nil
	}else{
		sl.size--
		sl.head=sl.head.next
	}
}
func (sl slist) lastNode()*node{
	var currNode *node
	if(sl.size==1){
		return sl.head
	}
	for i:=0;i<sl.size;i++{
		currentNode:=sl.head
		if(sl.size>1){
			sl.head=currentNode.next
		}
		currNode=currentNode
	}
	return currNode
}
func (sl *slist) insertAtEnd(val int){
	newNode:=node{value:val}
	sl.lastNode().next=&newNode
	sl.size++
}
func (sl *slist) removeAtEnd(){

	
}
func (sl *slist) reverse(){

}
func (sl slist)print(){
	for i:=0;i<sl.size;i++{
		currentNode:=sl.head
		fmt.Print(currentNode.value," ")
		if(sl.size>1){
			sl.head=currentNode.next
		}
	}
}
func main(){
	list:=createList()
	list.insertAtHead(5)
	list.insertAtHead(6)
	list.insertAtEnd(9)
	list.print()
	list.removeAtHead()
	fmt.Println()
	list.print()
	fmt.Println()
	fmt.Print(list.lastNode())
}
