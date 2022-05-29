module github.com/Nebojsa1999/XMLProjekat/agent-app-backend

go 1.18

replace github.com/Nebojsa1999/XMLProjekat/agent-app-backend/common => ../common

replace github.com/Nebojsa1999/XMLProjekat/agent-app-backend/domain => ../domain

replace github.com/Nebojsa1999/XMLProjekat/agent-app-backend/common/proto => ../common/proto

replace github.com/Nebojsa1999/XMLProjekat/agent-app-backend/application => ../application

replace github.com/Nebojsa1999/XMLProjekat/agent-app-backend/infrastructure/api => ../infrastructure/api

replace github.com/Nebojsa1999/XMLProjekat/agent-app-backend/infrastructure/persistence => ../infrastructure/persistence

replace github.com/Nebojsa1999/XMLProjekat/agent-app-backend/startup/config => ../startup/config

replace github.com/Nebojsa1999/XMLProjekat/agent-app-backend/startup => ../startup

require go.mongodb.org/mongo-driver v1.9.1
