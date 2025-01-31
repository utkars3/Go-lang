package main

import "fmt"

// Implementing stack
type Stack struct {
	stack []int
}

func (st *Stack) push(val int) {
	st.stack = append(st.stack, val)
}

func (st *Stack) pop() int {
	if len(st.stack) == 0 {
		fmt.Println("Stack is empty")
		return -1
	}

	returnVal := st.stack[len(st.stack)-1]
	st.stack = st.stack[:len(st.stack)-1]
	fmt.Println("Popped element is ", returnVal)
	return returnVal
}

func (st *Stack) top() int {
	if len(st.stack) == 0 {
		fmt.Println("Stack is empty")
		return -1
	}
	returnVal := st.stack[len(st.stack)-1]
	fmt.Println("Top value is ", returnVal)
	return returnVal
}

func (st *Stack) size() int {
	fmt.Println("Size of the stack is ", len(st.stack))
	return len(st.stack)

}

// queue
type Queue struct {
	queue []int
}

func (st *Queue) push(val int) {
	st.queue = append(st.queue, val)
}

func (st *Queue) pop() int {
	if len(st.queue) == 0 {
		fmt.Println("Queue is empty")
		return -1
	}

	returnVal := st.queue[0]
	st.queue = st.queue[1:len(st.queue)]
	fmt.Println("Popped element is ", returnVal)
	return returnVal
}

func (st *Queue) top() int {
	if len(st.queue) == 0 {
		fmt.Println("Queue is empty")
		return -1
	}
	returnVal := st.queue[0]
	fmt.Println("Top value is ", returnVal)
	return returnVal
}

func (st *Queue) size() int {
	fmt.Println("Size of the stack is ", len(st.queue))
	return len(st.queue)

}

func main() {
	// fmt.Println("stack with slices")

	// st := Stack{}
	// st.push(2)
	// st.push(22)
	// st.push(24)
	// st.pop()
	// st.pop()
	// st.pop()
	// st.pop()
	// st.top()
	// st.size()

	//queue
	q := Queue{}

	q.push(25)
	q.push(23)
	q.size()
	q.top()
	q.pop()
	q.top()

}
