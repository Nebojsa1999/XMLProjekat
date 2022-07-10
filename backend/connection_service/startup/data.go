package startup

import (
	"github.com/Nebojsa1999/XMLProjekat/backend/connection_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var connections = []*domain.Connection{
	{
		Id:         getObjectId("62706d1b624b3da748f63fe1"),
		Type:       domain.Following,
		IssuerId:   getObjectId("623b0cc3a34d25d8567f9f85"),
		SubjectId:  getObjectId("623b0cc3a34d25d8567f9f82"),
		IsApproved: true,
		Date:       time.Now(),
	},
	{
		Id:         getObjectId("62706c1b624b3da748f63fe2"),
		Type:       domain.Following,
		IssuerId:   getObjectId("623b0cc3a34d25d8567f9f82"),
		SubjectId:  getObjectId("623b0cc3a34d25d8567f9f83"),
		IsApproved: true,
		Date:       time.Now(),
	},
	{
		Id:         getObjectId("62706c1b624b3da748f63fe3"),
		Type:       domain.Following,
		IssuerId:   getObjectId("623b0cc3a34d25d8567f9f83"),
		SubjectId:  getObjectId("623b0cc3a34d25d8567f9f82"),
		IsApproved: true,
		Date:       time.Now(),
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}

	return primitive.NewObjectID()
}
