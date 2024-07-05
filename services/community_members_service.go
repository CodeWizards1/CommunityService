package services

import (
	"context"
	"fmt"
	"time"

	com "github.com/Projects/ComunityService/genproto/CommunityService"
	user "github.com/Projects/ComunityService/genproto/UserManagementService"
	"github.com/Projects/ComunityService/storage/postgres"
)

func (cs *communityService) JoinCommunity(ctx context.Context, comReq *com.JoinCommunityRequest) (*com.JoinCommunityResponse, error) {
	jComRes := com.JoinCommunityResponse{}
	userID := comReq.UserId
	fmt.Println(comReq)
	if userID == "" {
		errMsg := "error: user ID is empty"
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
