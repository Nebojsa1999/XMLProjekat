package api

import (
	pb "github.com/Nebojsa1999/XMLProjekat/backend/common/proto/connection_service"
	"github.com/Nebojsa1999/XMLProjekat/backend/connection_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapConnectionToPb(connection *domain.Connection) *pb.Connection {
	return &pb.Connection{
		Id:      connection.Id.Hex(),
		UserAId: connection.UserAId.Hex(),
		UserBId: connection.UserBId.Hex(),
	}
}

func mapPbToConnection(pb *pb.Connection) *domain.Connection {
	return &domain.Connection{
		Id:      getObjectId(pb.Id),
		UserAId: getObjectId(pb.UserAId),
		UserBId: getObjectId(pb.UserBId),
	}
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
