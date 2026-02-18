CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE messages (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  mailbox_id UUID NOT NULL,
  from_email TEXT NOT NULL,
  subject TEXT,
  body TEXT,
  raw_payload JSONB,
  created_at TIMESTAMPTZ DEFAULT now(),
  CONSTRAINT fk_mailbox
    FOREIGN KEY (mailbox_id)
    REFERENCES mailboxes(id)
    ON DELETE CASCADE
);