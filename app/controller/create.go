package controller

import (
	"github.com/Artem230819/itogWork/app/modal"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"path/filepath"
)

func AddPost(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//получаем значение из параметра text, переданного в форме запроса
	text := r.FormValue("text")

	//передаем значение в пост
	post := modal.NewPost(text)
	err := post.Add()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	defer http.Redirect(rw, r, "/", http.StatusSeeOther)
}

func CreatePosts(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//указываем путь к нужному файлу
	path := filepath.Join("public", "create.html")

	//создаем html-шаблон
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	//выводим шаблон в браузер
	err = tmpl.Execute(rw, nil)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}
