package faces

type LoginRes struct {
	HasAccess bool   `json:"hasAccess"`
	Token     string `json:"token"`
	Message   string `json:"message"`
	Link      string `json:"link"`
}

type LoginReq struct {
	Username   string `json:"username" form:"username" validate:"required"`
	Password   string `json:"password" form:"password" validate:"required"`
	DeviceHash string `json:"deviceHash" form:"deviceHash" validate:"required"`
}
