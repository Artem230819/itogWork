package controller

import (
	"github.com/Artem230819/itogWork/app/modal"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"path/filepath"
)

func StartPage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//указываем путь к нужному файлу
	path := filepath.Join("public", "homePage.html")

	//создаем html-шаблон
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	// получаем все посты
	posts, err := modal.GetAllPosts()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	//выводим шаблон в браузер
	err = tmpl.Execute(rw, posts)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}
