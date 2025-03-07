package repository

import (
	"encoding/json"
	"net/http"

	"github.com/vadgun/goApp/user_registration/internal/entity"
)

type ExternalAPIRepositoryImp struct {
}

func NewExternalAPIRepository() *ExternalAPIRepositoryImp {
	return &ExternalAPIRepositoryImp{}
}

func (r *ExternalAPIRepositoryImp) FetchUser() (*entity.ExternalUser, error) {

	response, err := http.Get("https://randomuser.me/api/")
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var result struct {
		Results []struct {
			Gender string `json:"gender"`
			Name   struct {
				First string `json:"first"`
				Last  string `json:"last"`
			}
			Email string `json:"email"`
		}
	}

	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, err
	}

	user := &entity.ExternalUser{
		Gender: result.Results[0].Gender,
		Name:   result.Results[0].Name.First + " " + result.Results[0].Name.Last,
		Email:  result.Results[0].Email,
	}

	return user, nil

}
