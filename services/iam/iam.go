package iam

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/mohammedfajer/Storj-REST-API/database"
	"github.com/mohammedfajer/Storj-REST-API/models"
	"github.com/mohammedfajer/Storj-REST-API/resources"
)

var jwtKey = []byte(os.Getenv("JWTKEY"))

func Login(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "from Login()")

	fmt.Println("from Login()")

	// check user address is already registered 
	// generate jwt token

	params := mux.Vars(r)

	var user models.DappUser 
	database.Db.First(&user, params["id"])
	// json.NewEncoder(w).Encode(user)

	if user.EthereumAddress == "" {
		w.WriteHeader(http.StatusUnauthorized)
		// json.NewEncoder(w).Encode(resources.Unauthorized{Message: "User is not registered"})
		return 
	}

	expirationTime := time.Now().Add(time.Minute * 5)
	claims := &resources.Claims{
		EthereumAddress: user.EthereumAddress,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
   	
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		// json.NewEncoder(w).Encode(resources.Unauthorized{Message: "Token is invalid"})
		return
	}
 
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: tokenString,
		Expires: expirationTime,
		Path: "/",
	})

	// json.NewEncoder(w).Encode(resources.TokenCreated{Message: "Token is created check your cookies"})
}

func Home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "from Home()")
	fmt.Println("from Home()")

	cookie, err := r.Cookie("token")
	
	if err != nil {
		fmt.Println(err)
		if err == http.ErrNoCookie {
			fmt.Println("error :", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		fmt.Println("error :", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := cookie.Value

	claims := &resources.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return 
		}
		fmt.Println("error: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return 
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Write([]byte(fmt.Sprintf("Hello, %s", claims.EthereumAddress)))
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "from Refresh()")

	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := cookie.Value

	claims := &resources.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return 
		}

		w.WriteHeader(http.StatusBadRequest)
		return 
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// if time.Unix(claims.ExpiresAt, 0).Sub(time.Now())  > 30 * time.Second {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	expirationTime := time.Now().Add(time.Minute * 5)
	claims.ExpiresAt = expirationTime.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, 
		&http.Cookie{
		Name: "token",
		Value: tokenString,
		Expires: expirationTime,
		Path: "/",
	})
}

func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "from Register()")

	
}
