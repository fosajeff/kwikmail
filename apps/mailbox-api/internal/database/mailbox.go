package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type MailBoxModel struct {
	DB *sql.DB
}

type Mailbox struct {
	Id             string `json:"id"`
	UserId         string `json:"user_id"`
	Address        string `json:"address"`
	CreatedAt      string `json:"created_at"`
	LastActivityAt string `json:"last_activity_at"`
}

func (m *MailBoxModel) Insert(mailbox *Mailbox) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	now := time.Now().UTC().Format(time.RFC3339)
	mailbox.CreatedAt = now
	mailbox.LastActivityAt = now
	mailbox.Id = uuid.New().String()

	query := "INSERT INTO mailboxes (id, user_id, address, created_at, last_activity_at) VALUES ($1, $2, $3, $4, $5)"

	_, err := m.DB.ExecContext(ctx, query, mailbox.Id, mailbox.UserId, mailbox.Address, mailbox.CreatedAt, mailbox.LastActivityAt)
	return err
}

func (m *MailBoxModel) GetAll(userId string) ([]*Mailbox, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "SELECT id, user_id, address, created_at, last_activity_at FROM mailboxes WHERE user_id = $1"
	rows, err := m.DB.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	mailboxes := []*Mailbox{}

	for rows.Next() {
		var mailbox Mailbox
		err := rows.Scan(&mailbox.Id, &mailbox.UserId, &mailbox.Address, &mailbox.CreatedAt, &mailbox.LastActivityAt)
		if err != nil {
			return nil, err
		}
		mailboxes = append(mailboxes, &mailbox)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return mailboxes, nil
}

func (m *MailBoxModel) UpdateLastActivity(mailboxAddress string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	now := time.Now().Format(time.RFC3339)
	query := "UPDATE mailboxes SET last_activity_at = $1 WHERE address = $2"

	_, err := m.DB.ExecContext(ctx, query, now, mailboxAddress)
	return err
}
