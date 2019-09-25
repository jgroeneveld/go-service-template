package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func NewThirdPartyPostsApiClient() *ThirdPartyPostsApiClient {
	return &ThirdPartyPostsApiClient{
		Url: "https://jsonplaceholder.typicode.com",
	}
}

type ThirdPartyPostsApiClient struct {
	Url string
}

func (api *ThirdPartyPostsApiClient) GetPost(id int) (Post, error) {
	response, err := http.Get(fmt.Sprintf("%s/posts/%d", api.Url, id))
	if err != nil {
		return Post{}, err
	}
	defer response.Body.Close()

	var post Post
	return post, json.NewDecoder(response.Body).Decode(&post)
}

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
