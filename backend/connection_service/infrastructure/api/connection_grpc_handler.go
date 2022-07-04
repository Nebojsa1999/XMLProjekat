package api

import (
	"context"
	pb "github.com/Nebojsa1999/XMLProjekat/backend/common/proto/connection_service"
	"github.com/Nebojsa1999/XMLProjekat/backend/connection_service/application"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ConnectionHandler struct {
	pb.UnimplementedConnectionServiceServer
	service    *application.ConnectionService
}

func NewConnectionHandler(service *application.ConnectionService) *ConnectionHandler {
	return &ConnectionHandler{
		service:    service,
	}
}

func (handler *ConnectionHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	connection, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}

	connectionPb := mapDomainConnectionToPbConnection(connection)
	response := &pb.GetResponse{
		Connection: connectionPb,
	}

	return response, nil
}

func (handler *ConnectionHandler) GetByUserId(ctx context.Context, request *pb.GetByUserIdRequest) (*pb.GetMultipleResponse, error) {
	userId := request.UserId
	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	connections, err := handler.service.GetByUserId(objectId)
	if err != nil {
		return nil, err
	}

	response := &pb.GetMultipleResponse{
		Connections: []*pb.Connection{},
	}

	for _, connection := range connections {
		current := mapDomainConnectionToPbConnection(connection)
		response.Connections = append(response.Connections, current)
	}

	return response, nil
}

func (handler *ConnectionHandler) GetFollowingByUserId(ctx context.Context, request *pb.GetByUserIdRequest) (*pb.GetMultipleResponse, error) {
	userId := request.UserId
	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	connections, err := handler.service.GetFollowingByUserId(objectId)
	if err != nil {
		return nil, err
	}

	response := &pb.GetMultipleResponse{
		Connections: []*pb.Connection{},
	}

	for _, connection := range connections {
		current := mapDomainConnectionToPbConnection(connection)
		response.Connections = append(response.Connections, current)
	}

	return response, nil
}

func (handler *ConnectionHandler) GetFollowersByUserId(ctx context.Context, request *pb.GetByUserIdRequest) (*pb.GetMultipleResponse, error) {
	userId := request.UserId
	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	connections, err := handler.service.GetFollowersByUserId(objectId)
	if err != nil {
		return nil, err
	}

	response := &pb.GetMultipleResponse{
		Connections: []*pb.Connection{},
	}

	for _, connection := range connections {
		current := mapDomainConnectionToPbConnection(connection)
		response.Connections = append(response.Connections, current)
	}

	return response, nil
}

func (handler *ConnectionHandler) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	connection := mapPbConnectionToDomainConnection(request.Connection)

	newConnection, err := handler.service.Create(connection)
	if err != nil {
		return nil, err
	}

	return &pb.CreateResponse{Connection: mapDomainConnectionToPbConnection(newConnection)}, nil
}

func (handler *ConnectionHandler) Delete(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	err := handler.service.Delete(request.Id)
	if err != nil {
		return nil, err
	}

	handler.service.Delete(request.Id)

	return &pb.DeleteResponse{}, nil
}

func (handler *ConnectionHandler) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	updatedConnection, err := handler.service.Update(objectId)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateResponse{Connection: mapDomainConnectionToPbConnection(updatedConnection)}, nil
}
