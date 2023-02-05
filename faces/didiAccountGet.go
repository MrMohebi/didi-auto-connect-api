package faces

import "go.mongodb.org/mongo-driver/bson/primitive"

type DidiAccountGetRes struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	Username  string             `json:"username"`
	Password  string             `json:"password"`
	CreatedAt int64              `json:"createdAt"`
	UpdatedAt int64              `json:"updatedAt"`
}
