package stl

type Queue []interface{}

func (q *Queue) Push(x interface{}) {
	*q = append(*q, x)
}

func (q *Queue) Pop() interface{} {
	h := *q
	var el interface{}
	l := len(h)
	el, *q = h[0], h[1:l]
	// Or use this instead for a Stack
	// el, *q = h[l-1], h[0:l-1]
	return el
}

func NewQueue() *Queue {
	return &Queue{}
}
