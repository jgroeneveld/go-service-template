package app

import (
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

type ThirdPartyPostsApi interface {
	GetPost(id int) (Post, error)
}

func getPostHandler(api ThirdPartyPostsApi) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		postID, err := strconv.Atoi(chi.URLParam(r, "postID"))
		PanicIf(err)

		post, err := api.GetPost(postID)
		PanicIf(err)

		response := GetPostHandlerResponse{
			UserId: post.UserId,
			Id:     post.Id,
			Title:  post.Title,
			Body:   post.Body,
		}

		err = WriteJson(w, response)
		PanicIf(err)
	}
}

type GetPostHandlerResponse struct {
	UserId int    `json:"user_id"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
