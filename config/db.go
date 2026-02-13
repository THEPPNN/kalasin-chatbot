package config

import (
 "context"
 "github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func InitDB() {
 conn, _ := pgx.Connect(context.Background(), Get("DB_URL"))
 DB = conn
}