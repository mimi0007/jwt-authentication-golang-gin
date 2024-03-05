package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type user struct {
	ID           primitive.ObjectID `bson:_id`
	userId       string             `json:"user_id" validate:"required"`
	firstName    string             `json:"first_name" validate:"required min=2 max=10"`
	lastName     string             `json:"last_name" validate:"required min=2 max=10"`
	email        string             `json:"email" validate:"required, email`
	phone        string             `json:"phone" validate:"required"`
	password     string             `json:"password" validate:"required min=6"`
	createAt     time.Time          `json:"created_at"`
	updatedAt    time.Time          `json:"updated_at"`
	deletedAt    time.Time          `json:"deleted_at"`
	accessToken  string             `json:"access_token" validate:"required`
	refreshToken string             `json:"refresh_token" validate:"required`
	userType     string             `json:"user_type" validate:"required eq = ADMIN | USER`
}
