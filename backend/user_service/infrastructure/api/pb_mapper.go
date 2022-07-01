package api

import (
	pb "github.com/Nebojsa1999/XMLProjekat/backend/common/proto/user_service"
	"github.com/Nebojsa1999/XMLProjekat/backend/user_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapDomainUserToPbUser(user *domain.User) *pb.User {
	userPb := &pb.User{
		Id:                user.Id.Hex(),
		Role:              mapDomainRoleToPbRole(user.Role),
		Username:          user.Username,
		Password:          user.Password,
		IsPrivate:         user.IsPrivate,
		FirstName:         user.FirstName,
		LastName:          user.LastName,
		Email:             user.Email,
		Phone:             user.Phone,
		Gender:            mapDomainGenderToPbGender(user.Gender),
		DateOfBirth:       timestamppb.New(user.DateOfBirth),
		Biography:         user.Biography,
		WorkExperience:    user.WorkExperience,
		Education:         user.Education,
		Skills:            user.Skills,
		Interests:         user.Interests,
		JobOffersAPIToken: user.JobOffersAPIToken,
	}

	return userPb
}

func mapPbUserToDomainUser(userPb *pb.User) *domain.User {
	var id primitive.ObjectID
	if objectId, err := primitive.ObjectIDFromHex(userPb.Id); err == nil {
		id = objectId
	} else {
		id = primitive.NewObjectID()
	}

	user := &domain.User{
		Id:                id,
		Role:              mapPbRoleToDomainRole(userPb.Role),
		Username:          userPb.Username,
		Password:          userPb.Password,
		IsPrivate:         userPb.IsPrivate,
		FirstName:         userPb.FirstName,
		LastName:          userPb.LastName,
		Email:             userPb.Email,
		Phone:             userPb.Phone,
		Gender:            mapPbGenderToDomainGender(userPb.Gender),
		DateOfBirth:       userPb.DateOfBirth.AsTime(),
		Biography:         userPb.Biography,
		WorkExperience:    userPb.WorkExperience,
		Education:         userPb.Education,
		Skills:            userPb.Skills,
		Interests:         userPb.Interests,
		JobOffersAPIToken: userPb.JobOffersAPIToken,
	}

	return user
}

func mapDomainRoleToPbRole(role domain.Role) pb.User_Role {
	if role == domain.CommonUser {
		return pb.User_CommonUser
	} else {
		return pb.User_Administrator
	}
}

func mapPbRoleToDomainRole(role pb.User_Role) domain.Role {
	if role == pb.User_CommonUser {
		return domain.CommonUser
	} else {
		return domain.Administrator
	}
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

func mapDomainCredentialsToPbCredentials(credentials *domain.Credentials) *pb.Credentials {
	credentialsPb := &pb.Credentials{
		Username: credentials.Username,
		Password: credentials.Password,
	}

	return credentialsPb
}

func mapPbCredentialsToDomainCredentials(credentialsPb *pb.Credentials) *domain.Credentials {
	credentials := &domain.Credentials{
		Username: credentialsPb.Username,
		Password: credentialsPb.Password,
	}

	return credentials
}
