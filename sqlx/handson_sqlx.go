package main

import (
	//    "database/sql"

	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Actor struct {
	ActorId    int       `db:"actor_id"`
	FirstName  string    `db:"first_name"`
	LastName   string    `db:"last_name"`
	LastUpdate time.Time `db:"last_update"`
}

type FilmActor struct {
	ActorId    int       `db:"actor_id"`
	FilmId     string    `db:"film_id"`
	LastUpdate time.Time `db:"last_update"`
}

type ActorAndFilmActor struct {
	ActorId    int       `db:"actor_id"`
	FirstName  string    `db:"first_name"`
	LastName   string    `db:"last_name"`
	LastUpdate time.Time `db:"last_update"`
	FilmActor  `db:"film_actor"`
}

func main() {
	// this Pings the database trying to connect
	// use sqlx.Open() for sql.Open() semantics
	db, err := sqlx.Connect("mysql", "root:root@tcp(127.0.0.1:3306)/sakila?parseTime=true")

	// select
	var actor Actor
	row := db.QueryRowx("SELECT * FROM actor WHERE actor_id=?", 2)
	err = row.StructScan(&actor)
	fmt.Println(actor)

	// 複数行　select
	actors := []Actor{}
	rows, err := db.Queryx("SELECT * FROM actor ORDER BY actor_id LIMIT 10")
	for rows.Next() {
		err = rows.StructScan(&actor)
		fmt.Println(actor)
		actors = append(actors, actor)
	}
	fmt.Println(actors)

	// join 実行例
	aafa := []ActorAndFilmActor{}
	sql := `SELECT
        a.*,
        f.actor_id "film_actor.actor_id",
        f.film_id "film_actor.film_id",
        f.last_update "film_actor.last_update"
      FROM
        actor AS a LEFT JOIN film_actor AS f ON a.actor_id = f.actor_id
        WHERE a.actor_id=1;`
	// Select は Queryx を実行して Scan まで行ってくれる wrapper
	// 全ての取得結果を一度にメモリにロードすることになるの注意
	err = db.Select(&aafa, sql)
	fmt.Println(aafa)
	// range : 配列の中身を変数に展開しつつループを回す
	for _, a := range aafa {
		fmt.Println(a)
	}

	if err != nil {
		log.Fatalln(err)
	}
}
