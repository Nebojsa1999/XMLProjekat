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

func (handler *ConnectionHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := getObjectId(request.Id)

	connection, err := handler.service.Get(id)
	if err != nil {
		return nil, err
	}

	connectionPb := mapDomainConnectionToPbConnection(connection)
	response := &pb.GetResponse{
		Connection: connectionPb,
	}

	return response, nil
}

func (handler *ConnectionHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetMultipleResponse, error) {
	connections, err := handler.service.GetAll()
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

func (handler *ConnectionHandler) GetByUserId(ctx context.Context, request *pb.GetByUserIdRequest) (*pb.GetMultipleResponse, error) {
	userId := getObjectId(request.UserId)

	connections, err := handler.service.GetByUserId(userId)
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
	userId := getObjectId(request.UserId)

	connections, err := handler.service.GetFollowingByUserId(userId)
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
	userId := getObjectId(request.UserId)

	connections, err := handler.service.GetFollowersByUserId(userId)
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

func (handler *ConnectionHandler) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	connectionUpdateDTO := mapPbConnectionUpdateDTOToDomainConnectionUpdateDTO(request.ConnectionUpdateDTO)

	updatedConnection, err := handler.service.Update(connectionUpdateDTO)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateResponse{Connection: mapDomainConnectionToPbConnection(updatedConnection)}, nil
}

func (handler *ConnectionHandler) Delete(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	issuerId := getObjectId(request.IssuerId)
	subjectId := getObjectId(request.SubjectId)

	err := handler.service.Delete(issuerId, subjectId)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteResponse{}, nil
}
