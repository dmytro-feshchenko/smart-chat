package main

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"net/http"

	"github.com/dgrijalva/jwt-go"
	models "github.com/technoboom/smart-chat/users/models"
)

// private and public RSA key paths
const (
	privKeyPath = "app.rsa"
	pubKeyPath  = "app.rsa.pub"
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

// fatal - checks error object and generates fatal log in case of error
func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// initKeys - reads keys from file system, generates fatal logs in case of error
func initKeys() {
	signBytes, err := ioutil.ReadFile(privKeyPath)
	fatal(err)

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	fatal(err)

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	fatal(err)

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	fatal(err)
}

// Response - represents server response data model
type Response struct {
	Data string `json:"data"`
}

// LoginHandler - provides an ability to login user and generate JWT token in
// case of success. Otherwise, generates error message.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user models.UserCredentials

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Println(w, "Error in request")
		return
	}

	if strings.ToLower(user.Username) != "someone" && user.Password != "p@ssword" {
		w.WriteHeader(http.StatusForbidden)
		fmt.Println("Error logging in")
		fmt.Fprint(w, "Invalid credentials")
		return
	}

	token := jwt.New(jwt.SigningMethodRS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error extracting the key")
		fatal(err)
	}

	tokenString, err := token.SignedString(signKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error while signing the token")
		fatal(err)
	}

	response := models.Token{Token: tokenString}
	JSONResponse(response, w)
}

// JSONResponse - writes response as application/json in http.ResponseWriter
func JSONResponse(response interface{}, w http.ResponseWriter) {
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

// Entry point for the program
func main() {
	initKeys()

	http.HandleFunc("/login", LoginHandler)

	http.ListenAndServe(":8080", nil)
}
