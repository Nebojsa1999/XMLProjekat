package api

import (
	pb "github.com/Nebojsa1999/XMLProjekat/backend/common/proto/user_service"
	"github.com/Nebojsa1999/XMLProjekat/backend/user_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapDomainUserToPbUser(user *domain.User) *pb.User {
	userPb := &pb.User{
		Id:             user.Id.Hex(),
		Username:       user.Username,
		Password:       user.Password,
		IsPrivate:      user.IsPrivate,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Email:          user.Email,
		Gender:         mapDomainGenderToPbGender(user.Gender),
		DateOfBirth:    timestamppb.New(user.DateOfBirth),
		Biography:      user.Biography,
		WorkExperience: user.WorkExperience,
		Education:      user.Education,
		Skills:         user.Skills,
		Interests:      user.Interests,
	}

	return userPb
}

func mapPbUserToDomainUser(userPb *pb.User) *domain.User {
	newUser := &domain.User{
		Id:             primitive.NewObjectID(),
		Username:       userPb.Username,
		Password:       userPb.Password,
		IsPrivate:      userPb.IsPrivate,
		FirstName:      userPb.FirstName,
		LastName:       userPb.LastName,
		Email:          userPb.Email,
		Gender:         mapPbGenderToDomainGender(userPb.Gender),
		DateOfBirth:    userPb.DateOfBirth.AsTime(),
		Biography:      userPb.Biography,
		WorkExperience: userPb.WorkExperience,
		Education:      userPb.Education,
		Skills:         userPb.Skills,
		Interests:      userPb.Interests,
	}

	return newUser
}

func mapDomainGenderToPbGender(gender domain.Gender) pb.User_Gender {
	if gender == domain.Male {
		return pb.User_Male
	} else {
		return pb.User_Female
	}
}

func mapPbGenderToDomainGender(gender pb.User_Gender) domain.Gender {
	if gender == pb.User_Male {
		return domain.Male
	} else {
		return domain.Female
	}
}
