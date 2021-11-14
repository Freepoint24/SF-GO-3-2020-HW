package main

import (
	"fmt"
	"sf-golang-13-11-1/bst"
	"sf-golang-13-11-1/graph"
)

func main() {
	//
	// Бинарное дерево
	//
	fmt.Print("======\nБинарное дерево\n======\n\n")
	t := bst.NewTree()
	t.Insert(10, "Москва")
	t.Insert(5, "Тверь")
	t.Insert(7, "Торжок")
	t.Insert(12, "Клин")
	t.Insert(20, "Саратов")
	t.Insert(21, "Балаково")
	t.Insert(14, "Одинцово")
	t.Insert(1, "Санкт-Петербург")
	t.Insert(2, "Выборг")
	t.Print()

	fmt.Println("\n* Удаляем значение 5")
	t.Delete(5)
	t.Print()

	fmt.Println("\n* Ищем значение 21")
	result, found := t.Find(21)
	if !found {
		fmt.Println("Значение не найдено")
	} else {
		fmt.Println(result)
	}

	//
	// Неориентированный граф
	//
	fmt.Println("\n======\nНеориентированный граф\n======")
	u := graph.NewUndirected()
	u.Add("Блок")
	u.Add("Пастернак")
	u.Add("Мандельштам")
	u.Add("Цветаева")
	u.Add("Ахматова")
	u.Add("Маяковский")
	u.Add("Гумилев")
	u.Add("Есенин")
	u.Add("Хармс")
	u.Add("Белый")

	_ = u.Connect("Блок", "Пастернак")
	_ = u.Connect("Блок", "Цветаева")
	_ = u.Connect("Блок", "Маяковский")
	_ = u.Connect("Пастернак", "Мандельштам")
	_ = u.Connect("Пастернак", "Ахматова")
	_ = u.Connect("Пастернак", "Маяковский")
	_ = u.Connect("Цветаева", "Ахматова")
	_ = u.Connect("Цветаева", "Мандельштам")
	_ = u.Connect("Цветаева", "Хармс")
	_ = u.Connect("Маяковский", "Гумилев")
	_ = u.Connect("Гумилев", "Есенин")

	fmt.Println("\n* Ищем путь между `Блок` и `Есенин`")
	_ = u.BFS("Блок", "Есенин")
	fmt.Println("\n* Ищем путь между `Ахматова` и `Белый`")
	_ = u.BFS("Ахматова", "Белый")

	//
	// Ориентированный граф
	//
	fmt.Println("\n======\nОриентированный граф\n======")
	d := graph.NewDirected()
	d.Add("Москва")
	d.Add("Рига")
	d.Add("Варшава")
	d.Add("Прага")
	d.Add("Париж")

	_ = d.Connect("Москва", "Рига", 900)
	_ = d.Connect("Москва", "Варшава", 1100)
	_ = d.Connect("Рига", "Варшава", 600)
	_ = d.Connect("Рига", "Прага", 1400)
	_ = d.Connect("Варшава", "Прага", 900)
	_ = d.Connect("Рига", "Париж", 1300)
	_ = d.Connect("Варшава", "Париж", 1200)
	//_ = d.Connect("Прага", "Париж", 100)

	fmt.Println("\n* Ищем путь между `Москва` и `Париж`")
	p, _ := d.Dijkstra("Москва", "Париж")
	if p == nil {
		fmt.Println("Путь не найден")
	} else {
		fmt.Printf("Кратчайшее расстояние: %d => %v\n", p.Distance, p.Keys)
	}

	fmt.Println("\n* Ищем путь между `Москва` и `Прага`")
	p, _ = d.Dijkstra("Москва", "Прага")
	if p == nil {
		fmt.Println("Путь не найден")
	} else {
		fmt.Printf("Кратчайшее расстояние: %d => %v\n", p.Distance, p.Keys)
	}
}
