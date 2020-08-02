package auth

import (
	"net/http"
)

func GetHandler() {
	http.HandleFunc("/auth/login", LoginPage)
}

func LoginPage(res http.ResponseWriter, req *http.Request)  {
	res.Write([]byte(""))
}

func RegistrationPage()  {

}

func ForgottenPasswordPage()  {
	
}

func PasswordResetPage() {

}

func Login() {
	
}

func Registration() {
	
}

func ForgottenPassword()  {
	
}

func PasswordReset()  {
	
}