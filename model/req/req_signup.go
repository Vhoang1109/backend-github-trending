package req

type ReqSignUp struct {
	FullName string `json:"fullName,omitempty" validate:"required"` // tags name
	Email    string `json:"email,omitempty" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
}
