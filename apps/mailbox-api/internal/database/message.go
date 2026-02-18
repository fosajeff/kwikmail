package database

import "database/sql"

type MessageModel struct {
	DB *sql.DB
}

type Message struct {
	Id         int    `json:"id"`
	MailBoxId  int    `json:"mailbox_id"`
	FromEmail  string `json:"from_email"`
	Subject    string `json:"subject"`
	Body       string `json:"body"`
	RawPayload string `jsonb:"raw_payload"`
	CreatedAt  string `json:"created_at"`
}
