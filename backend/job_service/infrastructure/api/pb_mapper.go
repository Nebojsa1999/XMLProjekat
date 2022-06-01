package api

import "time"

func mapJob(job *domain.Job) *pb.Job {
	jobPb := &pb.Job{
		Id:           job.Id.Hex(),
		UserId:       job.UserId.Hex(),
		Requirements: job.Requirements,
		Description:  job.Description,
		Position:     job.Position,
	}
	jobPb.CreatedAt = job.CreatedAt.Time().String()
	return jobPb
}

func mapNewJob(jobPb *pb.Job) *domain.Job {
	id, _ := primitive.ObjectIDFromHex(jobPb.UserId)
	job := &domain.Job{
		Id:           primitive.NewObjectID(),
		UserId:       id,
		Requirements: jobPb.Requirements,
		Description:  jobPb.Description,
		Position:     jobPb.Position,
	}
	t, _ := time.Parse(time.RFC3339, jobPb.CreatedAt)
	job.CreatedAt = primitive.NewDateTimeFromTime(t)
	return job
}
