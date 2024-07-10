package schemas

import "go.mongodb.org/mongo-driver/bson/primitive"

type BoardSchema struct {
	ID          primitive.ObjectID `bson:"_id"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	CreatedBy   int32              `bson:"created_by"`
	CreatedAt   int64              `bson:"created_at"`
	UpdatedAt   int64              `bson:"updated_at"`
}

const BoardCollection string = "boards"
