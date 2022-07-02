package startup

import (
	"github.com/Nebojsa1999/XMLProjekat/backend/connection_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var connections = []*domain.Connection{
	{
		Id:         getObjectId("62706d1b624b3da748f63fe1"),
		IssuerId:   getObjectId("623b0cc3a34d25d8567f9f85"),
		SubjectId:  getObjectId("623b0cc3a34d25d8567f9f82"),
		IsApproved: true,
		Date:       time.Now(),
	},
	{
		Id:         getObjectId("62706c1b624b3da748f63fe2"),
		IssuerId:   getObjectId("623b0cc3a34d25d8567f9f82"),
		SubjectId:  getObjectId("623b0cc3a34d25d8567f9f83"),
		IsApproved: true,
		Date:       time.Now(),
	},
	{
		Id:         getObjectId("62706c1b624b3da748f63fe3"),
		IssuerId:   getObjectId("623b0cc3a34d25d8567f9f83"),
		SubjectId:  getObjectId("623b0cc3a34d25d8567f9f82"),
		IsApproved: true,
		Date:       time.Now(),
	},
}

var profilesPrivacy = []*domain.ProfilePrivacy{
	{
		Id:        primitive.NewObjectID(),
		UserId:    getObjectId("623b0cc3a34d25d8567f9f82"),
		IsPrivate: true,
	},
	{
		Id:        primitive.NewObjectID(),
		UserId:    getObjectId("623b0cc3a34d25d8567f9f83"),
		IsPrivate: true,
	},
	{
		Id:        primitive.NewObjectID(),
		UserId:    getObjectId("623b0cc3a34d25d8567f9f84"),
		IsPrivate: false,
	},
	{
		Id:        primitive.NewObjectID(),
		UserId:    getObjectId("623b0cc3a34d25d8567f9f85"),
		IsPrivate: false,
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}

	return primitive.NewObjectID()
}
