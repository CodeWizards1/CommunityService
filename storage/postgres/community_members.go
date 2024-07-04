package postgres

import (
	"context"
	"fmt"
)

type JoinCommunity struct {
	CommunityID string `json:"community_id,omitempty"`
	UserID      string `json:"user_id,omitempty"`
	JoinedAt    string `json:"joined_at,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

type LeaveCommunity struct {
	CommunityId string `json:"community_id,omitempty"`
	UserID      string `json:"user_id,omitempty"`
}

func (cs *CommunityRepository) JoinCommunity(ctx context.Context, jCom *JoinCommunity) (*JoinCommunity, *Message) {
	query :=
		`
			INSERT INTO community_members (community_id, user_id, joined_at)
            VALUES ($1, $2, NOW())
            RETURNING community_id, user_id, joined_at, created_at, updated_at
        `

	err := cs.db.QueryRowContext(ctx, query, jCom.CommunityID, jCom.UserID).Scan(&jCom.CommunityID, &jCom.UserID, &jCom.JoinedAt, &jCom.CreatedAt, &jCom.UpdatedAt)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to join community: %v", err)
		return nil, &Message{Error: &errMsg}
	}

	successMsg := "Community joined successfully"
	return jCom, &Message{Message: &successMsg}
}

func (cs *CommunityRepository) LeaveCommunity(ctx context.Context, lCom *LeaveCommunity) Message {
	query :=
		`
	    UPDATE community_members SET Deleted_at = NOW() WHERE community_id = $1 and user_id = $2
		`

	_, err := cs.db.ExecContext(ctx, query, lCom.CommunityId, lCom.UserID)

	if err != nil {
		errMsg := fmt.Sprintf("Failed to execute the leave community query: %v", err)
		return Message{Error: &errMsg}
	}

	successMsg := "Community left successfully"
	return Message{Message: &successMsg}
}
