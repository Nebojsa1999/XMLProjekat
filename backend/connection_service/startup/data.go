package startup

import (
	"github.com/Nebojsa1999/XMLProjekat/backend/connection_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var connections = []*domain.Connection{
	{
		Id:      getObjectId("62706d1b624b3da748f63fe1"),
		UserAId: getObjectId("62706d1b624b3da748f63fe3"),
		UserBId: getObjectId("62706d1b624b3da748f63fe5"),
	},
}

var privacy = []*domain.ProfilePrivacy{
	{
		Id:        primitive.NewObjectID(),
		UserId:    getObjectId("62706d1b624b3da748f63fe3"),
		IsPrivate: false,
	},
	{
		Id:        primitive.NewObjectID(),
		UserId:    getObjectId("62706d1b624b3da748f63fe5"),
		IsPrivate: false,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
