package integrationtest

import (
	"go-cloud-run-template/app"
)

type FakeThirdPartyPostsApi struct {
}

func (api *FakeThirdPartyPostsApi) GetPost(id int) (app.Post, error) {
	return app.Post{
		UserId: 42,
		Id:     12,
		Title:  "fake title",
		Body:   "fake body",
	}, nil
}
