package types

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost = 12
)

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name              string             `bson:"name" json:"name"`
	Email             string             `bson:"email" json:"email"`
	EncryptedPassword string             `bson:"encryptedPassword" json:"-"`
}

type CreateUserParams struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (cup CreateUserParams) Validate() map[string]string {
	errs := make(map[string]string)

	if len(cup.Name) < 2 {
		errs["name"] = "Length of name is less than 2"
	}
	if len(cup.Email) < 2 {
		errs["email"] = "Length of email is less than 2"
	}
	if len(cup.Password) < 2 {
		errs["email"] = "Length of password is less than 2"
	}
	return errs
}

func NewUserFromParams(params CreateUserParams) (*User, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
	if err != nil {
		return nil, err
	}
	return &User{
		Name:              params.Name,
		Email:             params.Email,
		EncryptedPassword: string(encpw),
	}, nil

}

func (u *User) IsValidPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password))
	if err != nil {
		return fmt.Errorf("invalid credentials")
	}
	return nil
}
