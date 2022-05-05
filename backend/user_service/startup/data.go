package startup

import (
	"github.com/Nebojsa1999/XMLProjekat/backend/user_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var users = []*domain.User{
	{
		Id: getObjectId("623b0cc3a34d25d8567f9f82"),
		Username: "Darijan98",
		Password: getHashOfThePassword("malaMaca9"),
		FirstName: "Darijan",
		LastName: "Mićić",
		Email: "darijan.micic10@gmail.com",
		Gender: domain.Gender("Male"),
		DateOfBirth: getParsedDateOfBirthFrom("1998-07-10"),
		Biography: "Biografija Mićića.",
		WorkExperience: "Radno iskustvo Darijana Mićića.",
		Education: "Obrazovanje Darijana Mićića.",
		Skills: "Veštine Darijana Mićića.",
		Interests: "Interesovanja Darijana Mićića.",
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}

	return primitive.NewObjectID()
}

func getHashOfThePassword(password string) []byte {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	return hash
}

func getParsedDateOfBirthFrom(dateOfBirthAsString string) time.Time {
	dateOfBirth, _ := time.Parse("2010-01-30", dateOfBirthAsString)

	return dateOfBirth
}
