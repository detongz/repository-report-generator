package graphql

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"strings"
	"time"

	"github.com/machinebox/graphql"
	"github.com/tbbrave/repository-report-generator/graphql/model"
)

const timeFormat = "2006-01-02T15:04:05Z"

func parseTime(timeStr string) time.Time {
	timestamp, _ := time.Parse(timeFormat, timeStr)
	return timestamp
}

var startTime, endTime time.Time

type PullRequestCollection struct {
	Weeks   int
	WIP     []string
	BugFix  []string
	Improve []string
	Feature []string
	Other   []string
	Authors []string
}

type PullRequestContent struct {
	pullRequestURL string
	title          string
}

type PullRequestType string

const (
	WIP         PullRequestType = "WIP"
	Improvement PullRequestType = "Improvement"
	Improve     PullRequestType = "Improve"
	Feature     PullRequestType = "Feature"
	BugFix      PullRequestType = "Bug"
	Bug         PullRequestType = "BugFix"
	Other       PullRequestType = "Other"
)

var templ = `
	- %s
	%s`

func PullGithubData(repoOwner, repoName, start, end, token string) (collection PullRequestCollection) {

	wipAuthors, wip := pullGithubData(repoOwner, repoName, start, end, token, model.PullRequestStateOpen)
	for _, pr := range wip {
		collection.WIP = append(collection.WIP, fmt.Sprintf(templ, pr.title, pr.pullRequestURL))
	}

	authors, mergedPRs := pullGithubData(repoOwner, repoName, start, end, token, model.PullRequestStateMerged)

	collection.Authors = append(authors, wipAuthors...)

	for _, pr := range mergedPRs {
		switch getPullRequestType(pr.title) {
		case Improvement, Improve:
			collection.Improve = append(collection.Improve, fmt.Sprintf(templ, pr.title, pr.pullRequestURL))
		case BugFix, Bug:
			collection.BugFix = append(collection.BugFix, fmt.Sprintf(templ, pr.title, pr.pullRequestURL))
		case Feature:
			collection.Feature = append(collection.Feature, fmt.Sprintf(templ, pr.title, pr.pullRequestURL))
		default:
			collection.Other = append(collection.Other, fmt.Sprintf(templ, pr.title, pr.pullRequestURL))
		}
	}
	collection.Weeks = int(math.Ceil(float64(endTime.Sub(startTime)) / float64(time.Hour) / 7 / 24))

	return
}

func pullGithubData(repoOwner, repoName, start, end, token string, pullRequestType model.PullRequestState) (authorsList []string, prs []PullRequestContent) {
	startTime, endTime = parseTime(start), parseTime(end)

	authorsMap := map[string]int{}

	var repository struct {
		Repository model.Repository `json:"repository"`
	}

	httpClient := &http.Client{}
	client := graphql.NewClient("https://api.github.com/graphql", graphql.WithHTTPClient(httpClient))

	pageStart := ""
	for i := 0; i < 10; i++ {
		// make a request
		query := fmt.Sprintf(`
  query {
      repository(owner:"%s",name:"%s"){
          pullRequests(first:50,states:%s,orderBy:{field:UPDATED_AT,direction:DESC}%s) {
              totalCount
              pageInfo {
                  hasNextPage
                  hasPreviousPage
                  endCursor
              }
              nodes {
                  id
                  title
                  updatedAt
                  mergedAt
                  author {
					login
					url
					avatarUrl
					resourcePath
                  }
                  state
                  url
              }
          }
      }
  }
  `, repoOwner, repoName, pullRequestType, pageStart)

		req := graphql.NewRequest(query)

		// set header fields
		req.Header.Set("Authorization", "Bearer "+token)

		// run it and capture the response
		if err := client.Run(context.TODO(), req, &repository); err != nil {
			fmt.Println("request error: ", err)
		}

		authors, pullRequests, isBreak := processPullRequests(repository.Repository.PullRequests.Nodes)

		for i := range authors {
			authorsMap[authors[i]] = 1
		}

		prs = append(prs, pullRequests...)

		if isBreak {
			fmt.Println(pullRequestType, "request end: ", i)
			break
		}

		pageStart = fmt.Sprintf(`,after:"%s"`, *repository.Repository.PullRequests.PageInfo.EndCursor)

		time.Sleep(time.Second * 2) //query every 5sec, avoid flood github api
	}

	authorsList = make([]string, 0, len(authorsMap))

	for a := range authorsMap {
		authorsList = append(authorsList, a)
	}

	return
}

func processPullRequests(prs []*model.PullRequest) (authors []string, prResult []PullRequestContent, isBreak bool) {
	if len(prs) < 20 {
		isBreak = true
	}

	for i := range prs {
		aus, u, t, state := processPullRequest(prs[i])

		if aus != "" {
			authors = append(authors, aus)
		}
		if u != "" && t != "" {
			prResult = append(prResult, PullRequestContent{u, t})
		}

		if state {
			isBreak = true
			return
		}
	}

	return
}

func processPullRequest(pr *model.PullRequest) (author string, url, title string, isBreak bool) {
	var timeStr string

	if pr.Merged {
		timeStr = *pr.MergedAt
	} else {
		timeStr = pr.UpdatedAt
	}

	modTime := parseTime(timeStr)

	if modTime.After(endTime) { //skip
		fmt.Println("skip ", modTime, pr.Title, endTime)
		return
	}

	author = pr.Author.Login
	url = pr.URL
	title = pr.Title

	if modTime.Before(startTime) {
		isBreak = true
	}

	return
}

func getPullRequestType(title string) PullRequestType {
	if strings.Contains(title, fmt.Sprintf("[%s]", Feature)) {
		return Feature
	}

	if strings.Contains(title, fmt.Sprintf("[%s]", BugFix)) ||
		strings.Contains(title, fmt.Sprintf("[%s]", Bug)) {
		return Bug
	}

	if strings.Contains(title, fmt.Sprintf("[%s]", Improve)) ||
		strings.Contains(title, fmt.Sprintf("[%s]", Improvement)) {
		return Improve
	}

	return Other
}
