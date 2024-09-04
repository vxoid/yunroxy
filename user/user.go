package user

// constants
const minPass = 10
const maxPass = 45 
const minEmail = 8
const maxEmail = 20

type User struct{
	mail string
	pass string
	
}

func UserConstruct() *User {
	return &User{mail : EmailGenerator(minEmail, maxEmail),
			     pass : PassGenerator(minPass, maxPass)}
}