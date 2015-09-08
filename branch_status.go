/*
branch_status implements a simple github api call to retrevie status of a branch (or a commit).
I am using it along with a bash shortcut.

All my projects are cloned in one directroy. So a simple alias will do what I need.


	alias bstat='for i in `ls`; do printf "$i: "; branch_status -t $(cat ~/.boost/git_token) -o $(cat ~/.boost/git_org) -r $i;done'


	$ bstat
	my_branch: success
	my_nextbr: failed
	...

*/
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Status struct {
	State string
}

func _error(err error) {
	if err != nil {
		panic(err)
	}
}

func get_content(url string) string {
	res, err := http.Get(url)
	_error(err)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	_error(err)
	return string(body)
}

func main() {
	token := flag.String("t", "", "Github access tokens. repo:status access is sufficient.")
	owner := flag.String("o", "", "Repository owenr.")
	repo := flag.String("r", "", "Repository name")
	branch := flag.String("b", "master", "Repository name")
	flag.Parse()
	s := get_content(fmt.Sprintf("https://api.github.com/repos/%s/%s/commits/%s/statuses?access_token=%s", *owner, *repo, *branch, *token))
	var data []Status
	json.Unmarshal([]byte(s), &data)

	if len(data) > 0 {
		fmt.Printf("%s", data[0].State)
	} else {
		fmt.Printf("no status")
	}

	fmt.Println("")
}
