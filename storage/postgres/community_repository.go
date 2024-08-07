package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	com "github.com/Projects/ComunityService/genproto/CommunityService"

	"github.com/jmoiron/sqlx"
)

type CommunityRepository struct {
	db *sqlx.DB
}

type Community struct {
	ID          string    `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Location    string    `json:"location,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type Event struct {
	ID          string    `json:"id,omitempty"`
	CommunityID string    `json:"community_id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	EventType   string    `json:"event_type,omitempty"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Location    string    `json:"location,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type CommunityUpdateFilter struct {
	ID          *string `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Location    *string `json:"location,omitempty"`
}

type CommunityGetFilter struct {
	Name     *string `json:"name,omitempty"`
	Location *string `json:"location,omitempty"`
	Limit    *int32  `json:"limit,omitempty"`
	Offset   *int32  `json:"offset,omitempty"`
}

type Message struct {
	Error   *string `json:"error,omitempty"`
	Message *string `json:"message,omitempty"`
}

func NewCommunityRepository(db *sqlx.DB) *CommunityRepository {
	return &CommunityRepository{db: db}
}

func (c *CommunityRepository) CreateCommunity(ctx context.Context, community *Community) (*Community, *Message) {
	community.CreatedAt = time.Now()
	community.UpdatedAt = time.Now()

	query :=
		`
		INSERT INTO communities (name, description, location, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id, name, description, location, created_at, updated_at
    `

	err := c.db.QueryRowContext(ctx, query, community.Name, community.Description, community.Location, community.CreatedAt, community.UpdatedAt).Scan(
		&community.ID,
		&community.Name,
		&community.Description,
		&community.Location,
		&community.CreatedAt,
		&community.UpdatedAt,
	)

	if err != nil {
		errMsg := fmt.Sprintf("Failed to create community: %v", err)
		return nil, &Message{Error: &errMsg}
	}

	successMsg := "Community created successfully"
	return community, &Message{Message: &successMsg}
}

func (c *CommunityRepository) GetCommunity(ctx context.Context, comId string) (*Community, *Message) {
	query :=
		`
		SELECT id, name, description, location, created_at, updated_at
		FROM communities 
		WHERE deleted_at IS NULL AND id = $1
	`

	community := &Community{}

	err := c.db.QueryRowContext(ctx, query, comId).Scan(&community.ID, &community.Name, &community.Description, &community.Location, &community.CreatedAt, &community.UpdatedAt)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to get community: %v", err)
		return nil, &Message{Error: &errMsg}
	}

	successMsg := "Community retrieved successfully"
	return community, &Message{Message: &successMsg}
}

func (c *CommunityRepository) UpdateCommunity(ctx context.Context, com *CommunityUpdateFilter) (*Community, *Message) {
	params := []string{}
	args := []interface{}{}
	argIdx := 1

	if com.Name != nil {
		params = append(params, fmt.Sprintf("name = $%d", argIdx))
		args = append(args, *com.Name)
		argIdx++
	}
	if com.Description != nil {
		params = append(params, fmt.Sprintf("description = $%d", argIdx))
		args = append(args, *com.Description)
		argIdx++
	}
	if com.Location != nil {
		params = append(params, fmt.Sprintf("location = $%d", argIdx))
		args = append(args, *com.Location)
		argIdx++
	}

	if len(params) == 0 {
		errMsg := "Failed to update community: no parameters to update"
		return nil, &Message{Error: &errMsg}
	}

	args = append(args, *com.ID)
	query := fmt.Sprintf("UPDATE communities SET %s, updated_at = NOW() WHERE id = $%d AND deleted_at IS NULL RETURNING id, name, description, location, created_at, updated_at", strings.Join(params, ", "), argIdx)
	fmt.Println(query, args)

	community := &Community{}
	err := c.db.QueryRowContext(ctx, query, args...).Scan(&community.ID, &community.Name, &community.Description, &community.Location, &community.CreatedAt, &community.UpdatedAt)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to update community: %v", err)
		return nil, &Message{Error: &errMsg}
	}

	successMsg := "Community updated successfully"
	return community, &Message{Message: &successMsg}
}

func (c *CommunityRepository) DeleteCommunity(ctx context.Context, comId string) *Message {
	query :=
		`
        UPDATE communities
        SET deleted_at = NOW()
        WHERE id = $1
    `
	_, err := c.db.ExecContext(ctx, query, comId)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to delete community: %v", err)
		return &Message{Error: &errMsg}
	}

	successMsg := "Community deleted successfully"
	return &Message{Message: &successMsg}
}

func (c *CommunityRepository) GetAllCommunities(ctx context.Context, comFilter *CommunityGetFilter) ([]*Community, *Message) {
	params := []string{"deleted_at IS NULL"}
	args := []interface{}{}
	argIdx := 1

	if comFilter.Name != nil {
		params = append(params, fmt.Sprintf("name = $%d", argIdx))
		args = append(args, *comFilter.Name)
		argIdx++
	}

	if comFilter.Location != nil {
		params = append(params, fmt.Sprintf("location = $%d", argIdx))
		args = append(args, *comFilter.Location)
		argIdx++
	}
	query := ""
	if len(params) > 1 {

		query = fmt.Sprintf("SELECT id, name, description, location, created_at, updated_at FROM communities WHERE %s", strings.Join(params, " AND "))
	} else {
		query = "SELECT id, name, description, location, created_at, updated_at FROM communities"

	}

	if comFilter.Limit != nil {
		query += fmt.Sprintf(" LIMIT $%d", argIdx)
		args = append(args, *comFilter.Limit)
		argIdx++
	}

	if comFilter.Offset != nil {
		query += fmt.Sprintf(" OFFSET $%d", argIdx)
		args = append(args, *comFilter.Offset)
		argIdx++
	}

	log.Println("query: ", query)
	log.Println("args: ", args)

	rows, err := c.db.QueryxContext(ctx, query, args...)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to get communities: %v", err)
		return nil, &Message{Error: &errMsg}
	}
	defer rows.Close()

	communities := []*Community{}
	for rows.Next() {
		var id, name, description, location string
		var createdAt, updatedAt time.Time

		if err := rows.Scan(&id, &name, &description, &location, &createdAt, &updatedAt); err != nil {
			errMsg := fmt.Sprintf("Failed to scan community: %v", err)
			return nil, &Message{Error: &errMsg}
		}

		community := &Community{
			ID:          id,
			Name:        name,
			Description: description,
			Location:    location,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
		}
		communities = append(communities, community)
	}
	fmt.Println(communities)
	successMsg := "Communities retrieved successfully"
	return communities, &Message{Message: &successMsg}
}

func (c *CommunityRepository) IsValidCommunity(ctx context.Context, req *com.IsCommunityValidRequest) (*com.IsCommunityValidResponse, error) {
	query :=
		`
			SLECT id FROM community WHERE deleted_at IS NOT NULL AND id = $1 
			RETURNING id
		`

	row := c.db.QueryRowContext(ctx, query, req.Id)
	var id string
	if err := row.Scan(id); err != nil {
		if err == sql.ErrNoRows {
			return &com.IsCommunityValidResponse{Valid: false}, fmt.Errorf("no community found")
		}
	}

	return &com.IsCommunityValidResponse{Valid: true}, nil
}
