package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"

	"github.com/modood/table"
	"github.com/tomnomnom/linkheader"
)

type repo struct {
	Name              string `json:"name"`
	Language          string `json:"language"`
	Stars             int    `json:"stargazers_count"`
	Forks             int    `json:"forks_count"`
	Open_issues_count int    `json:"open_issues_count"`
	Created_at        string `json:"created_at"`
	Updated_at        string `json:"updated_at"`
	Pushed_at         string `json:"pushed_at"`
	Description       string `json:"description"`
}

var sortBy *string

func main() {
	org := flag.String("org-name", "kubernetes", "The org name you want to list")
	sortBy = flag.String("sort-by", "star", "The field name you want to sort by, "+
		"include: star, fork, issue, createtime, updatetime, pushedtime. (default: star)")
	flag.Parse()

	allRepos := repos{}

	intUrl := fmt.Sprintf("https://api.github.com/orgs/%s/repos", *org)

	for intUrl != "" {
		resp, err := http.Get(intUrl)
		if err != nil {
			fmt.Println("failed to get all repos, err: %v", err)
		}

		data := make([]byte, 0)
		data, _ = ioutil.ReadAll(resp.Body)

		repos := []repo{}
		err = json.Unmarshal(data, &repos)
		if err != nil {
			fmt.Println(err)
			return
		}
		allRepos = append(allRepos, repos...)

		intUrl = ""
		links := linkheader.Parse(strings.Join(resp.Header["Link"], ","))
		for _, link := range links {
			if link.Rel == "next" {
				intUrl = link.URL
				break
			}
		}
	}

	sort.Sort(allRepos)

	fmt.Println("repos count: ", len(allRepos))

	t := table.Table(allRepos)
	fmt.Println(t)
}

type repos []repo

func (r repos) Len() int { return len(r) }

func (r repos) Less(i, j int) bool {
	switch *sortBy {
	case "star":
		return r[i].Stars > r[j].Stars
	case "fork":
		return r[i].Forks > r[j].Forks
	case "issue":
		return r[i].Open_issues_count > r[j].Open_issues_count
	case "createtime":
		return r[i].Created_at > r[j].Created_at
	case "updatetime":
		return r[i].Updated_at > r[j].Updated_at
	case "pushedtime":
		return r[i].Pushed_at > r[j].Pushed_at
	}
	return r[i].Stars > r[j].Stars
}

func (r repos) Swap(i, j int) { r[i], r[j] = r[j], r[i] }
