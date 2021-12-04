package gitlab

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

func NewGitLabWrapStats(userName string) (*WrapStats, error) {
	return getWrapStats(userName)
}

func getCount(url string) (int, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Add("PRIVATE-TOKEN", os.Getenv("GITLAB_TOKEN"))
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	count, err := strconv.Atoi(resp.Header.Get("x-total"))
	if err != nil {
		return 0, err
	}
	return count, nil
}

func getProjectsCount(id, startpage, count int) (int, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://gitlab.com/api/v4/users/%d/projects?page=%d&per_page=100&membership=true", id, startpage), nil)
	if err != nil {
		return 0, err
	}
	req.Header.Add("PRIVATE-TOKEN", os.Getenv("GITLAB_TOKEN"))
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		var projects Projects
		if err := json.NewDecoder(resp.Body).Decode(&projects); err != nil {
			return 0, err
		}
		start := time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC)
		for _, project := range projects {
			if project.CreatedAt.After(start) {
				count++
			}
		}
		if resp.Header.Get("x-next-page") != "" {
			nextpage, err := strconv.Atoi(resp.Header.Get("x-next-page"))
			if err != nil {
				return 0, err
			}
			return getProjectsCount(id, nextpage, count)
		}
		return count, nil
	}
	return 0, fmt.Errorf(fmt.Sprintf("%d", resp.StatusCode))
}

func getUserDetails(id int) (*User, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://gitlab.com/api/v4/users/%d", id), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("PRIVATE-TOKEN", os.Getenv("GITLAB_TOKEN"))
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		var user User
		if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
			return nil, err
		}

		return &user, nil
	}
	return nil, fmt.Errorf(fmt.Sprintf("%d", resp.StatusCode))
}

func getWrapStats(userName string) (*WrapStats, error) {
	resp, err := http.Get(fmt.Sprintf("https://gitlab.com/api/v4/users?username=%s", url.QueryEscape(userName)))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		var users Users
		if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
			return nil, err
		}
		if len(users) == 0 {
			return nil, errors.New("user not found")
		}
		details, err := getUserDetails(users[0].ID)
		if err != nil {
			return nil, err
		}
		projects, err := getProjectsCount(details.ID, 1, 0)
		if err != nil {
			return nil, err
		}
		issues, err := getCount(fmt.Sprintf("https://gitlab.com/api/v4/issues?author_id=%d&scope=all&created_after=2021-01-01T08:00:00Z", details.ID))
		if err != nil {
			return nil, err
		}
		mergeRequests, err := getCount(fmt.Sprintf("https://gitlab.com/api/v4/merge_requests?author_id=%d&scope=all&created_after=2021-01-01T08:00:00Z", details.ID))
		if err != nil {
			return nil, err
		}
		starredProject, err := getCount(fmt.Sprintf("https://gitlab.com/api/v4/users/%d/starred_projects", details.ID))
		if err != nil {
			return nil, err
		}

		return &WrapStats{
			User:             details,
			IssuesSum:        issues,
			MergeRequestsSum: mergeRequests,
			ProjectSum:       projects,
			StarredProjects:  starredProject,
		}, nil
	}
	return nil, errors.New("no GitLab user found")
}
