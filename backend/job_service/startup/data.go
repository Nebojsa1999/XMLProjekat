package startup

import (
	"github.com/Nebojsa1999/XMLProjekat/backend/job_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var jobs = []*domain.Job{
	{
		Id:           getObjectId("623b0cc3a34d25d8567f9f81"),
		UserId:       getObjectId("123a0cc3a34d25d8567f9f04"),
		CreatedAt:    getParsedDateOfCreationFrom("2022-06-01T09:25:00Z"),
		Position:     "Junior developer",
		Description:  "Programer prednjeg dela.",
		Requirements: "Diploma osnovnih akademskih studija fakulteta.",
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}

	return primitive.NewObjectID()
}

func getParsedDateOfCreationFrom(dateOfCreationAsString string) time.Time {
	dateOfCreation, _ := time.Parse(time.RFC3339, dateOfCreationAsString)

	return dateOfCreation
}
