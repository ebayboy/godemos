package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"time"

	_ "github.com/ClickHouse/clickhouse-go/v2"
)

func ConnectDSN(source string) (*sql.DB, error) {

	conn, err := sql.Open("clickhouse", source)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}

func TestInsertTx(source string) {
	conn, err := ConnectDSN(source)
	if err != nil {
		fmt.Println("error:", err.Error())
		return
	}

	defer func() {
		_ = conn.Close()
	}()

	tx, err := conn.Begin()
	if err != nil {
		fmt.Println("error:", err.Error())
		return
	}

	stmt, err := tx.Prepare("INSERT INTO client_log (cg_id, conn_id, level, msg, service, time, create_time) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal("error:", err.Error())
	}

	if _, err := stmt.Exec("cgid_111", "connid_111", "info", msg, "clickhouse_test.1", "2023-07-03T16:13:21+08:00", time.Now()); err != nil {
		log.Fatal("error:", err.Error())
	}

	if err = tx.Commit(); err != nil {
		log.Fatal("error:", err.Error())
	}

}

var (
	host     string
	password string
	database string
	username string
	msg      string
)

func main() {

	flag.StringVar(&host, "host", "10.226.133.79", "host")
	flag.StringVar(&password, "password", "", "password")
	flag.StringVar(&database, "database", "testhi", "database")
	flag.StringVar(&username, "username", "default", "username")
	flag.StringVar(&msg, "msg", "msg_"+time.Now().String(), "msg")

	flag.Parse()

	source := fmt.Sprintf("http://%s:8123?debug=true&username=%s&password=%s&database=%s", host, username, password, database)
	log.Println("source:", source)

	//TestInsertTx(source)

	TestInsert(source)

	TestSelect(source)
}

func TestSelect(source string) {
	conn, err := ConnectDSN(source)
	if err != nil {
		fmt.Println("error:", err.Error())
		return
	}

	defer func() {
		_ = conn.Close()
	}()

	cmd := fmt.Sprintf("select count(*) from client_log where msg='%s'", msg)
	fmt.Println("select cmd:", cmd)
	rows, err := conn.Query(cmd)
	if err != nil {
		log.Fatal("error:", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var cnt uint64
		if err := rows.Scan(&cnt); err != nil {
			log.Fatal("error:", err.Error())
			panic(err)
		}
		log.Println("count:", cnt)
	}
}

func TestInsert(source string) {
	conn, err := ConnectDSN(source)
	if err != nil {
		fmt.Println("error:", err.Error())
		return
	}

	defer func() {
		_ = conn.Close()
	}()

	cmd := fmt.Sprintf("INSERT INTO client_log (cg_id, conn_id, level, msg, service, time) VALUES ('%s', '%s', '%s', '%s', '%s', '%s')",
		"cgid_111", "connid_111", "info", msg, "clickhouse_test.1", "2023-07-03T16:13:21+08:00")

	res, err := conn.Exec(cmd)
	if err != nil {
		log.Fatal("error:", err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Fatal("error:", err.Error())
	}
	log.Println("affected_rows:", rows)
}
