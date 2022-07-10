package api

import (
	"context"
	pb "github.com/Nebojsa1999/XMLProjekat/backend/common/proto/connection_service"
	"github.com/Nebojsa1999/XMLProjekat/backend/connection_service/application"
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

func (handler *ConnectionHandler) GetConnectionOfFollowingType(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := getObjectId(request.Id)

	followingConnection, err := handler.service.GetConnectionOfFollowingType(id)
	if err != nil {
		return nil, err
	}

	followingConnectionPb := mapDomainConnectionToPbConnection(followingConnection)
	response := &pb.GetResponse{
		Connection: followingConnectionPb,
	}

	return response, nil
}

func (handler *ConnectionHandler) GetAllConnectionsOfFollowingType(ctx context.Context, request *pb.GetAllRequest) (*pb.GetMultipleResponse, error) {
	followingConnections, err := handler.service.GetAllConnectionsOfFollowingType()
	if err != nil {
		return nil, err
	}

	response := &pb.GetMultipleResponse{
		Connections: []*pb.Connection{},
	}

	for _, fC := range followingConnections {
		current := mapDomainConnectionToPbConnection(fC)
		response.Connections = append(response.Connections, current)
	}

	return response, nil
}

func (handler *ConnectionHandler) GetConnectionOfBlockingType(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := getObjectId(request.Id)

	blockingConnection, err := handler.service.GetConnectionOfBlockingType(id)
	if err != nil {
		return nil, err
	}

	blockingConnectionPb := mapDomainConnectionToPbConnection(blockingConnection)
	response := &pb.GetResponse{
		Connection: blockingConnectionPb,
	}

	return response, nil
}

func (handler *ConnectionHandler) GetAllConnectionsOfBlockingType(ctx context.Context, request *pb.GetAllRequest) (*pb.GetMultipleResponse, error) {
	blockingConnections, err := handler.service.GetAllConnectionsOfBlockingType()
	if err != nil {
		return nil, err
	}

	response := &pb.GetMultipleResponse{
		Connections: []*pb.Connection{},
	}

	for _, bC := range blockingConnections {
		current := mapDomainConnectionToPbConnection(bC)
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

func (handler *ConnectionHandler) GetConnectionsOfFollowingTypeByUserId(ctx context.Context, request *pb.GetByUserIdRequest) (*pb.GetMultipleResponse, error) {
	userId := getObjectId(request.UserId)

	followingConnections, err := handler.service.GetConnectionsOfFollowingTypeByUserId(userId)
	if err != nil {
		return nil, err
	}

	response := &pb.GetMultipleResponse{
		Connections: []*pb.Connection{},
	}

	for _, fC := range followingConnections {
		current := mapDomainConnectionToPbConnection(fC)
		response.Connections = append(response.Connections, current)
	}

	return response, nil
}

func (handler *ConnectionHandler) GetConnectionsOfBlockingTypeByUserId(ctx context.Context, request *pb.GetByUserIdRequest) (*pb.GetMultipleResponse, error) {
	userId := getObjectId(request.UserId)

	blockingConnections, err := handler.service.GetConnectionsOfBlockingTypeByUserId(userId)
	if err != nil {
		return nil, err
	}

	response := &pb.GetMultipleResponse{
		Connections: []*pb.Connection{},
	}

	for _, bC := range blockingConnections {
		current := mapDomainConnectionToPbConnection(bC)
		response.Connections = append(response.Connections, current)
	}

	return response, nil
}

func (handler *ConnectionHandler) GetFollowingByUserId(ctx context.Context, request *pb.GetByUserIdRequest) (*pb.GetMultipleResponse, error) {
	userId := getObjectId(request.UserId)

	connectionsInWhichTheGivenUserIsFollowing, err := handler.service.GetFollowingByUserId(userId)
	if err != nil {
		return nil, err
	}

	response := &pb.GetMultipleResponse{
		Connections: []*pb.Connection{},
	}

	for _, c := range connectionsInWhichTheGivenUserIsFollowing {
		current := mapDomainConnectionToPbConnection(c)
		response.Connections = append(response.Connections, current)
	}

	return response, nil
}

func (handler *ConnectionHandler) GetFollowersByUserId(ctx context.Context, request *pb.GetByUserIdRequest) (*pb.GetMultipleResponse, error) {
	userId := getObjectId(request.UserId)

	connectionsInWhichTheGivenUserIsFollowed, err := handler.service.GetFollowersByUserId(userId)
	if err != nil {
		return nil, err
	}

	response := &pb.GetMultipleResponse{
		Connections: []*pb.Connection{},
	}

	for _, c := range connectionsInWhichTheGivenUserIsFollowed {
		current := mapDomainConnectionToPbConnection(c)
		response.Connections = append(response.Connections, current)
	}

	return response, nil
}

func (handler *ConnectionHandler) GetConnectionsInWhichTheGivenUserIsBlocker(ctx context.Context, request *pb.GetByUserIdRequest) (*pb.GetMultipleResponse, error) {
	userId := getObjectId(request.UserId)

	connectionsInWhichTheGivenUserIsBlocker, err := handler.service.GetConnectionsInWhichTheGivenUserIsBlocker(userId)
	if err != nil {
		return nil, err
	}

	response := &pb.GetMultipleResponse{
		Connections: []*pb.Connection{},
	}

	for _, c := range connectionsInWhichTheGivenUserIsBlocker {
		current := mapDomainConnectionToPbConnection(c)
		response.Connections = append(response.Connections, current)
	}

	return response, nil
}

func (handler *ConnectionHandler) GetConnectionsInWhichTheGivenUserIsBlockedOne(ctx context.Context, request *pb.GetByUserIdRequest) (*pb.GetMultipleResponse, error) {
	userId := getObjectId(request.UserId)

	connectionsInWhichTheGivenUserIsBlockedOne, err := handler.service.GetConnectionsInWhichTheGivenUserIsBlockedOne(userId)
	if err != nil {
		return nil, err
	}

	response := &pb.GetMultipleResponse{
		Connections: []*pb.Connection{},
	}

	for _, c := range connectionsInWhichTheGivenUserIsBlockedOne {
		current := mapDomainConnectionToPbConnection(c)
		response.Connections = append(response.Connections, current)
	}

	return response, nil
}

func (handler *ConnectionHandler) GetFollowingUsersIds(ctx context.Context, request *pb.GetByUserIdRequest) (*pb.GetFollowingUsersIdsResponse, error) {
	userId := getObjectId(request.UserId)

	followingUsersIds, err := handler.service.GetFollowingUsersIds(userId)
	if err != nil {
		return nil, err
	}

	response := &pb.GetFollowingUsersIdsResponse{
		Ids: []string{},
	}

	for _, id := range followingUsersIds {
		current := id.Hex()
		response.Ids = append(response.Ids, current)
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
