package gateway

import "github.com/pkg/errors"

func getUserFromAuthToken(authToken string) (string, error) {
	//todo should get user token from auth logic server

	return authToken, nil

	if authToken == "001" {
		// this token is user identification
		token := "kuip-gaosheng"
		return token, nil
	}

	return "", errors.New("authToken is invalid")
}
