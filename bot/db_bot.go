// db_bot.go
package bot
import (
	"database/sql"
	"fmt"
	"time"
	"log"
		_ "github.com/lib/pq"
	)
	type AdditionallyData struct {
		RegistrationDate time.Time
		DateOfApproval  time.Time
		WhoApproved     int
		StatusAdmin     string
	}
	
	func checkForUserInSystem(tg_id string) int {
		sqlRequest := `SELECT COUNT(*) FROM "public.user" WHERE tg_id = $1;`
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
	func checkAdminStatus(tgID int) (bool, error) {
		// Подключение к базе данных
		db, err := sql.Open("postgres", "postgres://postgres:1@localhost/evg_bot?sslmode=disable")
		if err != nil {
			return false, err
		}
		defer db.Close()
	
		
	
		// Выполняем SQL-запрос для получения статуса админа
		var adminLevel int

		err = db.QueryRow(`
			SELECT ad."level"
			FROM "public.user" u
			JOIN "public.Additionally" a ON u."Additional_information" = a."id_additionally"
			JOIN "public.Admin" ad ON a."admin_status" = ad."id_admin"
			WHERE u.tg_id = $1`, tgID).Scan(&adminLevel)

			if err != nil {
				return false, err
			}
		
			fmt.Printf("Admin Level: %d\n", adminLevel)
		
			isAdmin := adminLevel > 1
			return isAdmin, nil
		

	}
	
	
	
	
	func CreateUserAccount(tgID, name, surname, middleName, phoneNumber, region, email string) error {
		db, err := sql.Open("postgres", "postgres://postgres:1@localhost/evg_bot?sslmode=disable")
		if err != nil {
			return err
		}
		defer db.Close()
	
		err = db.Ping()
		if err != nil {
			return err
		}
	
		// Шаг 1: Создание записи в таблице "public.Additionally" и получение ID
		additionallyID, err := CreateAdditionallyRecord()
		if err != nil {
			log.Println(err)
			return err
		}
	
		// Шаг 2: Создание записи в таблице "public.user" и указание значения внешнего ключа
		_, err = db.Exec(`
			INSERT INTO "public.user" (name, surname, middle_name, email, phone_number, region, tg_id, "Additional_information")
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
			name, surname, middleName, email, phoneNumber, region, tgID, additionallyID)
	
		if err != nil {
			return err
		}
	
		fmt.Println("Пользователь успешно создан")
	
		return nil
	}
	
	func CreateAdditionallyRecord() (int, error) {
		db, err := sql.Open("postgres", "postgres://postgres:1@localhost/evg_bot?sslmode=disable")
		if err != nil {
			return 0, err
		}
		defer db.Close()
	
		err = db.Ping()
		if err != nil {
			return 0, err
		}
	
		var additionallyID int
		err = db.QueryRow(`
			INSERT INTO "public.Additionally" ("registration_date", "admin_status")
			VALUES (NOW(), 1)
			RETURNING "id_additionally"`).Scan(&additionallyID)
	
		if err != nil {
			return 0, err
		}
	
		return additionallyID, nil
	}
	func getUsersWithStatusZero() ([]string, error) {
		db, err := sql.Open("postgres", "postgres://postgres:1@localhost/evg_bot?sslmode=disable")
		if err != nil {
			return nil, err
		}
		defer db.Close()
	
		rows, err := db.Query(`
			SELECT u."id_user",u."name", u."surname", u."middle_name"
			FROM "public.user" u
			JOIN "public.Additionally" a ON u."Additional_information" = a."id_additionally"
			WHERE a."status" = 0`)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
	
		users := []string{}
		for rows.Next() {
			var id_user ,name, surname, middleName string
			if err := rows.Scan(&id_user, &name, &surname, &middleName); err != nil {
				return nil, err
			}
			users = append(users, fmt.Sprintf("%s %s %s %s",id_user, name, surname, middleName))
		}
		if err := rows.Err(); err != nil {
			return nil, err
		}
	
		return users, nil
	}
	func updateUserStatusToApproved(userID int) error {
		db, err := sql.Open("postgres", "postgres://postgres:1@localhost/evg_bot?sslmode=disable")
		if err != nil {
			return err
		}
		defer db.Close()
	
		// Выполняем SQL-запрос для обновления статуса
		_, err = db.Exec(`
			UPDATE "public.Additionally" a
			SET "status" = 1
			FROM "public.user" u
			WHERE u."id_user" = $1 AND u."Additional_information" = a."id_additionally"`,
			userID)
	
		if err != nil {
			return err
		}
	
		return nil
	}