package model

type UserEmail string

type UserId string

type UserCredentials struct {
	Email  UserEmail
	UserId UserId
}
