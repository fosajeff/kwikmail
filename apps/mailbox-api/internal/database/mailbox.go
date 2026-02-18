package database

import "database/sql"

type MailBoxModel struct {
	DB *sql.DB
}

type Mailbox struct {
	Id             int    `json:"id"`
	UserId         int    `json:"user_id"`
	Address        string `json:"address"`
	CreatedAt      string `json:"created_at"`
	LastActivityAt string `json:"last_activity_at"`
}
