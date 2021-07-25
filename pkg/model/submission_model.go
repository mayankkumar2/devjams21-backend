package model

type Submission struct {
	BaseModel
	Meta map[string]interface{} `json:"meta"`
}
