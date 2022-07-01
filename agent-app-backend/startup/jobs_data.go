package startup

import (
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/domain"
	"github.com/Nebojsa1999/XMLProjekat/agent-app-backend/domain/enums"
)

var jobs = []*domain.Job{
	{
		Id:        getObjectId("62bdba74ab73b862988b9fa5"),
		UserId:    getObjectId("123a0cc3a34d25d8567f9f04"),
		CreatedAt: getParsedDateOfBirthFrom("2022-06-26T00:00:00Z"),
		Position:  enums.DevOps,
		Description: "Have you ever dreamed of creating and designing the future? Making a great impact to the whole world? Being the part of privileged team? Well… we happen to have JUST the opportunity for you! \n\n" +
			"At this company we are helping to build the first cognitive city in the world. Based in Middle East - a seamless AI-orchestrated metropolis, a home and a workplace for several million citizens from around the world. \n\n" +
			"We are building a multi-disciplinary team of professionals to skyrocket this exciting initiative. To reach our ambitious goals, we are looking for DevOps Engineer. One more thing… This position can be both office-based and remote and can be based anywhere in Serbia.",
		Requirements: "Skilful in tools required to automate, test, deploy, manage and monitor applications \n\n" +
			"Jenkins, Azure and/or GitLab CI/CD knowledge is a must	\n\n" +
			"Hands-on participation in CI/CD build for cloud-based applications and environments \n\n" +
			"Track record on ways to automate and improve development and release processes \n\n" +
			"Working experience with Cloud environments (AWS, Azure or GCP or Huawei Cloud) \n\n" +
			"Working experience with Kubernetes and Microservice Architecture \n\n" +
			"Knowledge of infrastructure monitoring systems (e.g. The Elastic Stack, Zabbix, Nagios) \n\n" +
			"Strong background in Linux administration, Jira and Confluence \n\n" +
			"About three years of relevant experience,",
	},
}
