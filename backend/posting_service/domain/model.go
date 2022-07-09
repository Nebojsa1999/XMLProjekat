package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Post struct {
	Id            primitive.ObjectID `bson:"_id"`
	OwnerId       primitive.ObjectID `bson:"owner_Id"`
	Content       string             `bson:"content"`
	Image         string             `bson:"image"`
	LikesCount    int64              `bson:"likes"`
	DislikesCount int64              `bson:"dislikes"`
	Comments      []Comment          `bson:"comments"`
	Link          []string           `bson:"link"`
	WhoLiked      []string           `bson:"liked_by"`
	WhoDisliked   []string           `bson:"disliked_by"`
	User          User               `bson:"user"`
	PostedAt      primitive.DateTime `bson:"posted_at"`
}
type Comment struct {
	Code    string `bson:"code"`
	Content string `bson:"content"`
}

type LikeOrDislike struct {
	Id                primitive.ObjectID `bson:"_id"`
	PostId            primitive.ObjectID `bson:"post_id"`
	LikedOrDislikedBy primitive.ObjectID `bson:"liked_or_disliked_by"`
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

type User struct {
	Id                primitive.ObjectID `json:"id" bson:"_id"`
	Role              Role               `json:"role" bson:"role"`
	Username          string             `json:"username" bson:"username"`
	Password          string             `json:"password" bson:"password"`
	IsPrivate         bool               `json:"is_private" bson:"is_private"`
	FirstName         string             `json:"first_name" bson:"first_name"`
	LastName          string             `json:"last_name" bson:"last_name"`
	Email             string             `json:"email" bson:"email"`
	Phone             string             `json:"phone" bson:"phone"`
	Gender            Gender             `json:"gender" bson:"gender"`
	DateOfBirth       time.Time          `json:"date_of_birth" bson:"date_of_birth"`
	Biography         string             `json:"biography" bson:"biography"`
	WorkExperience    string             `json:"work_experience" bson:"work_experience"`
	Education         string             `json:"education" bson:"education"`
	Skills            string             `json:"skills" bson:"skills"`
	Interests         string             `json:"interests" bson:"interests"`
	JobOffersAPIToken string             `json:"jobOffersAPIToken" bson:"job_offers_api_token"`
}

type Role string

const (
	UndefinedRole Role = ""
	CommonUser         = "CommonUser"
	Administrator      = "Administrator"
)

type Gender string

const (
	UndefinedGender Gender = ""
	Male                   = "Male"
	Female                 = "Female"
)
