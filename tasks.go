package main

import (
    "fmt"
    "time"
)

type Task struct {
    Title       string
    Description string
    DueDate     time.Time
    Priority    int
}

func main() {
    // Создание новой задачи
    newTask := Task{
        Title:       "Покупка товаров",
        Description: "Купить товары 1,2,3 и т.д.",
        DueDate:     time.Date(2022, time.July, 31, 0, 0, 0, 0, time.UTC),
        Priority:    1,
    }

    // Вывод информации о задаче
    fmt.Println("Название задачи:", newTask.Title)
    fmt.Println("Описание задачи:", newTask.Description)
    fmt.Println("Дата выполнения:", newTask.DueDate.Format("02.01.2006"))
    fmt.Println("Приоритет:", newTask.Priority)
}