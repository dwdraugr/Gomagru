package auth

import (
	"net/http"
)

func GetHandler() {
	http.HandleFunc("/auth/login", LoginPage)
}

func LoginPage(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("BIIIIIIIIIIIIBA"))
	res.WriteHeader(http.StatusOK)
}

func RegistrationPage() {

}

func ForgottenPasswordPage() {

}

func PasswordResetPage() {

}

func Login() {

}

func Registration() {

}

func ForgottenPassword() {

}

func PasswordReset() {

}
