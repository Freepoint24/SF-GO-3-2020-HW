package bst

import (
	"fmt"
)

// Tree - дерево бинарного поиска (BST)
type Tree struct {
	root *node
}

func NewTree() *Tree {
	return &Tree{root: nil}
}

// Insert - добавляет элемент в дерево
func (t *Tree) Insert(value int, data string) {
	if t.root == nil {
		t.root = &node{value: value, data: data}
	} else {
		t.root.insert(value, data)
	}
}

// Delete - удаляет элемент из дерева
func (t *Tree) Delete(value int) {
	if t.root == nil {
		return
	}
	tempRoot := &node{right: t.root}
	t.root.delete(value, tempRoot)
	if tempRoot.right == nil {
		t.root = nil
	}
}

// Find - ищет в дереве узел с заданным значением.
// Возвращает содержимое найденного узла и признак того, что узел найден
func (t *Tree) Find(value int) (string, bool) {
	return t.root.find(value)
}

// Print - выводит дерево на печать.
func (t *Tree) Print() {
	if t.root == nil {
		fmt.Println("<Пустое дерево>")
	} else {
		t.root.print(0, prefixRoot)
	}
}
