package main

import "fmt"
func main(){
	// list :=10
	// something :=[]int{12312,123123123,012321312,-13123,1,3,1242,51251,912312,512,12312}
	// list= append(list,10)
	// fmt.Print("i'm here",list)

	// fmt.Print(bubbleSort(something))
	hello()
}
func bubbleSort(list []int) ([]int){
	for i := 0; i < len(list); i++ {
		for j:=i+1; j<len(list);j++{
			if list[i]>list[j]{
				tmp :=list[i]
				list[i]=list[j]
				list[j]=tmp
			}
		}
	}
	return list
}
func hello(){
	fmt.Println("Hey there, I like you can you please tell me your name?")
	var name string
	fmt.Scanln(&name)
	fmt.Println(name)
}
