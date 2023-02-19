package models

type UserEmail string

type UserModel struct {
	Model
	Email    UserEmail `bson:"email"`
	Nickname string    `bson:"nickname"`
	Picture  string    `bson:"picture,omitempty"`
}
