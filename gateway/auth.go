package gateway

func getLinkKeyFromToken(token string) (string, error) {
	//todo
	//mock get user token from auth logic server
	tokenKey := map[string]string{
		"001": "key1",
		"002": "key2",
	}

	return tokenKey[token], nil
}
