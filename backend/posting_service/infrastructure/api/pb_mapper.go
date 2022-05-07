package api

import (
	pb "github.com/Nebojsa1999/XMLProjekat/backend/common/proto/posting_service"
	"github.com/Nebojsa1999/XMLProjekat/backend/posting_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapPost(post *domain.Post) *pb.Post {
	postPb := &pb.Post{
		Id:            post.Id.Hex(),
		Content:       post.Content,
		Image:         post.Image,
		LikesCount:    post.LikesCount,
		DislikesCount: post.DislikesCount,
		Link:          post.Link,
	}

	for _, comment := range post.Comments {
		postPb.Comments = append(postPb.Comments, &pb.Comment{
			Code:    comment.Code,
			Content: comment.Content,
		})
	}

	return postPb
}

func mapPostRequest(postPb *pb.Post) *domain.Post {
	id, _ := primitive.ObjectIDFromHex(postPb.Id)
	ownerId, _ := primitive.ObjectIDFromHex(postPb.OwnerId)
	Post := &domain.Post{
		Id:            id,
		OwnerId:       ownerId,
		Content:       postPb.Content,
		Image:         postPb.Image,
		LikesCount:    postPb.LikesCount,
		DislikesCount: postPb.DislikesCount,
		Comments:      make([]domain.Comment, 0),
		Link:          postPb.Link,
	}

	for _, commentPb := range postPb.Comments {
		comment := domain.Comment{
			Code:    commentPb.Code,
			Content: commentPb.Content,
		}

		Post.Comments = append(Post.Comments, comment)
	}

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
