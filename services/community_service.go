package services

import (
	"context"
	"fmt"
	"time"

	com "github.com/Projects/ComunityService/genproto/CommunityService"
	"github.com/Projects/ComunityService/storage/postgres"
	"github.com/jmoiron/sqlx"
)

type communityService struct {
	CommunityRepository *postgres.CommunityRepository
	com.UnimplementedCommunityServiceServer
}

func NewCommunityService(db *sqlx.DB) *communityService {
	return &communityService{CommunityRepository: postgres.NewCommunityRepository(db)}
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
		CreatedAt:   repoCommunity.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   repoCommunity.UpdatedAt.Format(time.RFC3339),
	}
}

func parseTime(timeStr string) time.Time {
	t, _ := time.Parse(time.RFC3339, timeStr)
	return t
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

	if communities == nil {
		communities = []*com.Community{}
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
