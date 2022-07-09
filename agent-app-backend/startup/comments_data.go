package startup

import (
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/domain"
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/domain/enums"
)

var comments = []*domain.Comment{
	{
		Id:              getObjectId("62bdbe44b55d3b5c0fa891e1"),
		CompanyId:       getObjectId("123a0cc3a34d25d8567f9f01"),
		Position:        enums.DevOps,
		Engagement:      enums.FullTime,
		ExperienceLevel: enums.Junior,
		Content: "Puno prostora za ucenje i napredovanje, uz adekvatnu podrsku. \n\n" +
			"Plate su prilicno konkurentne a projekti na kojima kolege rade deluju vrlo izazovno i interesantno.",
	},
}
