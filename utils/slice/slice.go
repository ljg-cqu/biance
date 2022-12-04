package slice

type Slice struct {
	capacity int
	store    []any
}

func New(capacity int) *Slice {
	return &Slice{
		capacity: capacity,
		store:    make([]any, 0, capacity),
	}
}

func (s *Slice) AddElem(a any) {
	s.store = append(s.store, a)
	if len(s.store) > s.capacity {
		s.store = s.store[1:]
	}
}

func (s *Slice) Elem(index int) any {
	if index >= len(s.store) {
		panic("index overflow")
	}
	return s.store[index]
}

func (s *Slice) Len() int {
	return len(s.store)
}
