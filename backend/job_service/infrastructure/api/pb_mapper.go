package api

import (
	pb "github.com/Nebojsa1999/XMLProjekat/backend/common/proto/job_service"
	"github.com/Nebojsa1999/XMLProjekat/backend/job_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapJob(job *domain.Job) *pb.Job {
	jobPb := &pb.Job{
		Id:           job.Id.Hex(),
		UserId:       job.UserId.Hex(),
		CreatedAt:    timestamppb.New(job.CreatedAt),
		Requirements: job.Requirements,
		Description:  job.Description,
		Position:     job.Position,
	}

	return jobPb
}

func mapNewJob(jobPb *pb.Job) *domain.Job {
	var id primitive.ObjectID
	if objectId, err := primitive.ObjectIDFromHex(jobPb.Id); err == nil {
		id = objectId
	} else {
		id = primitive.NewObjectID()
	}

	userId, _ := primitive.ObjectIDFromHex(jobPb.UserId)

	job := &domain.Job{
		Id:           id,
		UserId:       userId,
		CreatedAt:    jobPb.CreatedAt.AsTime(),
		Requirements: jobPb.Requirements,
		Description:  jobPb.Description,
		Position:     jobPb.Position,
	}

	return job
}

func mapChangesOfJob(jobPb *pb.Job) *domain.Job {
	id, _ := primitive.ObjectIDFromHex(jobPb.Id)
	userId, _ := primitive.ObjectIDFromHex(jobPb.UserId)

	job := &domain.Job{
		Id:           id,
		UserId:       userId,
		CreatedAt:    jobPb.CreatedAt.AsTime(),
		Requirements: jobPb.Requirements,
		Description:  jobPb.Description,
		Position:     jobPb.Position,
	}
	return job
}
