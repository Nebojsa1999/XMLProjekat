package api

import (
	"github.com/Nebojsa1999/XMLProjekat/backend/user_service/domain"
	pb "github.com/Nebojsa1999/XMLProjekat/backend/common/proto/user_service"
)

func mapUser(user *domain.User) *pb.User {
	userPb := &pb.User{
		Id: user.Id.Hex(),
		Username: user.Username,
		Password: user.Password,
		FirstName: user.FirstName,
		LastName: user.LastName,
		Email: user.Email,
		Gender: user.Gender,
		DateOfBirth: user.DateOfBirth,
		Biography: user.Biography,
		WorkExperience: user.WorkExperience,
		Education: user.Education,
		Skills: user.Skills,
		Interests: user.Interests,
	}

	return userPb
}
