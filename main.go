package main




func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	todos.add("Buy Pussy")
	todos.add("Buy Licker")
	todos.toggle(0)
	todos.print()
	storage.Save(todos)
}