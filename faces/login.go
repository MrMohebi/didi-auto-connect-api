package faces

type LoginRes struct {
	Token     string `json:"token"`
	HasAccess bool   `json:"hasAccess"`
	IsLimit   bool   `json:"isLimit"`
}

type LoginReq struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	DeviceHash string `json:"deviceHash"`
}
