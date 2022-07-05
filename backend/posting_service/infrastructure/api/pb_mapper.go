package api

import (
	pb "github.com/Nebojsa1999/XMLProjekat/backend/common/proto/posting_service"
	"github.com/Nebojsa1999/XMLProjekat/backend/posting_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func mapPost(post *domain.Post) *pb.Post {
	postPb := &pb.Post{
		Id:            post.Id.Hex(),
		OwnerId:       post.OwnerId.Hex(),
		Content:       post.Content,
		Image:         post.Image,
		LikesCount:    post.LikesCount,
		DislikesCount: post.DislikesCount,
		Link:          post.Link,
		WhoLiked:      post.WhoLiked,
		WhoDisliked:   post.WhoDisliked,
	}

	for _, comment := range post.Comments {
		postPb.Comments = append(postPb.Comments, &pb.Comment{
			Code:    comment.Code,
			Content: comment.Content,
		})
	}
	postPb.PostedAt = post.PostedAt.Time().String()
	return postPb
}

func mapPostRequest(postPb *pb.Post) *domain.Post {
	id := getObjectId(postPb.Id)
	ownerId := getObjectId(postPb.OwnerId)
	Post := &domain.Post{
		Id:            id,
		OwnerId:       ownerId,
		Content:       postPb.Content,
		Image:         postPb.Image,
		LikesCount:    postPb.LikesCount,
		DislikesCount: postPb.DislikesCount,
		Comments:      make([]domain.Comment, 0),
		Link:          postPb.Link,
		WhoLiked:      postPb.WhoLiked,
		WhoDisliked:   postPb.WhoDisliked,
	}

	for _, commentPb := range postPb.Comments {
		comment := domain.Comment{
			Code:    commentPb.Code,
			Content: commentPb.Content,
		}

		Post.Comments = append(Post.Comments, comment)
	}
	t, _ := time.Parse(time.RFC3339, postPb.PostedAt)
	Post.PostedAt = primitive.NewDateTimeFromTime(t)

	return Post
}

func mapComment(commentPb *domain.Comment) *pb.Comment {
	Comment := &pb.Comment{
		Code:    commentPb.Code,
		Content: commentPb.Content,
	}

	return Comment
}

func mapCommentOnPostRequest(commentPb *pb.Comment) *domain.Comment {
	Comment := &domain.Comment{
		Code:    commentPb.Code,
		Content: commentPb.Content,
	}

	return Comment
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}

	return primitive.NewObjectID()
}
