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
	msg := &postgres.Message{}

	communityRes, msg := cs.CommunityRepository.CreateCommunity(ctx, community)
	if msg.Error != nil {
		return nil, fmt.Errorf("error creating community: %v", msg.Error)
	}

	return &com.CreateCommunityResponse{Community: RepoToProtoCommunity(communityRes)}, nil
}
