package models

type Login struct {
	Username string `json:"username" validate:"required" default:"admin"`
	Password string `json:"password" validate:"required" default:"admin"`
}
