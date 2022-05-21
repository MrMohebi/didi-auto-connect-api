package faces

type LoginRes struct {
	Token     string `json:"token"`
	HasAccess bool   `json:"hasAccess"`
	IsLimit   bool   `json:"isLimit"`
}

type LoginReq struct {
	Username   string `json:"username" form:"username" validate:"required"`
	Password   string `json:"password" form:"password" validate:"required"`
	DeviceHash string `json:"deviceHash" form:"deviceHash" validate:"required"`
}
