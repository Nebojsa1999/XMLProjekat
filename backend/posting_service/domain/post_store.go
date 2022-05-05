package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type PostStore interface {
	GetAll() ([]*Post, error)
	GetAllFromCollection(id primitive.ObjectID) ([]*Post, error)
	Get(id, post_id primitive.ObjectID) (*Post, error)
	Insert(id primitive.ObjectID, post *Post) (*Post, error)
	InsertComment(id primitive.ObjectID, post_id primitive.ObjectID, comment *Comment) (*Comment, error)
	UpdateLikes(likeordislike *LikeOrDislike) (*Post, error)
	UpdateDislikes(likeordislike *LikeOrDislike) (*Post, error)
	DeleteAll()
}
