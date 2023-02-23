package queue

type Queue struct {
	items    []int
	capacity int
}

func New(capacity int) Queue {
	return Queue{
		items:    make([]int, 0, capacity),
		capacity: capacity,
	}
}

func (q *Queue) Append(item int) {
	q.items = append(q.items, item)
	if len(q.items) > q.capacity {
		q.items = q.items[1:]
	}
}
