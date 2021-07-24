package schema

type CreateUserRequest struct {
	IdToken string `json:"id_token"`
	Meta    struct {
		RegNo   string `json:"reg_no"`
		College string `json:"college"`
	} `json:"meta"`
}
