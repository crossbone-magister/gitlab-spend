package logic

import (
	"fmt"
	"gitlab-spend/config"
	"gitlab-spend/issue"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func RegisterTimeSpent(issue issue.Issue, client *http.Client, config config.GitlabSpendConfiguration) (*http.Response, error) {
	var requestBody = url.Values{"duration": {issue.Duration.String()}, "summary": {issue.Details()}}
	log.Println(requestBody)
	var apiUrl = fmt.Sprintf("https://%s/api/v4/projects/%s/issues/%s/add_spent_time", config.Host(), url.PathEscape(issue.Project), issue.Iid)
	var request, _ = http.NewRequest(http.MethodPost, apiUrl, strings.NewReader(requestBody.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("PRIVATE-TOKEN", config.Token())
	log.Println(request)
	return client.Do(request)
}
