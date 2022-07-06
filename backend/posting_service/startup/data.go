package startup

import (
	"github.com/Nebojsa1999/XMLProjekat/backend/posting_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var posts = []domain.Post{
	{
		Id:            getObjectId("623b0cc3a34d25d8567f9f82"),
		OwnerId:       getObjectId("623b0cc3a34d25d8567f9f84"),
		Content:       "Zdravo! Ja sam Aleksandar Dujin, ovo je moja prva objava.",
		LikesCount:    1,
		DislikesCount: 0,
		WhoLiked: []string{
			"623b0cc3a34d25d8567f9f82",
		},
		WhoDisliked: []string{},
		Comments: []domain.Comment{
			{
				Code:    "1",
				Content: "Odlican post",
			},
		},
		Image: "programiranje1.jpg",
	},
	{
		Id:            getObjectId("623b0cc3a34d25d8567f9f83"),
		OwnerId:       getObjectId("623b0cc3a34d25d8567f9f84"),
		Content:       "Druga objava Aleksandra Dujina.",
		LikesCount:    0,
		DislikesCount: 1,
		WhoLiked:      []string{},
		WhoDisliked: []string{
			"623b0cc3a34d25d8567f9f82",
		},
		Comments: []domain.Comment{
			{
				Code:    "2",
				Content: "Sjajno",
			},
			{
				Code:    "3",
				Content: "Sjajno jako",
			},
		},
		Image: "programiranje2.png",
	},
	{
		Id:      getObjectId("623b0cc3a34d25d8567f9f84"),
		OwnerId: getObjectId("623b0cc3a34d25d8567f9f85"),
		Content: "Zdravo, ja sam Marko Trifunović! Ovo je moja prva objava.",
		Image:   "programiranje3.jpg",
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
