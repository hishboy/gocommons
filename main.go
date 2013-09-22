package main

import "fmt"
import "github.com/hishboy/gocommons/lang"

func main() { 
	queue := lang.NewQueue()
	queue.Push("Hello")
	queue.Push(4)
	queue.Push(5)
	queue.Push(8)
	queue.Push(6)
	queue.Push("World")
	fmt.Println("Total items before poll: ", queue.Count())
	fmt.Println(queue.Poll(), queue.Poll(), queue.Poll(), queue.Poll(), queue.Poll(), queue.Poll(), queue.Poll())
	fmt.Println("Total items after poll: ", queue.Count())
	
	fmt.Println("\n")
	
	stack := lang.NewStack()
	stack.Push("World")
	stack.Push(4)
	stack.Push(5)
	stack.Push(8)
	stack.Push(6)
	stack.Push("Hello")
	fmt.Println("Total items before pop: ", stack.Count())
	fmt.Println(stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop())
	fmt.Println("Total items after pop: ", stack.Count())
}