package startup

import (
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/domain"
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/domain/enums"
)

var wages = []*domain.Wage{
	{
		Id:              getObjectId("62bdbf68457873bd3b14a1b6"),
		CompanyId:       getObjectId("123a0cc3a34d25d8567f9f01"),
		Position:        enums.DevOps,
		Engagement:      enums.FullTime,
		ExperienceLevel: enums.Junior,
		NetoWage:        "550",
	},
}
