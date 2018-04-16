package dataaccess

import (
	"database/sql"
	"fmt"
	"gossip/domain"

	_ "github.com/lib/pq"
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
	fmt.Println("**** Query")
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

type scannable interface {
	Scan(...interface{}) error
}

func scanMessage(row scannable) (msg domain.Message, err error) {
	err = row.Scan(
		&msg.Id,
		&msg.CreatedAt,
		&msg.EditedAt,
		&msg.UserName,
		&msg.Message)
	return
}

func (conn Connection) GetMessages() (result []domain.Message, err error) {
	var rows *sql.Rows
	fmt.Println("Get messages")
	defer func() {
		fmt.Println("Result", result, err)
	}()
	rows, err = conn.db.Query(`select * from messages order by created_at`)
	for err == nil && rows.Next() {
		var msg domain.Message
		msg, err = scanMessage(rows)
		result = append(result, msg)
	}
	return
}

func (conn Connection) GetMessage(id string) (domain.Message, error) {
	row := conn.db.QueryRow(`select
		id, created_at, edited_at, user_name, message
		from messages where id=$1`, id)
	return scanMessage(row)
}

func (conn Connection) UpdateMessage(message domain.Message) error {
	_, err := conn.db.Exec(
		`update messages
		set edited_at=$2, user_name=$3, message=$4
		where
		id=$1`,
		message.Id,
		message.EditedAt,
		message.UserName,
		message.Message)
	return err
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
