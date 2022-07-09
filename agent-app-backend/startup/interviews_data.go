package startup

import (
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/domain"
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/domain/enums"
)

var interviews = []*domain.Interview{
	{
		Id:                 getObjectId("62bdc05639c6e9c9d3b2757f"),
		CompanyId:          getObjectId("123a0cc3a34d25d8567f9f01"),
		Position:           enums.DevOps,
		Title:              "Jako dobar regruterski tim",
		YearOfInterview:    "2022",
		HRInterview:        "Vrlo opušten i prijatan intervju o prethodnim iskustvima i generalnim očekivanjima. Stekla sam utisak da se vodi računa o potrebama zaposlenog. Nakon par dana je usledio tehnički intervju.",
		TechnicalInterview: "Isto tako opušten i prijatan razgovor. Pitanja su mahom teorijska, neka su fokusirana na diskusiju o nekim rešenjima i na kraju postoji mali case study koji je usmeren ka tome da se proveri kako osoba razmišlja. Nema live codinga, niti nekog zadatka pre toga (što ubrzava proces selekcije).",
	},
}
