package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type PostStore interface {
	GetAllPosts([]string) ([]*Post, error)
	GetAllPostsFromUser(id primitive.ObjectID) ([]*Post, error)
	GetPostFromUser(id, post_id primitive.ObjectID) (*Post, error)
	UpdateLikes(liked_or_disliked_by *LikeOrDislike) (*Post, error)
	UpdateDislikes(liked_or_disliked_by *LikeOrDislike) (*Post, error)
	CreatePost(id primitive.ObjectID, post *Post) (*Post, error)
	CreateComment(id primitive.ObjectID, post_id primitive.ObjectID, comment *Comment) (*Comment, error)
	DeleteAll()
}
