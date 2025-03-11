package leetcode

type Submission struct {
	ID string `json:"id"`
	Title string `json:"title"`
	TitleSlug string `json:"titleSlug"`
	Timestamp int64 `json:"timestamp"`
}

type submissionResponse struct {
	Data struct {
		RecentAcSubmissionList []Submission `json:"recentAcSubmissionList"`
	} `json:"data"`
}