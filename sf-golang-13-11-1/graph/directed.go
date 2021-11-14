package graph

import (
	"errors"
)

// edgeWeight - ребро графа с весом
type edgeWeight struct {
	weight int
	node   *nodeWeight
}

// nodeWeight - вершина графа c весами
type nodeWeight struct {
	key   string
	edges map[string]*edgeWeight // Ребра вершины
}

type Path struct {
	Distance int
	Keys     []string
}

// Directed - ориентированный граф c весами
type Directed struct {
	nodes map[string]*nodeWeight // Вершины графа
}

func NewDirected() *Directed {
	return &Directed{nodes: make(map[string]*nodeWeight)}
}

// Add - добавляет вершину в граф
func (d *Directed) Add(key string) {
	// Если такая вершина уже существует, ничего не делаем
	if _, exists := d.nodes[key]; exists {
		return
	}
	d.nodes[key] = &nodeWeight{key: key, edges: make(map[string]*edgeWeight)}
}

// Connect - связывает 2 вершины графа
// Возвращает ошибку, если вершина отсуствует в графе
func (d *Directed) Connect(key1, key2 string, weight int) error {
	node1 := d.nodes[key1]
	node2 := d.nodes[key2]

	if node1 == nil || node2 == nil {
		return errors.New("неизвестная вершина графа")
	}

	node1.edges[key2] = &edgeWeight{weight, node2}
	return nil
}

// Dijkstra - поиск кратчайшего расстояния между заданными вершинами алгоритмом Дейкстры.
// Возвращает:
// - кратчайший путь Path или nil, если путь не найден
// - ошибку, если начальная точка не найдена
func (d *Directed) Dijkstra(startKey, targetKey string) (*Path, error) {
	startNode := d.nodes[startKey]
	if startNode == nil {
		return nil, errors.New("неизвестная вершина графа")
	}

	q := newQueueWeight()
	q.put(startNode)
	paths := make(map[string]*Path)
	paths[startNode.key] = &Path{0, []string{startNode.key}}

	if targetKey == startNode.key {
		return paths[targetKey], nil
	}

	for {
		n := q.get()
		if n == nil {
			break
		}

		p := paths[n.key]

		for _, near := range n.edges {
			q.put(near.node)
			if _, exists := paths[near.node.key]; !exists {
				paths[near.node.key] = &Path{
					Distance: p.Distance + near.weight,
					Keys:     append(p.Keys, near.node.key),
				}
			} else {
				if p.Distance+near.weight < paths[near.node.key].Distance {
					paths[near.node.key] = &Path{
						Distance: p.Distance + near.weight,
						Keys:     append(p.Keys, near.node.key),
					}
				}
			}
		}

	}
	return paths[targetKey], nil
}
