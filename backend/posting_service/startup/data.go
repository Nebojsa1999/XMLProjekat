package startup

import (
	"github.com/Nebojsa1999/XMLProjekat/backend/posting_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
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
		User: domain.User{
			Id:                getObjectId("623b0cc3a34d25d8567f9f84"),
			Role:              domain.CommonUser,
			Username:          "Aleksandar97",
			Password:          "aleksandar",
			IsPrivate:         false,
			FirstName:         "Aleksandar",
			LastName:          "Dujin",
			Email:             "aleksandar.dujin@gmail.com",
			Phone:             "060/132-345",
			Gender:            domain.Male,
			DateOfBirth:       getParsedDateOfBirthFrom("1997-02-03T00:00:00Z"),
			Biography:         "Biografija Aleksandra Dujina.",
			WorkExperience:    "Bez radnog iskustva.",
			Education:         "Fakultet tehničkih nauka Novi Sad",
			Skills:            "Veb programiranje.",
			Interests:         "Video igre.",
			JobOffersAPIToken: "",
		},
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
		User: domain.User{
			Id:                getObjectId("623b0cc3a34d25d8567f9f84"),
			Role:              domain.CommonUser,
			Username:          "Aleksandar97",
			Password:          "aleksandar",
			IsPrivate:         false,
			FirstName:         "Aleksandar",
			LastName:          "Dujin",
			Email:             "aleksandar.dujin@gmail.com",
			Phone:             "060/132-345",
			Gender:            domain.Male,
			DateOfBirth:       getParsedDateOfBirthFrom("1997-02-03T00:00:00Z"),
			Biography:         "Biografija Aleksandra Dujina.",
			WorkExperience:    "Bez radnog iskustva.",
			Education:         "Fakultet tehničkih nauka Novi Sad",
			Skills:            "Veb programiranje.",
			Interests:         "Video igre.",
			JobOffersAPIToken: "",
		},
	},
	{
		Id:      getObjectId("623b0cc3a34d25d8567f9f84"),
		OwnerId: getObjectId("623b0cc3a34d25d8567f9f85"),
		Content: "Zdravo, ja sam Marko Trifunović! Ovo je moja prva objava.",
		Image:   "programiranje3.jpg",
		User: domain.User{
			Id:                getObjectId("623b0cc3a34d25d8567f9f85"),
			Role:              domain.Administrator,
			Username:          "Marko99",
			Password:          "marko99",
			IsPrivate:         false,
			FirstName:         "Marko",
			LastName:          "Trifunović",
			Email:             "marko.trifunovic@gmail.com",
			Phone:             "063/763-6897",
			Gender:            domain.Male,
			DateOfBirth:       getParsedDateOfBirthFrom("1999-11-30T00:00:00Z"),
			Biography:         "Biografija Marka Trifunovića.",
			WorkExperience:    "Bez radnog iskustva.",
			Education:         "Fakultet tehničkih nauka Novi Sad",
			Skills:            "Programer WPF aplikacija.",
			Interests:         "Video igre.",
			JobOffersAPIToken: "",
		},
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
