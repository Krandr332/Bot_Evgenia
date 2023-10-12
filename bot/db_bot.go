package bot
import (
	"database/sql"
	"fmt"
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
	
	func CreateUserAccaunt(tg_id,name,surname,middle_name,phone_number,region,email,password string) error {
		db, err := sql.Open("postgres", "postgres://postgres:1@localhost/base?sslmode=disable")
		if err != nil {
			return err
		}
		defer db.Close()
	
		// Проверяем соединение с базой данных
		err = db.Ping()
		if err != nil {
			return err
		}
	
		// Создаем нового пользователя в таблице users
		_, err = db.Exec(`INSERT INTO "public.user" (username, email, password) VALUES ($1, $2, $3)`, name, email, password)
		if err != nil {
			return err
		}
	
		fmt.Println("Пользователь успешно создан")
	
		return nil
	}