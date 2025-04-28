package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	ID    int
	Name  string
	Email string
}

type Pagination struct {
	Current int
	Pages   []int
}

type PageData struct {
	Data       []User
	Pagination Pagination
}

var tmpl *template.Template

func main() {
	var err error
	tmpl, err = template.ParseFiles("index.html", "table-content.html")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/data", dataHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	// Параметры запроса
	query := r.URL.Query().Get("search")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page == 0 {
		page = 1
	}

	// Имитация данных из БД с поиском и пагинацией
	data := mockDatabase(query, page)

	// Генерация пагинации
	pagination := Pagination{
		Current: page,
		Pages:   []int{1, 2, 3, 4, 5}, // Можно динамически рассчитать
	}

	// Рендер шаблона
	tmpl.ExecuteTemplate(w, "table-content", PageData{
		Data:       data,
		Pagination: pagination,
	})
}

func mockDatabase(query string, page int) []User {
	// Здесь должна быть реальная логика работы с БД
	// Это пример с моковыми данными
	var users []User
	for i := 1; i <= 20; i++ {
		users = append(users, User{
			ID:    i,
			Name:  "User " + strconv.Itoa(i),
			Email: "user" + strconv.Itoa(i) + "@example.com",
		})
	}

	// Фильтрация по запросу
	if query != "" {
		var filtered []User
		for _, u := range users {
			if strings.Contains(u.Name, query) || strings.Contains(u.Email, query) {
				filtered = append(filtered, u)
			}
		}
		return filtered
	}

	return users
}
