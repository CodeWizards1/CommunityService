package services

import (
	"context"
	"fmt"
	"time"

	com "github.com/Projects/ComunityService/genproto/CommunityService"
	user "github.com/Projects/ComunityService/genproto/UserManagementService"
	"github.com/Projects/ComunityService/storage/postgres"
	"github.com/jmoiron/sqlx"
)

const timeLayout = time.RFC3339

type communityService struct {
	CommunityRepository *postgres.CommunityRepository
	userClient          user.UserManagementServiceClient
	com.UnimplementedCommunityServiceServer
}

func NewCommunityService(db *sqlx.DB, userClient user.UserManagementServiceClient) *communityService {
	return &communityService{
		CommunityRepository: postgres.NewCommunityRepository(db),
		userClient:          userClient,
	}
}

func ProtoToRepoCommunity(protoCommunity *com.Community) *postgres.Community {
	return &postgres.Community{
		ID:          protoCommunity.Id,
		Name:        protoCommunity.Name,
		Description: protoCommunity.Description,
		Location:    protoCommunity.Location,
		CreatedAt:   parseTime(protoCommunity.CreatedAt),
		UpdatedAt:   parseTime(protoCommunity.UpdatedAt),
	}
}

func RepoToProtoCommunity(repoCommunity *postgres.Community) *com.Community {
	return &com.Community{
		Id:          repoCommunity.ID,
		Name:        repoCommunity.Name,
		Description: repoCommunity.Description,
		Location:    repoCommunity.Location,
		CreatedAt:   repoCommunity.CreatedAt.Format(timeLayout),
		UpdatedAt:   repoCommunity.UpdatedAt.Format(timeLayout),
	}
}

func parseTime(timeStr string) time.Time {
	t, _ := time.Parse(timeLayout, timeStr)
	return t
}

func (cs *communityService) JoinCommunity(ctx context.Context, comReq *com.JoinCommunityRequest) (*com.JoinCommunityResponse, error) {
	jComRes := com.JoinCommunityResponse{}
	userID := comReq.UserId

	if userID == "" {
		errMsg := "Error: user ID is empty"
		jComRes.Message = errMsg
		return &jComRes, fmt.Errorf(errMsg)
	}

	userIDReq := user.IdUserRequest{UserId: userID}
	userRes, err := cs.userClient.GetUserById(ctx, &userIDReq)
	if err != nil {
		errMsg := "Error: failed to get user details"
		jComRes.Message = errMsg
		return &jComRes, fmt.Errorf("%s: %v", errMsg, err)
	}

	jComRep := postgres.JoinCommunity{
		CommunityID: comReq.CommunityId,
		UserID:      userRes.UserId,
		JoinedAt:    time.Now().Format(timeLayout),
	}

	joinRes, msg := cs.CommunityRepository.JoinCommunity(ctx, &jComRep)
	if msg.Error != nil {
		errMsg := "Error: failed to join community"
		jComRes.Message = errMsg
		return &jComRes, fmt.Errorf("%s: %v", errMsg, *msg.Error)
	}

	jComRes.Message = fmt.Sprintf("%s successfully joined the community %s", userRes.Username, joinRes.CommunityID)
	return &jComRes, nil
}

func (cs *communityService) LeaveCommunity(ctx context.Context, c *com.LeaveCommunityRequest) (*com.LeaveCommunityResponse, error) {
	userIDReq := user.IdUserRequest{UserId: c.UserId}
	userRes, err := cs.userClient.GetUserById(ctx, &userIDReq)
	if err != nil {
		errMsg := "Error getting user failed"
		return &com.LeaveCommunityResponse{Message: errMsg}, fmt.Errorf("%s: %v", errMsg, err)
	}

	jComRep := postgres.LeaveCommunity{
		CommunityId: c.CommunityId,
		UserID:      userRes.UserId,
	}

	msg := cs.CommunityRepository.LeaveCommunity(ctx, &jComRep)
	if msg.Error != nil {
		errMsg := "Error: failed to leave community for user"
		return &com.LeaveCommunityResponse{Message: errMsg}, fmt.Errorf("%s: %v", errMsg, *msg.Error)
	}

	return &com.LeaveCommunityResponse{Message: fmt.Sprintf("%s successfully left the community %s", userRes.Username, c.CommunityId)}, nil
}

func (cs *communityService) CreateCommunity(ctx context.Context, comReq *com.CreateCommunityRequest) (*com.CreateCommunityResponse, error) {
	community := ProtoToRepoCommunity(comReq.Community)
	communityRes, msg := cs.CommunityRepository.CreateCommunity(ctx, community)
	if msg.Error != nil {
		return nil, fmt.Errorf("error creating community: %v", msg.Error)
	}
	return &com.CreateCommunityResponse{Community: RepoToProtoCommunity(communityRes)}, nil
}

func (cs *communityService) GetCommunityBy(ctx context.Context, comReq *com.GetCommunityRequest) (*com.GetCommunityResponse, error) {
	communityRes, msg := cs.CommunityRepository.GetCommunity(ctx, comReq.Id)
	if msg.Error != nil {
		return nil, fmt.Errorf("error getting community: %v", msg.Error)
	}
	return &com.GetCommunityResponse{Community: RepoToProtoCommunity(communityRes)}, nil
}

func (cs *communityService) GetAllCommunities(ctx context.Context, comReq *com.GetAllCommunityRequest) (*com.GetAllCommunityResponse, error) {
	filter := postgres.CommunityGetFilter{
		Name:   &comReq.Name,
		Limit:  &comReq.Limit,
		Offset: &comReq.Offset,
	}

	communityRes, msg := cs.CommunityRepository.GetAllCommunities(ctx, &filter)
	if msg.Error != nil {
		return nil, fmt.Errorf("error getting communities: %v", msg.Error)
	}

	var communities []*com.Community
	for _, community := range communityRes {
		communities = append(communities, RepoToProtoCommunity(community))
	}

	return &com.GetAllCommunityResponse{Communities: communities}, nil
}

func (cs *communityService) UpdateCommunity(ctx context.Context, upCom *com.UpdateCommunityRequest) (*com.UpdateCommunityResponse, error) {
	community := ProtoToRepoCommunity(upCom.Community)
	upFilter := postgres.CommunityUpdateFilter{
		ID:          &upCom.Community.Id,
		Name:        &community.Name,
		Description: &community.Description,
		Location:    &community.Location,
	}

	communityRes, msg := cs.CommunityRepository.UpdateCommunity(ctx, &upFilter)
	if msg.Error != nil {
		return nil, fmt.Errorf("error updating community: %v", msg.Error)
	}
	return &com.UpdateCommunityResponse{Community: RepoToProtoCommunity(communityRes)}, nil
}

func (cs *communityService) DeleteCommunity(ctx context.Context, comReq *com.DeleteCommunityRequest) (*com.DeleteCommunityResponse, error) {
	msg := cs.CommunityRepository.DeleteCommunity(ctx, comReq.Id)
	if msg.Error != nil {
		return nil, fmt.Errorf("error deleting community: %v", msg.Error)
	}
	return &com.DeleteCommunityResponse{Message: *msg.Message}, nil
}
