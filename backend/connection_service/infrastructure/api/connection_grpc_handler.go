package api

import (
	"context"
	pb "github.com/Nebojsa1999/XMLProjekat/backend/common/proto/connection_service"
	"github.com/Nebojsa1999/XMLProjekat/backend/connection_service/application"
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

func (handler *ConnectionHandler) GetByUserId(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	connections, err := handler.service.Get(request.UserId)
	if err != nil {
		return nil, err
	}

	response := &pb.GetResponse{
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

	if newConnection.IsApproved {
		_, err = handler.service.Create(newConnection)
		if err != nil {
			handler.service.Delete(newConnection.Id.Hex())
			return nil, err
		}
	}

	return &pb.CreateResponse{
		Connection: mapDomainConnectionToPbConnection(newConnection),
	}, nil
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
	connection, err := handler.service.Update(request.Id)
	if err != nil {
		return nil, err
	}

	if connection.IsApproved {
		_, err = handler.service.Create(connection)
		if err != nil {
			return nil, err
		}
	}

	return &pb.UpdateResponse{
		Connection: mapDomainConnectionToPbConnection(connection),
	}, nil
}
