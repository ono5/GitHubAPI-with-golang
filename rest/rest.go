package rest

import (
	"encoding/json"
	"githubapi-golang/domain"
	"net/http"
)

type getRequest struct {
	endPoint string
}

func NewGetRequest(endPoint string) domain.GetRequest {
	return &getRequest{
		endPoint: endPoint,
	}
}

func (g *getRequest) GetUser(username string) *domain.UserInfo {
	var u domain.UserInfo
	resp, err := http.Get(g.endPoint + username)
	if err != nil {
		return &u
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&u)
	if err != nil {
		return &u
	}
	return &u
}

func (g *getRequest) GetRepos(username string) *[]domain.Repo {
	var rs []domain.Repo
	resp, err := http.Get(g.endPoint + username + "/repos?sort=created")
	if err != nil {
		return &rs
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&rs)
	if err != nil {
		return &rs
	}
	return &rs
}
