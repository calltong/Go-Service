package authentication

import (
  "fmt"
  "time"
  "gopkg.in/mgo.v2/bson"
  jwt "gopkg.in/dgrijalva/jwt-go.v3"
)

const signatureKey string = "aventador"

type ProjectGeneral struct {
	Folder string `json:"folder" bson:"folder"`
  Database string `json:"db_name" bson:"db_name"`
  Address string `json:"address" bson:"address"`
}

type UserClaim struct {
    User bson.ObjectId `json:"user"`
    Project ProjectGeneral `json:"project"`
    // recommended having
    jwt.StandardClaims
}

func CreateToken(id bson.ObjectId, project ProjectGeneral) string {
  // Expires the token and cookie in 1 hour
  expireToken := time.Now().AddDate(1, 0, 0).Unix()
  // We'll manually assign the claims but in production you'd insert values from a database
  claims := UserClaim {
    id,
    project,
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
