package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	FirstName string             `json:"first_name" validate:"required, min=2, max=100" bson:"first_name"`
	LastName  string             `json:"last_name" validate:"required, min=2, max=100" bson:"last_name"`
	Password  string             `json:"password" validate:"required, min=6" bson:"password"`
	Email     string             `json:"email" validate:"email, required" bson:"email"`
	Phone     string             `json:"phone" validate:"required" bson:"phone"`
	Token     string             `json:"token" bson:"token"`
	// UserType			*string					`json:"user_type" validate:"required, eq=ADMIN|eq=USER"`
	// RefreshToken		*string					`json:"refresh_token"`
	// CreatedAt			time.Time				`json:"created_at"`
	// UpdatedAt			time.Time				`json:"updated_at"`
	// UserID				*string					`json:"user_id"`
}

type CreateUser struct {
	FirstName string `json:"first_name" validate:"required, min=2, max=100" bson:"first_name"`
	LastName  string `json:"last_name" validate:"required, min=2, max=100" bson:"last_name"`
	Password  string `json:"password" validate:"required, min=6" bson:"password"`
	Email     string `json:"email" validate:"email, required" bson:"email"`
	Phone     string `json:"phone" validate:"required" bson:"phone"`
}
