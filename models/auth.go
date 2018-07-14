package models

type IAuthenticate interface {
	authenticate(username string, password string)
}
