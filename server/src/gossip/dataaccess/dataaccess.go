package dataaccess

import (
	"database/sql"
	_ "github.com/lib/pq"
	"gossip/domain"
)

type Connection struct {
	db *sql.DB
}

func NewConnection(url string) (conn Connection, err error) {
	if conn.db, err = sql.Open("postgres", url); err == nil {
		err = conn.db.Ping()
	}
	return
}

func (conn Connection) InsertMessage(message domain.Message) error {
	_, err := conn.db.Exec(
		`insert into messages 
		( id, "created_at", "edited_at", "user_name", "message") values
		( $1, $2, $3, $4, $5)`,
		message.Id,
		message.CreatedAt,
		message.EditedAt,
		message.UserName,
		message.Message)
	return err
}

func (conn Connection) GetMessage(id string) (msg domain.Message, err error) {
	row := conn.db.QueryRow(`select
		id, created_at, edited_at, user_name, message
		from messages where id=$1`, id)
	err = row.Scan(
		&msg.Id,
		&msg.CreatedAt,
		&msg.EditedAt,
		&msg.UserName,
		&msg.Message)
	return
}

func (conn Connection) Migrate() error {
	sql := `CREATE TABLE IF NOT EXISTS messages (
    id text COLLATE pg_catalog."default" NOT NULL,
    created_at timestamp with time zone NOT NULL,
    edited_at timestamp with time zone,
    user_name text COLLATE pg_catalog."default" NOT NULL,
    message text COLLATE pg_catalog."default" NOT NULL
	)`
	_, err := conn.db.Exec(sql)
	return err
}
