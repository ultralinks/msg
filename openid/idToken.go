package openid

import (
    "errors"
    jwt "github.com/dgrijalva/jwt-go"
)

// https://github.com/dgrijalva/jwt-go/blob/master/claims.go
// https://github.com/dgrijalva/jwt-go/blob/0b96aaa707760d6ab28d9b9d1913ff5993328bae/example_test.go

type IdTokenClaims struct {
    AuthTime    int64   `json:"auth_time,omitempty"`
    Nonce       string  `json:"nonce,omitempty"`
    Acr         string  `json:"acr,omitempty"`          // Authentication Context Class
    Arm         string  `json:"arm,omitempty"`          // Authentication Methods References
    Azp         string  `json:"azp,omitempty"`          // Authorized party
    jwt.StandardClaims
}

func ParseIdToken(idTokenStr string, publicKey string) (*IdTokenClaims, error) {
    verifyKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
    if err != nil {
        return nil, err
    }

    token, err := jwt.ParseWithClaims(idTokenStr, &IdTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
        return verifyKey, nil
    })
    if err != nil {
        return nil, err
    }

    idToken, ok := token.Claims.(*IdTokenClaims)

    if ok && token.Valid {
        return idToken, nil
    }

    return idToken, errors.New("valid failed")
}
