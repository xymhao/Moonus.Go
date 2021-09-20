package structure

type Queue struct {
	value []interface{}
	Len   int
}

type Queer interface {
	Enqueue(value interface{})
	Dequeue() interface{}
}

func (q *Queue) Enqueue(value interface{}) {
	q.value = append(q.value, value)
	q.Len++
}

func (q *Queue) Dequeue() interface{} {
	result := q.value[0]
	q.value = q.value[1:]
	q.Len--
	return result
}
