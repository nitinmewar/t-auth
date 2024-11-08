package authsvc

/* --------------------------------- SIGNUP --------------------------------- */
type SingupObject struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginObject struct {
	Email     string      `json:"email"`
	TypingDNA interface{} `json:"typing_dna"`
	Password  string      `json:"password"`
}

type EmailCheckRequest struct {
	Email string `json:"email"`
}

type EMailCheckResponse struct {
	Exist bool `json:"exist"`
}

type UserObject struct {
	PID       string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	CreatedAt int    `json:"created_at"`
}
