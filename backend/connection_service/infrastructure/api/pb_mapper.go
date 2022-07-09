package api

import (
	pb "github.com/Nebojsa1999/XMLProjekat/backend/common/proto/connection_service"
	"github.com/Nebojsa1999/XMLProjekat/backend/connection_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapDomainConnectionToPbConnection(connection *domain.Connection) *pb.Connection {
	return &pb.Connection{
		Id:         connection.Id.Hex(),
		IssuerId:   connection.IssuerId.Hex(),
		SubjectId:  connection.SubjectId.Hex(),
		Date:       timestamppb.New(connection.Date),
		IsApproved: connection.IsApproved,
	}
}

func mapPbConnectionToDomainConnection(pbConnection *pb.Connection) *domain.Connection {
	return &domain.Connection{
		Id:         getObjectId(pbConnection.Id),
		IssuerId:   getObjectId(pbConnection.IssuerId),
		SubjectId:  getObjectId(pbConnection.SubjectId),
		Date:       pbConnection.Date.AsTime(),
		IsApproved: pbConnection.IsApproved,
	}
}

func mapDomainConnectionUpdateDTOToPbConnectionUpdateDTO(connectionUpdateDTO *domain.ConnectionUpdateDTO) *pb.ConnectionUpdateDTO {
	return &pb.ConnectionUpdateDTO{
		IssuerId:  connectionUpdateDTO.IssuerId.Hex(),
		SubjectId: connectionUpdateDTO.SubjectId.Hex(),
		IsApproved:  connectionUpdateDTO.IsApproved,
	}
}

func mapPbConnectionUpdateDTOToDomainConnectionUpdateDTO(pbConnectionUpdateDTO *pb.ConnectionUpdateDTO) *domain.ConnectionUpdateDTO {
	return &domain.ConnectionUpdateDTO{
		IssuerId:  getObjectId(pbConnectionUpdateDTO.IssuerId),
		SubjectId: getObjectId(pbConnectionUpdateDTO.SubjectId),
		IsApproved:  pbConnectionUpdateDTO.IsApproved,
	}
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}

	return primitive.NewObjectID()
}
