package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := New()

	// Проверка, что стек пуст
	if !stack.IsEmpty() {
		t.Errorf("Expected stack to be empty")
	}

	// Тестируем Push
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// Проверка, что стек не пуст
	if stack.IsEmpty() {
		t.Errorf("Expected stack to not be empty")
	}

	// Тестируем Top
	top, ok := stack.Top()
	if !ok || top != 3 {
		t.Errorf("Expected top element to be 3, got %v", top)
	}

	// Тестируем Pop
	item, ok := stack.Pop()
	if !ok || item != 3 {
		t.Errorf("Expected popped element to be 3, got %v", item)
	}

	// Проверяем новый верхний элемент
	top, ok = stack.Top()
	if !ok || top != 2 {
		t.Errorf("Expected top element to be 2, got %v", top)
	}

	// Тестируем Pop для всех элементов
	stack.Pop() // удаляем 2
	stack.Pop() // удаляем 1

	// Проверка, что стек пуст после удаления всех элементов
	if !stack.IsEmpty() {
		t.Errorf("Expected stack to be empty after popping all elements")
	}

	// Тестируем Pop на пустом стеке
	item, ok = stack.Pop()
	if ok {
		t.Errorf("Expected pop to return false on empty stack, got %v", item)
	}

	// Тестируем Top на пустом стеке
	top, ok = stack.Top()
	if ok {
		t.Errorf("Expected top to return false on empty stack, got %v", top)
	}
}

func TestIterate(t *testing.T) {
	stack := New()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	expected := []interface{}{1, 2, 3}
	i := 0

	for item := range stack.Iterate() {
		if item != expected[i] {
			t.Errorf("Expected item to be %v, got %v", expected[i], item)
		}
		i++
	}
}
