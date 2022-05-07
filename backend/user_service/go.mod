module github.com/Nebojsa1999/XMLProjekat/backend/user_service

go 1.18

replace github.com/Nebojsa1999/XMLProjekat/backend/common => ../common

replace github.com/Nebojsa1999/XMLProjekat/backend/user_service/domain => ../user_service/domain

replace github.com/Nebojsa1999/XMLProjekat/backend/common/proto/user_service => ../common/proto/user_service

replace github.com/Nebojsa1999/XMLProjekat/backend/user_service/application => ../user_service/application

replace github.com/Nebojsa1999/XMLProjekat/backend/user_service/infrastructure/api => ../user_service/infrastructure/api

replace github.com/Nebojsa1999/XMLProjekat/backend/user_service/infrastructure/persistence => ../user_service/infrastructure/persistence

replace github.com/Nebojsa1999/XMLProjekat/backend/user_service/startup/config => ../user_service/startup/config

replace github.com/Nebojsa1999/XMLProjekat/backend/user_service/startup => ../user_service/startup

require (
	golang.org/x/net v0.0.0-20220425223048-2871e0cb64e4 // indirect
	golang.org/x/sys v0.0.0-20220429121018-84afa8d3f7b3 // indirect
	google.golang.org/protobuf v1.28.0
)

require (
	github.com/Nebojsa1999/XMLProjekat/backend/common v1.0.0
	go.mongodb.org/mongo-driver v1.9.0
	google.golang.org/grpc v1.46.0
)

require (
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.0 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.0.2 // indirect
	github.com/xdg-go/stringprep v1.0.2 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	golang.org/x/crypto v0.0.0-20201216223049-8b5274cf687f // indirect
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220505152158-f39f71e6c8f3 // indirect
)
