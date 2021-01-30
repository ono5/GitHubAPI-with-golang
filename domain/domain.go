// domain.go
package domain

type UserInfo struct {
	AvatarURL   string `json:"avatar_url"`
	Name        string `json:"name"`
	Bio         string `json:"bio"`
	Followers   int    `json:"followers"`
	Following   int    `json:"following"`
	PublicRepos int    `json:"public_repos"`
}

type Repo struct {
	Name    string `json:"name"`
	HTMLURL string `json:"html_url"`
}

type Repos struct {
	Repos []Repo `json:"repos"`
}
