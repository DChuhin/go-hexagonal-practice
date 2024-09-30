package mongodbiotapplication

import "go-hexagonal-practice/internal/domain"

type Document struct {
	Id          string              `bson:"_id,omitempty"`
	Name        string              `bson:"name,omitempty"`
	UserId      string              `bson:"userId,omitempty"`
	Credentials *domain.Credentials `bson:"credentials,omitempty"`
}
