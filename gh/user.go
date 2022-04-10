package ghpr

func Me() (me string) {
	var client = getRestClient()
	var response User
	client.Get("user", &response)
	return response.Login
}
