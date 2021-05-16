package main

import (
	"fmt"
)


func main() {
	fmt.Println("Push de 2 elementos")
	newStack := NewStack()
	newStack.Push(90322)
	newStack.Push(903)
	newStack.print()	
	fmt.Println("Pop de 1 elemento")
	newStack.Pop()
	newStack.print()

	fmt.Println("\nPush de 2 elementos en cola")
	newQueue := NewQueue()
	newQueue.Push(90322)
	newQueue.Push(903)
	newQueue.print()
	fmt.Println("Pop de 1 elemento")
	newQueue.Pop()
	newQueue.print()

	fmt.Println("\nProblema 3")
	fmt.Println("a ()( le faltan: ", minimalParentheses("()("))
	fmt.Println("a ((()( le faltan: ", minimalParentheses("((()("))
	fmt.Println("a ()((( le faltan: ", minimalParentheses("()((("))
	fmt.Println("a ()))(( le faltan: ", minimalParentheses("()))(("))
}

func minimalParentheses(parentheses string) int {
	count := 0
	stack := NewStack()
	for _, r := range parentheses {
		if r == ')'{
			if stack.IsEmpty() {
				count++
			}else{
				stack.Pop()
			}
		}else{
			stack.Push(1)
		}
	}
	return count + stack.Elements
}

//Stack implementation

type Node struct{
	Value int
	Next *Node
}

func NewNode(item int) *Node{
	newNode := Node{}
	newNode.Value = item
	newNode.Next = nil
	return &newNode
}

type Stack struct{
	Head *Node
	Min int
	Elements int
}

func NewStack() *Stack{
	newStack := Stack{}
	newStack.Head = nil
	newStack.Min = 100000000000
	newStack.Elements = 0
	return &newStack
}

func (stack *Stack) IsEmpty() bool {
	if stack.Elements == 0 {
		return true
	}
	return false
}

func (stack *Stack) GetMin() int {
	return stack.Min
}

func (stack *Stack) Push(item int) {
	newNode := NewNode(item)
	if !stack.IsEmpty(){
		newNode.Next = stack.Head
	}
	stack.Head = newNode
	stack.Elements++
	if stack.Min > item{
		stack.Min = item
	}
}

func (stack *Stack) Peek() int{
	return stack.Head.Value
}

func (stack *Stack) Pop() {
	if stack.Head.Value == stack.Min {
		stack.Min = 1000000000
		stack.FindMin()
	}
	stack.Head = stack.Head.Next
	stack.Elements--
}

func (stack *Stack) FindMin() {
	node := stack.Head
	for node != nil {
		if node.Value < stack.Min {
			stack.Min = node.Value
		}
		node = node.Next
	}
}

func (stack *Stack) print() {
	node := stack.Head
	for node != nil {
		fmt.Printf("Elemento: %d\n", node.Value)
		node = node.Next
	}
}

type Queue struct {
	stackPush *Stack
	stackPop *Stack
}

//queue implementation

func NewQueue() *Queue {
	queue := Queue{}
	queue.stackPush = NewStack()
	queue.stackPop = NewStack()
	return &queue
}

func (queue *Queue) IsEmpty() bool {
	if (queue.stackPush.IsEmpty()) && (queue.stackPop.IsEmpty()) {
		return true
	}
	return false
}

func (queue *Queue) GetMin() int {
	if queue.stackPush.IsEmpty() {
		return queue.stackPop.GetMin()
	}
	return queue.stackPop.GetMin()
}

func (queue *Queue) Push(item int) {
	if !queue.stackPop.IsEmpty() {
		exchangeStacks(queue.stackPop, queue.stackPush)
	}
	queue.stackPush.Push(item)
}

func (queue *Queue) Peek() int {
	if queue.stackPop.IsEmpty() {
		exchangeStacks(queue.stackPush, queue.stackPop)
	}
	return queue.stackPop.Peek()
}

func (queue *Queue) Pop() {
	if queue.stackPop.IsEmpty() {
		exchangeStacks(queue.stackPush, queue.stackPop)
	}
	queue.stackPop.Pop()
}

func exchangeStacks(fullStack, emptyStack *Stack) {
	node := fullStack.Head
	for node != nil {
		emptyStack.Push(fullStack.Peek())
		fullStack.Pop()
		node = node.Next
	}
}

func (queue *Queue) print() {
	if queue.stackPop.IsEmpty() {
		exchangeStacks(queue.stackPush, queue.stackPop)
	}
	node := queue.stackPop.Head
	for node != nil {
		fmt.Printf("Elemento: %d\n", node.Value)
		node = node.Next
	}
}

