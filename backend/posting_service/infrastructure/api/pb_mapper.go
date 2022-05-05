package api

import "go.mongodb.org/mongo-driver/bson/primitive"

func mapPost(post *domain.Post) *pb.Post {
	postPb := &pb.Post{
		Id:          post.Id.Hex(),
		Title:       post.Title,
		DateCreated: post.DateCreated,
	}
	return postPb
}

func mapInsertPost(post *pb.Post) *domain.Post {
	id, _ := primitive.ObjectIDFromHex(post.Id)

	postPb := &domain.Post{
		Id:          id,
		Title:       post.Title,
		DateCreated: post.DateCreated,
	}

	return postPb
}
