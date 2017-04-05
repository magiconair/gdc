// gdc shows the Github downloads count for Github repositories.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	log.SetFlags(0)

	var err error
	var repos []string

	switch len(os.Args) {
	case 2:
		user := os.Args[1]
		if repos, err = listRepos(user); err != nil {
			log.Fatalf("Fetching repos for %s failed. %s", user, err)
		}
	case 3:
		user, repo := os.Args[1], os.Args[2]
		repos = []string{user + "/" + repo}
	default:
		log.Fatalf("Usage: %s github-user [github-project]", os.Args[0])
	}

	var total int
	for _, repo := range repos {
		u := "https://api.github.com/repos/" + repo + "/releases"
		data, err := get(u)
		if err != nil {
			log.Printf("Failed to get releases for %s. %s", repo, err)
			continue
		}

		var x []struct {
			Assets []struct {
				Name  string `json:"name"`
				Count int    `json:"download_count"`
			}
		}
		if err := json.Unmarshal(data, &x); err != nil {
			log.Printf("Failed to unmarshal releases for %s. %s", repo, err)
			continue
		}

		for _, o := range x {
			for _, a := range o.Assets {
				fmt.Printf("%d\t%s\n", a.Count, a.Name)
				total += a.Count
			}
		}
	}
	fmt.Printf("%d\tTotal downloads\n", total)
}

// listRepos returns the github repositories for the given user.
func listRepos(user string) ([]string, error) {
	u := "https://api.github.com/users/" + user + "/repos"
	data, err := get(u)
	if err != nil {
		return nil, err
	}
	var x []struct {
		Name string `json:"full_name"`
	}
	if err := json.Unmarshal(data, &x); err != nil {
		return nil, err
	}
	var repos []string
	for _, r := range x {
		repos = append(repos, r.Name)
	}
	return repos, nil
}

// get executes an authenticated GET request and returns
// the body if the request was successful with 200 OK.
func get(rawurl string) ([]byte, error) {
	req, err := http.NewRequest("GET", rawurl, nil)
	if err != nil {
		return nil, err
	}
	if token := os.Getenv("GITHUB_TOKEN"); token != "" {
		req.Header.Set("Authorization", "token "+token)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s: Status Code: %d, Body: %s", rawurl, resp.StatusCode, string(data))
	}
	return data, err
}
