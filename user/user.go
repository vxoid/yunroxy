package user

const minPass = 10
const maxPass = 45
const minEmail = 8
const maxEmail = 20

type UserCredentials struct {
	email string
	pass  string
}

func NewRandom() *UserCredentials {
	return &UserCredentials{email: EmailGenerator(minEmail, maxEmail), pass: PassGenerator(minPass, maxPass)}
}

func (u *UserCredentials) GetEmail() string {
	return u.email
}

func (u *UserCredentials) GetPassword() string {
	return u.pass
}
