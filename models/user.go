package models

import (
	"regexp"
	"strings"
	"task-management/helper"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `gorm:"unique" json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Token    string `json:"token" form:"token"`
	Projects []Project
	Tasks    []Task
}

type NewDataUser struct {
	Name            string `json:"name" form:"name"`
	Email           string `json:"email" form:"email"`
	Password        string `json:"password" form:"password"`
	ConfirmPassword string `json:"confirmpassword" form:"confirmpassword"`
}

type LoginUser struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type GetUser struct {
	ID    uint
	Name  string
	Email string
}

type UpdateDataUser struct {
	Name            string `json:"name" form:"name"`
	Email           string `json:"email" form:"email"`
	Password        string `json:"password" form:"password"`
	ConfirmPassword string `json:"confirmpassword" form:"confirmpassword"`
	CurrentPassword string `json:"currentpassword" form:"currentpassword"`
}

func (new *NewDataUser) Validate() string {
	var message string

	if len(new.Name) < 3 {
		message = "Length of name must equal or greater than 3"
		return message
	}

	if strings.Count(new.Name, "  ") > 0 {
		message = "Your name should'nt be contain double or more space in one place"
		return message
	}

	var regex, _ = regexp.Compile(helper.NameRegex)

	if !regex.MatchString(new.Name) {
		message = "Your name must be only contain alphabet! Check if name contain unnecessary space too"
		return message
	}

	if strings.Count(new.Name, " ") > 2 {
		message = "Your name should be only contain not greater than 3 words"
		return message
	}

	var regex2, _ = regexp.Compile(helper.EmailRegex)

	if !regex2.MatchString(new.Email) {
		message = "Email format is invalid"
		return message
	}

	if len(new.Password) < 3 {
		message = "Length of password must equal or greater than 3"
		return message
	}

	if new.Password != new.ConfirmPassword {
		message = "Your password and confirmation password do not match"
		return message
	}

	return "OK"
}

func (new *UpdateDataUser) Validate() string {
	var message string

	if len(new.Name) != 0 && len(new.Name) < 3 {
		message = "Length of name must equal or greater than 3"
		return message
	}

	if len(new.Name) != 0 && strings.Count(new.Name, "  ") > 0 {
		message = "Your name should'nt be contain double or more space in one place"
		return message
	}

	var regex, _ = regexp.Compile(helper.NameRegex)

	if len(new.Name) != 0 && !regex.MatchString(new.Name) {
		message = "Your name must be only contain alphabet! Check if name contain unnecessary space too"
		return message
	}

	if len(new.Name) != 0 && strings.Count(new.Name, " ") > 2 {
		message = "Your name should be only contain not greater than 3 words"
		return message
	}

	var regex2, _ = regexp.Compile(helper.EmailRegex)

	if len(new.Email) != 0 && !regex2.MatchString(new.Email) {
		message = "Email format is invalid"
		return message
	}

	if len(new.Password) != 0 && len(new.Password) < 3 {
		message = "Length of password must equal or greater than 3"
		return message
	}

	if len(new.Password) != 0 && new.Password != new.ConfirmPassword {
		message = "Your password and confirmation password do not match"
		return message
	}

	return "OK"
}
