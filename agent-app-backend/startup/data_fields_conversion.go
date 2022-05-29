package startup

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}

	return primitive.NewObjectID()
}

func getParsedDateOfBirthFrom(dateOfBirthAsString string) time.Time {
	dateOfBirth, _ := time.Parse(time.RFC3339, dateOfBirthAsString)

	return dateOfBirth
}
