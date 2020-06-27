package middleware

import (
	"context"

	md "github.com/ebikode/eLearning-core/model"
	tr "github.com/ebikode/eLearning-core/translation"
	ut "github.com/ebikode/eLearning-core/utils"

	// "fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

// JwtUserAuthentication authenticate a user
//
// @param tokenSecret is used to encrypt password
//
// @param appKey is used to assert if the request is coming from a reliable source
func JwtUserAuthentication(tokenSecret, appKey string) func(next http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {

			response := make(map[string]interface{})

			// Any request that doesn't have app key on its header will be aborted
			// This is to curb unwarranted request
			//Grab the token from the header
			appkeyHeader := r.Header.Get("AppKey")

			// Set Token Error Message
			invalidAppKeyMsg := ut.Translate(tr.TParam{
				Key:          "error.unauthorized",
				TemplateData: nil,
				PluralCount:  nil,
			}, r)
			//AppKey is missing, returns with error code 403 Unauthorized
			if appkeyHeader != appKey || appkeyHeader == "" {
				response = ut.Message(false, invalidAppKeyMsg)

				ut.ErrorRespond(http.StatusUnauthorized, w, r, response)
				return
			}

			//List of endpoints that doesn't require auth
			notAuth := []string{
				"/api/v1/user/authenticate",
			}
			requestPath := r.URL.Path

			//check if request does not need authentication, serve the request if it doesn't need it
			for _, value := range notAuth {

				if value == requestPath {
					next.ServeHTTP(w, r)
					return
				}
			}

			// Set Token Error Message
			invalidTokenMsg := ut.Translate(tr.TParam{Key: "validation.invalid_token", TemplateData: nil, PluralCount: nil}, r)

			//Grab the token from the header
			tokenHeader := r.Header.Get("Authorization")
			//Token is missing, returns with error code 403 Unauthorized
			if tokenHeader == "" {
				response = ut.Message(false, invalidTokenMsg)

				ut.ErrorRespond(http.StatusUnauthorized, w, r, response)
				return
			}

			splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
			if len(splitted) != 2 {
				response = ut.Message(false, invalidTokenMsg)

				ut.ErrorRespond(http.StatusUnauthorized, w, r, response)
				return
			}

			tokenPart := splitted[1] //Grab the token part, what we are truly interested in
			tk := &md.UserTokenData{}

			token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
				return []byte(tokenSecret), nil
			})

			if err != nil { //Malformed token, returns with http code 403 as usual
				response = ut.Message(false, invalidTokenMsg)

				ut.ErrorRespond(http.StatusUnauthorized, w, r, response)
				return
			}

			if !token.Valid { //Token is invalid, maybe not signed on this server
				response = ut.Message(false, invalidTokenMsg)

				ut.ErrorRespond(http.StatusUnauthorized, w, r, response)
				return
			}

			//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
			// fmt.Sprintf("User %s", tk.Username) //Useful for monitoring
			ctx := context.WithValue(r.Context(), "tokenData", tk)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r) //proceed in the middleware chain!
		}
		return http.HandlerFunc(fn)
	}
}

// JwtAdminAuthentication ...
func JwtAdminAuthentication(tokenSecret string) func(next http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			//List of endpoints that doesn't require auth
			notAuth := []string{
				"/api/v1/admin/authenticate",
			}
			requestPath := r.URL.Path

			//check if request does not need authentication, serve the request if it doesn't need it
			for _, value := range notAuth {

				if value == requestPath {
					next.ServeHTTP(w, r)
					return
				}
			}

			// Set Token Messages
			invalidTokenMsg := ut.Translate(tr.TParam{Key: "validation.invalid_token", TemplateData: nil, PluralCount: nil}, r)

			response := make(map[string]interface{})
			tokenHeader := r.Header.Get("Authorization") //Grab the token from the header

			if tokenHeader == "" { //Token is missing, returns with error code 403 Unauthorized
				response = ut.Message(false, invalidTokenMsg)

				ut.ErrorRespond(http.StatusUnauthorized, w, r, response)
				return
			}

			splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
			if len(splitted) != 2 {
				response = ut.Message(false, invalidTokenMsg)

				ut.ErrorRespond(http.StatusUnauthorized, w, r, response)
				return
			}

			tokenPart := splitted[1] //Grab the token part, what we are truly interested in
			tk := &md.AdminTokenData{}

			token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
				return []byte(tokenSecret), nil
			})

			if err != nil { //Malformed token, returns with http code 403 as usual
				response = ut.Message(false, invalidTokenMsg)

				ut.ErrorRespond(http.StatusUnauthorized, w, r, response)
				return
			}

			if !token.Valid { //Token is invalid, maybe not signed on this server
				response = ut.Message(false, invalidTokenMsg)

				ut.ErrorRespond(http.StatusUnauthorized, w, r, response)
				return
			}

			//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
			//fmt.Sprintf("Admin %s", tk.Username) //Useful for monitoring
			ctx := context.WithValue(r.Context(), "tokenData", tk)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r) //proceed in the middleware chain!
		}
		return http.HandlerFunc(fn)
	}
}
