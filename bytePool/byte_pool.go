package bytecache

type BytePool struct {
	c    chan []byte
	w    int
	wcap int
}

func NewBytePool(poolSize, size, cap int) *BytePool {
	return &BytePool{
		c:    make(chan []byte, poolSize),
		w:    size,
		wcap: cap,
	}
}

func (bp *BytePool) Get() (b []byte) {
	select {
	case b = <-bp.c:
	default:
		if bp.wcap > 0 {
			b = make([]byte, bp.w, bp.wcap)
		} else {
			b = make([]byte, bp.w)
		}
	}

	return
}

func (bp *BytePool) Put(b []byte) {
	select {
	case bp.c <- b:
	default:
	}
}
