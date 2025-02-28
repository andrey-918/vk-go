package stack

type Stack struct {
	items []interface{}
}

func New() *Stack {
	return &Stack{
		items: []interface{}{},
	}
}

// Push добавляет элемент на верх стека
func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

// Pop удаляет и возвращает верхний элемент стека
func (s *Stack) Pop() (interface{}, bool) {
	if len(s.items) == 0 {
		return nil, false // Стек пуст
	}
	topIndex := len(s.items) - 1
	item := s.items[topIndex]
	s.items = s.items[:topIndex] // Удаляем верхний элемент
	return item, true
}

// Top возвращает верхний элемент стека без его удаления
func (s *Stack) Top() (interface{}, bool) {
	if len(s.items) == 0 {
		return nil, false // Стек пуст
	}
	return s.items[len(s.items)-1], true
}

// IsEmpty проверяет, пуст ли стек
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Iterate() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		for i := 0; i <= len(s.items)-1; i++ { // Итерируемся от верхушки стека
			ch <- s.items[i]
		}
		close(ch)
	}()
	return ch
}
