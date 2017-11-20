package authentication

import (
  "fmt"
  "context"
  "time"
  "gopkg.in/mgo.v2/bson"
  "net/http"
  jwt "gopkg.in/dgrijalva/jwt-go.v3"
)

const signatureKey string = "aventador"

type UserClaim struct {
    Data bson.M `json:"data" bson:"data"`
    // recommended having
    jwt.StandardClaims
}

func CreateToken(data bson.M) string {
  // Expires the token and cookie in 1 hour
  expireToken := time.Now().AddDate(1, 0, 0).Unix()
  // We'll manually assign the claims but in production you'd insert values from a database
  claims := UserClaim {
    data,
    jwt.StandardClaims {
      ExpiresAt: expireToken,
    },
  }

  // Create the token using your claims
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

  // Signs the token with a key.
  signedToken, _ := token.SignedString([]byte(signatureKey))
  return signedToken
}

func DecodeToken(signedToken string) (UserClaim, error) {
	token, err := jwt.ParseWithClaims(signedToken, &UserClaim{},
    func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				 return nil, fmt.Errorf("Unexpected signing method")
			}
			return []byte(signatureKey), nil
		})


	if err != nil {
		return UserClaim{}, err
	}

	if claims, ok := token.Claims.(*UserClaim); ok && token.Valid {
    return *claims, nil
			//page(res, req.WithContext(ctx))
	} else {
		return UserClaim{}, fmt.Errorf("Unexpected signing method")
	}
}

const ContentKey int = 0
// middleware to protect private pages
func ValidateToken(page http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
    token := req.Header.Get("Authorization")

    if token != "" {
      claim, err := DecodeToken(token)
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

func GetClaim(req *http.Request) (UserClaim, bool) {
  claim, ok := req.Context().Value(ContentKey).(UserClaim)
  return claim, ok
}
