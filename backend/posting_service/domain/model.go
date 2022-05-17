package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	Id            primitive.ObjectID `bson:"_id"`
	OwnerId       primitive.ObjectID `bson:"owner_Id"`
	Content       string             `bson:"content"`
	Image         string             `bson:"image"`
	LikesCount    int64              `bson:"likes"`
	DislikesCount int64              `bson:"dislikes"`
	Comments      []Comment          `bson:"comments"`
	Link          string             `bson:"link"`
	WhoLiked      []string           `bson:"liked_by"`
	WhoDisliked   []string           `bson:"disliked_by"`
	PostedAt      primitive.DateTime `bson:"posted_at"`
}
type Comment struct {
	Code    string `bson:"code"`
	Content string `bson:"content"`
}

type LikeOrDislike struct {
	Id                primitive.ObjectID `bson:"_id"`
	LikedOrDislikedBy primitive.ObjectID `bson:"liked_or_disliked_by"`
	PostId            primitive.ObjectID `bson:"post_id"`
}

type NewPostRequest struct {
	Id   primitive.ObjectID `bson:"_id"`
	Post Post               `bson:"post"`
}

type CommentOnPostRequest struct {
	Id      primitive.ObjectID `bson:"_id"`
	PostID  primitive.ObjectID `bson:"post_id"`
	Comment Comment            `bson:"comment"`
}
