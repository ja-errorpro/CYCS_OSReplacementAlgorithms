package lib

type Queue struct {
	list []interface{}
}

func NewQueue() *Queue {
	q := new(Queue)
	q.list = make([]interface{}, 0)
	return q
}

func (q *Queue) Push(v interface{}) {
	q.list = append(q.list, v)
}

func (q *Queue) Pop() interface{} {
	if len(q.list) == 0 {
		return nil
	}
	v := q.list[0]
	q.list = q.list[1:]
	return v
}

func (q *Queue) Len() int {
	return len(q.list)
}

func (q *Queue) Empty() bool {
	return len(q.list) == 0
}

func (q *Queue) Front() interface{} {
	if len(q.list) == 0 {
		return nil
	}
	return q.list[0]
}
