package database

import "database/sql"

type Models struct {
	Mailbox MailBoxModel
	Message MessageModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Mailbox: MailBoxModel{DB: db},
		Message: MessageModel{DB: db},
	}
}
