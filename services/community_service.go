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

func (cs *communityService) CreateCommunity(ctx context.Context, comReq *com.CreateCommunityRequest) (*com.CreateCommunityResponse, error) {
	community := ProtoToRepoCommunity(comReq.Community)
	communityRes, msg := cs.CommunityRepository.CreateCommunity(ctx, community)
	if msg.Error != nil {
		return nil, fmt.Errorf("error creating community: %v", *msg.Error)
	}
	return &com.CreateCommunityResponse{Community: RepoToProtoCommunity(communityRes)}, nil
}

func (cs *communityService) GetCommunityBy(ctx context.Context, comReq *com.GetCommunityRequest) (*com.GetCommunityResponse, error) {
	communityRes, msg := cs.CommunityRepository.GetCommunity(ctx, comReq.Id)
	if msg.Error != nil {
		return nil, fmt.Errorf("error getting community: %v", *msg.Error)
	}

	return &com.GetCommunityResponse{Community: RepoToProtoCommunity(communityRes)}, nil
}

func (cs *communityService) GetAllCommunity(ctx context.Context, comReq *com.GetAllCommunityRequest) (*com.GetAllCommunityResponse, error) {
	filter := postgres.CommunityGetFilter{
		Name:   &comReq.Name,
		Limit:  &comReq.Limit,
		Offset: &comReq.Offset,
	}

	communityRes, msg := cs.CommunityRepository.GetAllCommunities(ctx, &filter)
	if msg.Error != nil {
		return nil, fmt.Errorf("error getting communities: %v", *msg.Error)
	}

	fmt.Println(communityRes)
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
		return nil, fmt.Errorf("error updating community: %v", *msg.Error)
	}
	return &com.UpdateCommunityResponse{Community: RepoToProtoCommunity(communityRes)}, nil
}

func (cs *communityService) DeleteCommunity(ctx context.Context, comReq *com.DeleteCommunityRequest) (*com.DeleteCommunityResponse, error) {
	msg := cs.CommunityRepository.DeleteCommunity(ctx, comReq.Id)
	if msg.Error != nil {
		return nil, fmt.Errorf("error deleting community: %v", *msg.Error)
	}
	return &com.DeleteCommunityResponse{Message: *msg.Message}, nil
}

func (cs *communityService) IsValidCommunity(ctx context.Context, comReq *com.IsCommunityValidRequest) (*com.IsCommunityValidResponse, error) {
	return cs.CommunityRepository.IsValidCommunity(ctx, comReq)
}
