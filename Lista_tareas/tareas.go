package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (t *taskList) appendTask(tl *task) {

	t.tasks = append(t.tasks, tl)
}

func (t *taskList) removeTask(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}



func(t *taskList) imprimirLista(){
	for _, tarea := range t.tasks {
		fmt.Println("Nombre",tarea.name)
		fmt.Println("Descripcion",tarea.description)
	}
}



func (t *taskList) imprimirListaCompletados(){
	for _, tarea := range t.tasks {
		if tarea.completed{
		fmt.Println("Nombre",tarea.name)
		fmt.Println("Descripcion",tarea.description)
		}
		
	}
}


type task struct {
	name        string
	description string
	completed   bool
}

func (t *task) markAsCompleted() {
	t.completed = true
}

func (t *task) updateName(name string) {
	t.name = name
}

func (t *task) updateDescription(description string) {
	t.description = description
}

func main() {
	t1 := &task{
		name:        "Curso de Go",
		description: "Curso de Go en Platzi en las proximas dos semanas",
	}
	t2 := &task{
		name:        "Curso de JavaScript",
		description: "mi curso de Async/Await en JavaScript",
	}
	list := &taskList{
		tasks: []*task{
			t1, t2,
		},// la coma al final es importante, ya que cuando defines structs idependientemente si existe un elemento o no despues, siempre debes colocarlo, para que no marque un error de sintaxis
	}
	t3 := &task{
		name:        "Construir mi propio servidor web",
		description: "Construir mi propio web server utilizando Go",
	}
	list.appendTask(t3)

	list.imprimirLista()
	list.tasks[0].markAsCompleted()
	fmt.Println("Tareas completadas")
	list.imprimirListaCompletados()

	mapaTareas := make(map[string]*taskList)

	mapaTareas["Leon"] =list

		t4 := &task{
		name:        "Curso de Rust",
		description: "Curso de Rust en Platzi en las proximas dos semanas",
	}
	t5 := &task{
		name:        "Curso de Node",
		description: "mi curso de Node",
	}
	list2 := &taskList{
		tasks: []*task{
			t4, t5,
		},// la coma al final es importante, ya que cuando defines structs idependientemente si existe un elemento o no despues, siempre debes colocarlo, para que no marque un error de sintaxis
	}
	fmt.Println("*************************************************************")
	mapaTareas["pepe"] =list2
	fmt.Println("tareas pepe")
	mapaTareas["pepe"].imprimirLista()
}