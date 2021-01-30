//plae.go
package template

import (
	"fmt"
	"githubapi-golang/domain"
	"io/ioutil"
	"log"
)

const (
	HTMLDocs = `
<!DOCTYPEl
l lng="en">
  <head>
  etacharset="UTF-8" />
  eta="iewport" content="width=device-width, initial-scale=1.0" />
    <link rel="stylesheet" href="style.css" />
    <title>Github Profiles</title>
  </head>
  <body>
	<form class="user-form" id="form">
		<div><input type="text" name="post" id="search" placeholder="Search a Github User"></div>
	</form>
	<main id="main">
		%s
	</main>
  </body>
<l
`

	UserCardHTML = `
<div class="card">
	<div>
	g sc="%s" alt="%s" class="avatar">
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
		<a class="repo" href="%s" target="_blank">%s</a>
`

	CardErrHTML = `
<div class="card">
	<h1>No profile with this usee</h1>
</div>
`
)

func outPutFile(user domain.UserInfo, repos []domain.Repo) {
	var result string
	if user != (domain.UserInfo{}) {
		userCardHTML := fmt.Sprintf(
			user.AvatarURL,
			user.Name,
			user.Name,
			user.Bio,
			user.Followers,
			user.Following,
			user.PublicRepos,
		)
		var repoHTML string
		for i := 0; i < len(repos); i++ {
			L
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
	err := ioutil.WriteFile("./html/index.html", []byte(result), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
