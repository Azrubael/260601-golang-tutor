package main
import "fmt"

type Stack struct{
	items []int
}

type Queue struct{
	items []int
}

// Push - will add an element at the end
func (s *Stack) Push(i int){
	s.items = append(s.items, i)
}

// Pop - will remove an element from the end
func (s *Stack) Pop() int{
	l := len(s.items) -1
	toRemove := s.items[l]
	s.items = s.items[:l]
}

// Enqueue - will add an element at the end
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Dequeue - - will remove an element from the end
func (q *Stack) Dequeue() int{
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.push(11)
	myStack.push(22)
	myStack.push(33)
	fmt.Println(myStack)
	
}

