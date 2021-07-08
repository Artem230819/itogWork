package controller

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func GetPosts(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//получаем список все посты
	posts, err := model.GetAllPosts()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	//возвращаем список клиенту в формате JSON
	err = json.NewEncoder(rw).Encode(posts)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}