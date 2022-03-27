package pool

type pool[T any] struct {
	noCopy struct{}
	c      chan T
	new    func() T
}

func NewPool[T any](poolSize int) *pool[T] {
	return &pool[T]{
		c: make(chan T, poolSize),
	}
}

func NewPoolWithNew[T any](poolSize int, f func() T) *pool[T] {
	return &pool[T]{
		c:   make(chan T, poolSize),
		new: f,
	}
}

func (p *pool[T]) Get() (o T) {
	select {
	case o = <-p.c:
	default:
		if p.new != nil {
			o = p.new()
		} else {
			o = *new(T)
		}
	}

	return
}

func (p *pool[T]) Put(o T) {
	select {
	case p.c <- o:
	default:
	}
}
