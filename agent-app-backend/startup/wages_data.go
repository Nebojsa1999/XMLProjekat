package startup

import (
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/domain"
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/domain/enums"
)

var wages = []*domain.Wage{
	{
		Id:              getObjectId("62bdbf68457873bd3b14a1b6"),
		UserId:          getObjectId("123a0cc3a34d25d8567f9f01"),
		JobId:           getObjectId("62bdba74ab73b862988b9fa5"),
		Position:        enums.DevOps,
		Engagement:      enums.FullTime,
		ExperienceLevel: enums.Junior,
		NetoWage:        "550",
	},
}
