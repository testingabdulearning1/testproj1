package schemas

import "github.com/go-playground/validator/v10"

type UserLogin1 struct {
	UserID     string `json:"user_id"`
	UserEmail  string `json:"user_email"`
	RequestOtp bool   `json:"request_otp"`
}

func (u *UserLogin1) Validate() error {
	validator := validator.New()
	err := validator.Struct(u)
	if err != nil {
		return err
	}
	return nil
}

type ValidateOtp struct {
	UserID    string `json:"user_id"`
	UserEmail string `json:"user_email"`
	Otp       string `json:"otp"`
}

func (v *ValidateOtp) Validate() error {
	validator := validator.New()
	err := validator.Struct(v)
	if err != nil {
		return err
	}

	return nil
}

type ChangePassword struct {
	UserID          string `json:"user_id"`
	UserEmail       string `json:"user_email"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}

func (u *ChangePassword) Validate() error {
	validator := validator.New()
	err := validator.Struct(u)
	if err != nil {
		return err
	}

	return nil
}
