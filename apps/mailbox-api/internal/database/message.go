package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type MessageModel struct {
	DB *sql.DB
}

type Message struct {
	Id         string `json:"id"`
	MailBoxId  string `json:"mailbox_id"`
	FromEmail  string `json:"from_email"`
	Subject    string `json:"subject"`
	Body       string `json:"body"`
	RawPayload string `jsonb:"raw_payload"`
	CreatedAt  string `json:"created_at"`
}

func (m *MessageModel) Insert(message *Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	now := time.Now().UTC().Format(time.RFC3339)
	message.CreatedAt = now
	message.Id = uuid.New().String()

	query := "INSERT INTO messages (id, mailbox_id, from_email, subject, body, raw_payload, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7)"

	_, err := m.DB.ExecContext(ctx, query, message.Id, message.MailBoxId, message.FromEmail, message.Subject, message.Body, message.RawPayload, message.CreatedAt)
	return err
}

func (m *MessageModel) GetByMailboxID(mailboxID string) ([]*Message, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "SELECT * FROM messages WHERE mailbox_id = $1"
	rows, err := m.DB.QueryContext(ctx, query, mailboxID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	messages := []*Message{}

	for rows.Next() {
		var message Message
		err := rows.Scan(&message.Id, &message.MailBoxId, &message.FromEmail, &message.Subject, &message.Body, &message.RawPayload, &message.CreatedAt)
		if err != nil {
			return nil, err
		}
		messages = append(messages, &message)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}
