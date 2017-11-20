package router

import (
  "context"
  "net/http"
  "time"
  "authentication"
)

const ContentKey int = 0

// middleware to protect private pages
func validateToken(page http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
    token := req.Header.Get("Authorization")

    if token != "" {
      claim, err := authentication.DecodeToken(token)
      if err == nil {
  			current := time.Now().Unix()
  			if current < claim.ExpiresAt {
          ctx := context.WithValue(req.Context(), ContentKey, claim)
			    page(res, req.WithContext(ctx))
  			} else {
          http.Error(res, "Token Expired", http.StatusUnauthorized)
  			}
  		} else {
  			http.Error(res, "Unauthorized", http.StatusUnauthorized)
  		}
    } else {
      http.Error(res, "Token Required", http.StatusUnauthorized)
    }
	})
}

func getClaim(req *http.Request) (authentication.UserClaim, bool) {
  claim, ok := req.Context().Value(ContentKey).(authentication.UserClaim)
  return claim, ok
}
