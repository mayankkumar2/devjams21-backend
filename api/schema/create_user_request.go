package schema

type CreateUserRequest struct {
	IdToken string `json:"id_token"`
	Meta    struct {
		FirstName             string `json:"first_name"`
		LastName              string `json:"last_name"`
		RegNo                 string `json:"reg_no"`
		College               string `json:"college"`
		PhoneNumber           string `json:"phone_number"`
		Gender                string `json:"gender"`
		Degree                string `json:"degree"`
		Stream                string `json:"stream"`
		GraduationYear        string `json:"graduation_year"`
		Age                   uint   `json:"age"`
		Address               string `json:"address"`
		TShirtSize            string `json:"t_shirt_size"`
		AgreeMLHCodeOfConduct bool   `json:"is_mlh_codeofconduct"`
		AgreeMLHPrivacyPolicy bool   `json:"is_mlh_privacypolicy"`
		AgreeMLHEventDetails  bool   `json:"is_mlh_eventdetails"`
	} `json:"meta"`
}
