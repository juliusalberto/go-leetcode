package leetcode 

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Client struct {
	endpoint string
}

type GraphQLRequest struct {
	Query string `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

func (c* Client) GetRecentSubmission(username string, limit int) ([]Submission, error) {
	query := `
		query recentAcSubmissions($username: String!, $limit: Int!) {
			recentAcSubmissionList(username: $username, limit: $limit) {
				id
				title
				titleSlug
				timestamp
			}
		}
	`

	variables := map[string]interface{}{
		"username": username,
		"limit": limit,
	}

	reqBody := GraphQLRequest{
		Query: query,
		Variables: variables,
	}

	jsonBody, err := json.Marshal(reqBody)
	if (err != nil) {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result submissionResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Data.RecentAcSubmissionList, nil
}

