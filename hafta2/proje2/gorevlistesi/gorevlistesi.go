package main

import (
	"fmt"
	"os"
)

type Task struct {
	Description string
	Completed   bool
}

var tasks []Task

func addTask(description string) {
	tasks = append(tasks, Task{Description: description, Completed: false})
}

func listTasks() {
	for i, task := range tasks {
		status := "Tamamlanmadı"
		if task.Completed {
			status = "Tamamlandı"
		}
		fmt.Printf("%d. %s (%s)\n", i+1, task.Description, status)
	}
}

func completeTask(taskNumber int) {
	if taskNumber >= 1 && taskNumber <= len(tasks) {
		tasks[taskNumber-1].Completed = true
		fmt.Println("Görev tamamlandı!")
	} else {
		fmt.Println("Geçersiz görev numarası.")
	}
}

func main() {
	for {
		fmt.Println("\nGörev Listesi Uygulaması")
		fmt.Println("1. Görev Ekle")
		fmt.Println("2. Görevleri Listele")
		fmt.Println("3. Görev Tamamla")
		fmt.Println("4. Çıkış")

		var choice int
		fmt.Print("Seçiminizi yapın: ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Hatalı giriş. Lütfen tekrar deneyin.")
			continue
		}

		switch choice {
		case 1:
			var description string
			fmt.Print("Görev açıklamasını girin: ")
			_, _ = fmt.Scan(&description)
			addTask(description)
		case 2:
			listTasks()
		case 3:
			var taskNumber int
			fmt.Print("Tamamlanan görevin numarasını girin: ")
			_, _ = fmt.Scan(&taskNumber)
			completeTask(taskNumber)
			listTasks()
		case 4:
			fmt.Println("Çıkış yapılıyor..")
			os.Exit(0)
		default:
			fmt.Println("Geçersiz seçim. Tekrar deneyin.")
		}
	}
}
