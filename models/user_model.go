package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserModel struct{
	ID primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name string `bson:"name" json:"name"`
	GithubUsername string `bson:"github_username" json:"github_username"`
	Email string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
	Bio string `bson:"bio" json:"bio,omitempty"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
	ProfileUrl string `bson:"profile_url" json:"profile_url"`
}