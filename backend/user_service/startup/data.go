package startup

import (
	"github.com/Nebojsa1999/XMLProjekat/backend/user_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var users = []*domain.User{
	{
		Id: getObjectId("623b0cc3a34d25d8567f9f82"),
		Username: "Darijan98",
		Password: "darijan",
		IsPrivate: true,
		FirstName: "Darijan",
		LastName: "Mićić",
		Email: "darijan.micic10@gmail.com",
		Gender: domain.Male,
		DateOfBirth: getParsedDateOfBirthFrom("1998-07-10"),
		Biography: "Biografija Darijana Mićića.",
		WorkExperience: "Bez radnog iskustva.",
		Education: "Fakultet tehničkih nauka Novi Sad",
		Skills: "Veb programiranje.",
		Interests: "Video igre.",
	},
	{
		Id: getObjectId("623b0cc3a34d25d8567f9f83"),
		Username: "Nebojsa99",
		Password: "nebojsa",
		IsPrivate: true,
		FirstName: "Nebojša",
		LastName: "Bogosavljev",
		Email: "nebojsa.bogosavljev@gmail.com",
		Gender: domain.Male,
		DateOfBirth: getParsedDateOfBirthFrom("1999-09-26"),
		Biography: "Biografija Nebojše Bogosavljeva.",
		WorkExperience: "Bez radnog iskustva",
		Education: "Fakultet tehničkih nauka Novi Sad",
		Skills: "Programer WPF aplikacija.",
		Interests: "Video igre.",
	},
	{
		Id: getObjectId("623b0cc3a34d25d8567f9f84"),
		Username: "Aleksandar97",
		Password: "aleksandar",
		IsPrivate: false,
		FirstName: "Aleksandar",
		LastName: "Dujin",
		Email: "aleksandar.dujin@gmail.com",
		Gender: domain.Male,
		DateOfBirth: getParsedDateOfBirthFrom("1997-02-03"),
		Biography: "Biografija Aleksandra Dujina.",
		WorkExperience: "Bez radnog iskustva.",
		Education: "Fakultet tehničkih nauka Novi Sad",
		Skills: "Veb programiranje.",
		Interests: "Video igre.",
	},
	{
		Id: getObjectId("623b0cc3a34d25d8567f9f85"),
		Username: "Marko99",
		Password: "marko",
		IsPrivate: false,
		FirstName: "Marko",
		LastName: "Trifunović",
		Email: "marko.trifunovic@gmail.com",
		Gender: domain.Male,
		DateOfBirth: getParsedDateOfBirthFrom("1999-11-30"),
		Biography: "Biografija Marka Trifunovića.",
		WorkExperience: "Bez radnog iskustva.",
		Education: "Fakultet tehničkih nauka Novi Sad",
		Skills: "Programer WPF aplikacija.",
		Interests: "Video igre.",
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}

	return primitive.NewObjectID()
}

func getParsedDateOfBirthFrom(dateOfBirthAsString string) time.Time {
	dateOfBirth, _ := time.Parse("2010-01-30", dateOfBirthAsString)

	return dateOfBirth
}
