//template.go
package template

import (
	"fmt"
	"githubapi-golang/domain"
	"io/ioutil"
	"log"
)

const (
	HTMLDocs = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <link rel="stylesheet" href="./html/style.css" />
    <title>Github Profiles</title>
  </head>
  <body>
    <form action="/create" method="post" class="user-form" id="form">
      <input type="text" id="search" name="username" placeholder="Search a Github User"/>
  	</form>
	<main id="main">
		%s
	</main>
  </body>
<html>
`

	UserCardHTML = `
<div class="card">
	<div>
    	<img src="%s" alt="%s" class="avatar">
	</div>
	<div class="user-info">
		<h2>%s</h2>
		<p>%s</p>
		<ul>
			<li>%d<strong>Followers</strong></li>
			<li>%d<strong>Following</strong></li>
			<li>%d<strong>Repos</strong></li>
		</ul>
	<div id="repos">
		%s
	</div>
</div>
`

	RepoHTML = `
		<a class="repo" href="%s" target="_blank">
		%s
	</a>
`

	CardErrHTML = `
<div class="card">
	<h1>No profile with this username</h1>
</div>
`
)

func OutPutFile(user domain.UserInfo, repos []domain.Repo) {
	var result string
	if user != (domain.UserInfo{}) {
		userCardHTML := fmt.Sprintf(
			UserCardHTML,
			user.AvatarURL,
			user.Name,
			user.Name,
			user.Bio,
			user.Followers,
			user.Following,
			user.PublicRepos,
			"%s",
		)
		var repoHTML string
		for i := 0; i < len(repos); i++ {
			repoHTML += fmt.Sprintf(
				RepoHTML,
				repos[i].HTMLURL,
				repos[i].Name,
			)
		}

		userCardHTML = fmt.Sprintf(userCardHTML, repoHTML)
		result = fmt.Sprintf(HTMLDocs, userCardHTML)

	} else {
		result = fmt.Sprintf(HTMLDocs, CardErrHTML)
	}

	// ファイルの書き込み
	err := ioutil.WriteFile("./html/result.html", []byte(result), 0777)
	if err != nil {
		log.Fatal(err)
	}
}
