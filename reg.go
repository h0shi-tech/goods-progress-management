package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/signup", signupHandler).Methods("POST")
    r.HandleFunc("/login", loginHandler).Methods("POST")

    log.Fatal(http.ListenAndServe(":8080", r))
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
    // Получение данных из формы регистрации
    username := r.FormValue("username")
    password := r.FormValue("password")

    // Здесь можно добавить проверки данных пользователя (например, уникальность имени пользователя)

    // Создание нового пользователя в базе данных или другом хранилище
    // TODO: Добавить логику для создания пользователя

    fmt.Fprintf(w, "Регистрация прошла успешно! Добро пожаловать, %s!", username)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    // Получение данных из формы входа
    username := r.FormValue("username")
    password := r.FormValue("password")

    // Проверка введенных данных пользователя
    // TODO: Добавить логику для проверки имени пользователя и пароля

    // Если проверка прошла успешно, можно выполнить действия при успешной аутентификации
    // Например, можно создать сессию для пользователя и установить cookie

    fmt.Fprintf(w, "Вход выполнен успешно! Добро пожаловать, %s!", username)
}