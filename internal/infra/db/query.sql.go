// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: query.sql

package db

import (
	"context"
)

const fIndChatByID = `-- name: FIndChatByID :one
SELECT id, user_id, initial_message_id, status, token_usage, model, model_max_tokens, temperature, top_p, n, stop, max_tokens, presence_penalty, frequency_penalty, created_at, updated_at FROM chats WHERE id = ?
`

func (q *Queries) FIndChatByID(ctx context.Context, id string) (Chat, error) {
	row := q.db.QueryRowContext(ctx, fIndChatByID, id)
	var i Chat
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.InitialMessageID,
		&i.Status,
		&i.TokenUsage,
		&i.Model,
		&i.ModelMaxTokens,
		&i.Temperature,
		&i.TopP,
		&i.N,
		&i.Stop,
		&i.MaxTokens,
		&i.PresencePenalty,
		&i.FrequencyPenalty,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
