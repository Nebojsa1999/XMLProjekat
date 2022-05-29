package startup

import (
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var users = []*domain.User{
	{
		Id:             getObjectId("123a0cc3a34d25d8567f9f01"),
		Role:           domain.CommonUser,
		Username:       "Darijan98",
		Password:       "darijan",
		FirstName:      "Darijan",
		LastName:       "Mićić",
		Email:          "darijan.micic10@gmail.com",
		Phone:          "062/100-6031",
		Gender:         domain.Male,
		DateOfBirth:    getParsedDateOfBirthFrom("1998-07-10T00:00:00Z"),
		Biography:      "Biografija Darijana Mićića.",
		WorkExperience: "Bez radnog iskustva.",
		Education:      "Fakultet tehničkih nauka Novi Sad",
		Skills:         "Veb programiranje.",
		Interests:      "Video igre.",
	},
	{
		Id:             getObjectId("123a0cc3a34d25d8567f9f02"),
		Role:           domain.Administrator,
		Username:       "Nebojsa99",
		Password:       "nebojsa",
		FirstName:      "Nebojša",
		LastName:       "Bogosavljev",
		Email:          "nebojsa.bogosavljev@gmail.com",
		Phone:          "064/788-400",
		Gender:         domain.Male,
		DateOfBirth:    getParsedDateOfBirthFrom("1999-09-26T00:00:00Z"),
		Biography:      "Biografija Nebojše Bogosavljeva.",
		WorkExperience: "Bez radnog iskustva",
		Education:      "Fakultet tehničkih nauka Novi Sad",
		Skills:         "Programer WPF aplikacija.",
		Interests:      "Video igre.",
	},
	{
		Id:             getObjectId("123a0cc3a34d25d8567f9f03"),
		Role:           domain.CommonUser,
		Username:       "Aleksandar97",
		Password:       "aleksandar",
		FirstName:      "Aleksandar",
		LastName:       "Dujin",
		Email:          "aleksandar.dujin@gmail.com",
		Phone:          "060/132-345",
		Gender:         domain.Male,
		DateOfBirth:    getParsedDateOfBirthFrom("1997-02-03T00:00:00Z"),
		Biography:      "Biografija Aleksandra Dujina.",
		WorkExperience: "Bez radnog iskustva.",
		Education:      "Fakultet tehničkih nauka Novi Sad",
		Skills:         "Veb programiranje.",
		Interests:      "Video igre.",
	},
	{
		Id:             getObjectId("123a0cc3a34d25d8567f9f04"),
		Role:           domain.CompanyOwner,
		Username:       "Marko99",
		Password:       "marko",
		FirstName:      "Marko",
		LastName:       "Trifunović",
		Email:          "marko.trifunovic@gmail.com",
		Phone:          "063/763-6897",
		Gender:         domain.Male,
		DateOfBirth:    getParsedDateOfBirthFrom("1999-11-30T00:00:00Z"),
		Biography:      "Biografija Marka Trifunovića.",
		WorkExperience: "Bez radnog iskustva.",
		Education:      "Fakultet tehničkih nauka Novi Sad",
		Skills:         "Programer WPF aplikacija.",
		Interests:      "Video igre.",
	},
}

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
