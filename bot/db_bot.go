package bot
import (
	"database/sql"

		_ "github.com/lib/pq"
	)
	func checkForUserInSystem(tg_id string) int {
		sqlRequest := `SELECT COUNT(*) FROM "public.User" WHERE tg_id = $1;`
		db, err := sql.Open("postgres", "postgres://postgres:1@localhost/evg_bot?sslmode=disable")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()
	
		var count int
		err = db.QueryRow(sqlRequest, tg_id).Scan(&count)
		if err != nil {
			panic(err.Error())
		}
	
		return count
	}