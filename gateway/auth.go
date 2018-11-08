package gateway

func getLinkKeyFromToken(token string) (string, error) {
	//todo
	//mock get user token from auth logic server
	tokenKey := map[string]string{
		"001":"001",
		"002":"002",
	}

	return tokenKey[token],nil
}
