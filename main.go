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
	
	
	array := lang.NewArrayList()
	array.Add("hello")
	array.Add("world")
	array.Add("Hicham")
	fmt.Println("total items in array:", array.Count())
	fmt.Println("array:", array)
	fmt.Println("array[1]:", array.Get(1))
	fmt.Println("array.indexOf(world):", array.IndexOf("world"))
	fmt.Println("array.Contains(world):", array.Contains("world"), array.Contains("random"))
	
	for i := 0; i < array.Count(); i++ {
		fmt.Println("item(", i, "):", array.Get(i))
	}
	
	fmt.Println("array.Remove(world):", array.Remove("world"), array)
	fmt.Println("array.Remove(Hicham):", array.Remove("Hicham"), array)
	fmt.Println("array.IsEmpty():", array.IsEmpty(), array.Remove("hello"), array.IsEmpty(), array.Count())
	
	for i := 0; i < 30; i++ {
		array.Add(i)
	}
	
	fmt.Println("array.indexOf(5):", array.IndexOf(5), array.Count())
	array.Clear()
	array.Add("one", "two", "three")
	
	array2 := lang.NewArrayList()
	array2.Add("four", "five", "six")
	
	array.AddFromArrayList(array2)
	
	fmt.Println("array:", array, array.Count())
	fmt.Println("array.First/Last:", array.First(), array.Last())
	fmt.Println("array.ToSlice:", array.ToSlice())
	
	
	set := lang.NewHashSet()
	set.Add("hello", "world", "hello")
	fmt.Println("total items in set:", set.Count(), set)
	setToSlice := set.ToSlice()
	fmt.Println("set.ToSlice:", setToSlice, len(setToSlice))
	
}