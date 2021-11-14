package bst

import (
	"fmt"
	"strings"
)

const (
	padding     = "  " // Отбивка высоты узла
	prefixLeft  = "Л"  // Левый узел
	prefixRight = "П"  // Правый узел
	prefixRoot  = "К"  // Корневой узел
)

// node - узел бинарного дерева
type node struct {
	left  *node
	right *node
	value int
	data  string
}

// insert - добавляет дочерний элемент к узлу
func (n *node) insert(value int, data string) {
	if n == nil {
		return
	}

	// Выбираем целевой дочерний элемент для вставки:
	if value <= n.value {
		// левый
		if n.left == nil {
			n.left = &node{value: value, data: data}
		} else {
			n.left.insert(value, data)
		}
	} else {
		// ...или правый
		if n.right == nil {
			n.right = &node{value: value, data: data}
		} else {
			n.right.insert(value, data)
		}
	}
}

// max - ищет узел с максимальным значением в поддереве, обходя дерево справа.
// Возвращает найденный узел и его родителя.
func (n *node) max(parent *node) (*node, *node) {
	if n == nil {
		return nil, parent
	}
	if n.right == nil {
		return n, parent
	}
	return n.right.max(n)
}

// replace - заменяет текущий узел дерева узлом node
func (n *node) replace(parent, node *node) {
	if n == nil || parent == nil {
		return
	}
	if n == parent.left {
		parent.left = node
	} else {
		parent.right = node
	}
}

// delete - ищет и удаляет первый узел с указанным значением, начиная с текущего узла
func (n *node) delete(value int, parent *node) {
	if n == nil || parent == nil {
		return
	}

	switch {
	// 1. Если удаляемое значение меньше значения текущего узла,
	// то выполняем рекурсивное удаление над левой веткой
	case value < n.value:
		n.left.delete(value, n)

	// 2. Если удаляемое значение больше значения текущего узла,
	// то выполняем рекурсивное удаление над правой веткой
	case value > n.value:
		n.right.delete(value, n)

	// 3. Если удаляемое значение находится в текущем узле
	default:
		// 3.1 Если у текущего узла нет дочерних элементов, то удаляем узел
		if n.left == nil && n.right == nil {
			n.replace(parent, nil)
			return
		}

		// 3.2 Если у текущего узла есть один дочерний элемент,
		// то заменяем узел этим дочерним элементом
		if n.left == nil {
			n.replace(parent, n.right)
			return
		}
		if n.right == nil {
			n.replace(parent, n.left)
			return
		}

		// 3.3 Если у текущего узла есть оба дочерних элемента,
		// находим в левой ветке узел с максимальным значением (и его родителя).
		r, rParent := n.left.max(n)
		// Заменяем значение текущего узла значением найденного узла.
		// Таким образом выполнится условие: значение левого дочернего узла
		// будет не больше текущего узла
		n.value = r.value
		n.data = r.data
		// Рекурсивно удаляем найденный узел из дерева
		r.delete(r.value, rParent)
	}
}

// find - ищет узел с заданным значением в поддереве, начиная с текущего узла.
// Возвращает содержимое найденного узла и признак того, что узел найден
func (n *node) find(value int) (string, bool) {
	if n == nil {
		return "", false
	}
	switch {
	case value == n.value:
		return n.data, true
	case value < n.value:
		return n.left.find(value)
	default:
		return n.right.find(value)
	}
}

// print - выводит узел и его дочериние элементы на печать
// height - высота узла в дереве
// prefix - префикс узла на печати (правый, левый, корневой)
func (n *node) print(height int, prefix string) {
	if n == nil {
		return
	}
	fmt.Printf("%s%s: %d - %s\n", strings.Repeat(padding, height), prefix, n.value, n.data)
	n.left.print(height+1, prefixLeft)
	n.right.print(height+1, prefixRight)
}
