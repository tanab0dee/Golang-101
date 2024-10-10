package main

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todo.json")
	storage.Load(&todos)
	cmdFlags := NewCmdFlags()
	cmdFlags.Excecute(&todos)
	storage.Save(todos)
}
