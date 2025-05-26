package webexgosdk


type webexApi struct {
	accessToken string
}

func NewWebexApi(accessToken string) *webexApi {
	return &webexApi{
		accessToken: accessToken,
	}
}