package main

import (
  "net/http"
  "strings"

  "github.com/dgrijalva/jwt-go"
  "github.com/dgrijalva/jwt-go/request"
  "github.com/golang/glog"

  "github.com/ckbball/dev-edge/auth"
)

// allowCORS allows Cross Origin Resoruce Sharing from any origin.
// Don't do this without consideration in production systems.
func allowCORS(h http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if origin := r.Header.Get("Origin"); origin != "" {
      w.Header().Set("Access-Control-Allow-Origin", origin)
      if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
        preflightHandler(w, r)
        return
      }
    }
    h.ServeHTTP(w, r)
  })
}

// preflightHandler adds the necessary headers in order to serve
// CORS from any origin using the methods "GET", "HEAD", "POST", "PUT", "DELETE"
// We insist, don't do this without consideration in production systems.
func preflightHandler(w http.ResponseWriter, r *http.Request) {
  headers := []string{"Content-Type", "Accept", "Authorization"}
  w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
  methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
  w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
  glog.Infof("preflight request for %s", r.URL.Path)
}

// add auth middleware here

func UserAuth(h http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    token, err := request.ParseFromRequest(c.Request, MyAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
      b := auth.GetKey()
      return b, nil
    }, request.WithClaims(&auth.CustomClaims{}))

    if err != nil {
      http.Error(w, http.StatusText("Unauthorized"), 401)
      return
    }

    claims, err := auth.DecodeWithCustomClaims(token)
    if err != nil {
      fmt.Printf(err)
      http.Error(w, http.StatusText("Unauthorized"), 401)
      return
    }

    user := claims.User

    ctx := context.WithValue(r.Context(), "user", user)
    h.ServeHTTP(w, r.WithContext(ctx))
  })
}

func stripBearerPrefixFromTokenString(tok string) (string, error) {
  if len(tok) > 5 && strings.ToUpper(tok[0:6]) == "TOKEN " {
    return tok[6:], nil
  }
  return tok, nil
}

// Extract  token from Authorization header
// Uses PostExtractionFilter to strip "TOKEN " prefix from header
var AuthorizationHeaderExtractor = &request.PostExtractionFilter{
  request.HeaderExtractor{"Authorization"},
  stripBearerPrefixFromTokenString,
}

// Extractor for OAuth2 access tokens.  Looks in 'Authorization'
// header then 'access_token' argument for a token.
var MyAuth2Extractor = &request.MultiExtractor{
  AuthorizationHeaderExtractor,
  request.ArgumentExtractor{"access_token"},
}
