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
		Image:         "",
		LikesCount:    0,
		DislikesCount: 0,
		Comments:      nil,
		Link:          "",
	},
	{
		Id:            getObjectId("623b0cc3a34d25d8567f9f83"),
		OwnerId:       getObjectId("623b0cc3a34d25d8567f9f84"),
		Content:       "Druga objava Aleksandra Dujina.",
		Image:         "",
		LikesCount:    0,
		DislikesCount: 0,
		Comments:      nil,
		Link:          "",
	},
	{
		Id:            getObjectId("623b0cc3a34d25d8567f9f84"),
		OwnerId:       getObjectId("623b0cc3a34d25d8567f9f85"),
		Content:       "Zdravo, ja sam Marko Trifunović! Ovo je moja prva objava.",
		Image:         "",
		LikesCount:    0,
		DislikesCount: 0,
		Comments:      nil,
		Link:          "",
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
