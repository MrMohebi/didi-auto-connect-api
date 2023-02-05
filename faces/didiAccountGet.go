package faces

type DidiAccountGetRes struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}
