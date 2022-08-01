package middleware

import (
	"net/http"

	"github.com/SangiliG/cms-first/pkg/models"
	"github.com/golang-jwt/jwt"
)

func VerifyLogin(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenstr := r.URL.Query()["token"]
		if tokenstr != nil {
			claims := &models.Claims{}
			jwtKey := []byte("abc123")
			tkn, err := jwt.ParseWithClaims(tokenstr[0], claims, func(t *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})
			if err != nil {
				if err == jwt.ErrSignatureInvalid {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("Invalid token or token expires,Invalid Signature"))
					return
				}
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Error Occured"))
				return
			}
			if !tkn.Valid {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Invalid token or token expires"))
				return
			}
			//fmt.Println("Username")
			//fmt.Println(claims.Name)
			if claims.Role == "Admin" {
				f(w, r)
				return
			} else {
				w.Write([]byte("You're not Authorized person to access this API"))
			}
		}
		w.Write([]byte("Token Missing or Invalid Token"))
	}
}
