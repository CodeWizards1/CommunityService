package services

// import (
// 	"context"
// 	"fmt"
// 	"time"

// 	com "github.com/Projects/ComunityService/genproto/CommunityService"
// 	user "github.com/Projects/ComunityService/genproto/UserManagementService"
// 	"github.com/Projects/ComunityService/storage/postgres"
// )

// func (cs *communityService) JoinCommunity(ctx context.Context, comReq *com.JoinCommunityRequest) (*com.JoinCommunityResponse, error) {
// 	jComRes := com.JoinCommunityResponse{}
// 	// extracting user id from request body
// 	userID := comReq.UserId
// 	// insert user id into request body
// 	userIDReq := user.IdUserRequest{UserId: userID}
// 	// call user get by id service
// 	if userID == "" {
// 		errMsg := "Error: user id is empty"
// 		jComRes.Message = errMsg
// 		return &jComRes, fmt.Errorf("%s", errMsg)
// 	}

// 	userRes, err := cs.userClient.GetUserById(ctx, &userIDReq)
// 	if err != nil {
// 		errMsg := "Error: failed to get user details"
// 		jComRes.Message = errMsg
// 		return &jComRes, fmt.Errorf("User not found: " + err.Error())
// 	}

// 	jComRep := postgres.JoinCommunity{
// 		CommunityID: comReq.CommunityId,
// 		UserID:      userRes.UserId,
// 		JoinedAt:    time.Now().String(),
// 	}

// 	joinRes, msg := cs.CommunityRepository.JoinCommunity(ctx, &jComRep)
// 	if msg.Error != nil {
// 		errMsg := "Error: failed to join community"
// 		jComRes.Message = errMsg
// 		return &jComRes, fmt.Errorf("" + *msg.Error)
// 	}

// 	jComRes.Message = fmt.Sprintf("%s successfully joined the community %s", userRes.Username, joinRes.CommunityID)

// 	return &jComRes, nil
// }

// func (cs *communityService) LeaveCommunity(ctx context.Context, c *com.LeaveCommunityRequest) (*com.LeaveCommunityResponse, error) {
// 	userId := c.UserId

// 	userIDReq := user.IdUserRequest{
// 		UserId: userId,
// 	}

// 	userRes, err := cs.userClient.GetUserById(ctx, &userIDReq)

// 	if err != nil {
// 		msgErr := "Error getting user failed"
// 		return &com.LeaveCommunityResponse{Message: msgErr}, fmt.Errorf("leavecomerr: getting user failed: " + err.Error())
// 	}

// 	jComRep := postgres.LeaveCommunity{
// 		CommunityId: c.CommunityId,
// 		UserID:      userRes.UserId,
// 	}

// 	msg := cs.CommunityRepository.LeaveCommunity(ctx, &jComRep)

// 	if msg.Error != nil {
// 		errMsg := "Error: failed to leave community for user "
// 		return &com.LeaveCommunityResponse{Message: errMsg}, fmt.Errorf("failed to leave community" + *msg.Error)
// 	}

// 	return &com.LeaveCommunityResponse{Message: fmt.Sprintf("%s successfully left the community %s", userRes.Username, c.CommunityId)}, nil
// }
