package types

import (
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

const (
	minNameLength = 2
	maxNameLength = 15

	minPasswordLength = 6
	maxPasswordLength = 15
	bcryptCost        = 12
)

var (
	nameref     *regexp.Regexp
	emailref    *regexp.Regexp
	passwordref *regexp.Regexp
)

type User struct {
	ID                string `bson:"_id,omitempty" json:"id,omitempty"`
	Name              string `bson:"name" json:"name"`
	Email             string `bson:"email" json:"email"`
	EncryptedPassword string `bson:"encryptedPassword" json:"-"`
}

type CreateUserParams struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (cup CreateUserParams) Validate() map[string]string {
	errs := make(map[string]string)
	fmt.Printf("%+v\n", cup)

	if !userNameValidate(cup.Name) {
		errs["name"] = fmt.Sprintf("Имя должно быть длиной[%d,%d] и содержать только латинские символы", minNameLength, maxNameLength)
	}
	if !userEmailValidate(cup.Email) {
		errs["email"] = "Неправильный email"
	}
	if !userPasswordValidate(cup.Password) {
		errs["password"] = fmt.Sprintf("Пароль должен быть длиной[%d,%d]", minPasswordLength, maxPasswordLength)
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

func userNameValidate(name string) bool {
	if len(name) < minNameLength || len(name) > maxNameLength {
		return false
	}
	return nameref.MatchString(name)
}

func userEmailValidate(email string) bool {
	return emailref.MatchString(email)
}

func userPasswordValidate(password string) bool {
	if len(password) < minPasswordLength || len(password) > maxPasswordLength {
		return false
	}
	return passwordref.MatchString(password)
}

func init() {
	nameref = regexp.MustCompile(`^[a-zA-Z0-9]\w*[a-z0-9A-Z]$`)
	emailref = regexp.MustCompile(`^[a-z0-9.+-_]{1,}@[a-z0-9.]{1,}\.\w{2,}$`)
	passwordref = regexp.MustCompile(`^.+$`)
}
