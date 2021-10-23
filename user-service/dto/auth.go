package dto

//Auth has the dto for auth routes
type Auth struct {
	//Email of the user
	Email string `json:"email"`
	//Password of the user
	Password string `json:"password"`
}
