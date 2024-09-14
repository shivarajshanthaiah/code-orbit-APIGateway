package model

type User struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type OTP struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Password represents the structure for updating password.
type Password struct {
	Old     string `json:"old_password"`
	New     string `json:"new_password"`
	Confirm string `json:"confirm_password"`
}
