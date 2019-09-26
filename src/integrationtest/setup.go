package integrationtest

import (
	"github.com/go-chi/chi"
	"github.com/jgroeneveld/go-cloud-run-template/src/app"
)

type TestSetup struct {
	ThirdPartyPostsApi *FakeThirdPartyPostsApi
}

func (setup *TestSetup) Router() *chi.Mux {
	return app.Router(setup.ThirdPartyPostsApi)
}

func NewTestSetup() *TestSetup {
	return &TestSetup{
		ThirdPartyPostsApi: &FakeThirdPartyPostsApi{},
	}
}