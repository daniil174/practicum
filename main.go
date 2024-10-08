package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const form = `<html>
    <head>
    <title></title>
    </head>
    <body>
        <form action="/" method="post">
            <label>Логин <input type="text" name="login"></label>
            <label>Пароль <input type="password" name="password"></label>
            <input type="submit" value="Login">
        </form>
    </body>
</html>`

func Auth(login, password string) bool {
	return login == `guest` && password == `demo`
}

type Subj struct {
	Product string `json:"name"`
	Price   int    `json:"price"`
}

func JSONHandler(w http.ResponseWriter, req *http.Request) {
	// собираем данные
	subj := Subj{"Milk", 50}
	// кодируем в JSON
	resp, err := json.Marshal(subj)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// устанавливаем заголовок Content-Type
	// для передачи клиенту информации, кодированной в JSON
	w.Header().Set("content-type", "application/json")
	// устанавливаем код 200
	w.WriteHeader(http.StatusOK)
	// пишем тело ответа
	w.Write(resp)
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		login := r.FormValue("login")
		password := r.FormValue("password")
		if Auth(login, password) {
			io.WriteString(w, "Добро пожаловать!")
		} else {
			http.Error(w, "Неверный логин или пароль", http.StatusUnauthorized)
		}
		return
	} else {
		io.WriteString(w, form)
	}
}

func apiPage(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Это страница /api."))
}

func WriteHandle(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "1")
	fmt.Fprint(w, "2")
	w.Write([]byte("3"))
}

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir(".."))

	mux.HandleFunc(`/api/`, apiPage)
	//mux.HandleFunc(`/`, mainPage)
	mux.HandleFunc(`/api/getSubject`, JSONHandler)
	mux.HandleFunc(`/test`, WriteHandle)
	mux.Handle(`/golang/`, http.StripPrefix(`/golang/`, fs))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./main.go")
	})

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
