package schema

type CreateUserRequest struct {
	IdToken string `json:"id_token"`
	Meta    struct {
		RegNo          string `json:"reg_no"`
		College        string `json:"college"`
		PhoneNumber    string `json:"phone_number"`
		Gender         string `json:"gender"`
		Degree         string `json:"degree"`
		Stream         string `json:"stream"`
		GraduationYear string `json:"graduation_year"`
		Age            uint   `json:"age"`
		Address        string `json:"address"`
		TShirtSize     string `json:"t_shirt_size"`
	} `json:"meta"`
}
