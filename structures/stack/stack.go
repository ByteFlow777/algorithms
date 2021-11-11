package stack

type Stack struct {
	items []interface{}
}

// NewStack return a new queue
func NewStack() *Stack {
	return &Stack{
		items: []interface{}{},
	}
}

func (q *Stack) Push(item interface{}) {
	q.items = append(q.items, item)
}

func (q *Stack) Pop() (interface{}, bool) {
	n := q.Len()
	if n == 0 {
		return nil, false
	}
	ans := q.items[n-1]
	q.items = q.items[0 : n-1]

	return ans, true
}

func (q Stack) Len() int {
	return len(q.items)
}
