package api

import (
	"context"

	pb "github.com/Nebojsa1999/XMLProjekat/backend/common/proto/connection_service"
	"github.com/Nebojsa1999/XMLProjekat/backend/connection_service/application"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ConnectionHandler struct {
	pb.UnimplementedConnectionServiceServer
	service *application.ConnectionService
}

func NewConnectionHandler(service *application.ConnectionService) *ConnectionHandler {
	return &ConnectionHandler{
		service: service,
	}
}

func (handler *ConnectionHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	userId := request.UserId
	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}
	connections, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	response := &pb.GetResponse{
		Connections: []*pb.Connection{},
	}
	for _, Connection := range connections {
		current := mapConnectionToPb(Connection)
		response.Connections = append(response.Connections, current)
	}
	return response, nil
}

func (handler *ConnectionHandler) CreateConnection(ctx context.Context, request *pb.CreateConnectionRequest) (*pb.CreateConnectionResponse, error) {
	newConnection, err := handler.service.CreateConnection(mapPbToConnection(request.Connection))
	if err != nil {
		return nil, err
	}
	response := &pb.CreateConnectionResponse{
		Connection: mapConnectionToPb(newConnection),
	}
	return response, err
}

func (handler *ConnectionHandler) UpdateConnection(ctx context.Context, request *pb.UpdateConnectionRequest) (*pb.UpdateConnectionResponse, error) {
	connection, err := handler.service.UpdateConnection(request.Id)
	if err != nil {
		return nil, err
	}

	_, err = handler.CreateConnection(context.TODO(), &pb.CreateConnectionRequest{
		Connection: mapConnectionToPb(connection),
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateConnectionResponse{
		Connection: mapConnectionToPb(connection),
	}, nil
}

func (handler *ConnectionHandler) DeleteConnection(ctx context.Context, request *pb.DeleteConnectionRequest) (*pb.DeleteConnectionResponse, error) {
	id, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}
	err = handler.service.DeleteConnection(id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteConnectionResponse{}, nil
}
