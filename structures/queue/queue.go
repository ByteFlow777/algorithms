package queue

type Queue struct {
	items []interface{}
}

// NewQueue return a new queue
func NewQueue() *Queue {
	return &Queue{
		items: []interface{}{},
	}
}

func (q *Queue) Push(item interface{}) {
	q.items = append(q.items, item)
}

func (q *Queue) Pop() (interface{}, bool) {
	n := q.Len()
	if n == 0 {
		return nil, false
	}
	ans := q.items[0]
	q.items = q.items[1:n]

	return ans, true
}

func (q Queue) Len() int {
	return len(q.items)
}
