package bot
import (
	"database/sql"
	"fmt"
	"time"
		_ "github.com/lib/pq"
	)
	type AdditionallyData struct {
		RegistrationDate time.Time
		DateOfApproval  time.Time
		WhoApproved     int
		StatusAdmin     string
	}
	
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
	
		// Создаем нового пользователя в таблице users
		_, err = db.Exec(`
			INSERT INTO "public.user" (name, surname, middle_name, email, phone_number, region, tg_id, "Additional_information")
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
			name, surname, middleName, email, phoneNumber, region, tgID, 0) // Устанавливаем значение по умолчанию (0)
	
		if err != nil {
			return err
		}
	
		fmt.Println("Пользователь успешно создан")
	
		
	
		return nil
	}
	
	func Additional_information(){
		// additionallyData := AdditionallyData{
		// 	RegistrationDate: time.Now(),
		// 	DateOfApproval:  time.Time{}, // Установите дату по необходимости
		// 	WhoApproved:     0,            // Установите значения по необходимости
		// 	StatusAdmin:     "no",         // Установите значения по необходимости
		// }
	
		// // Определите пользователя по email и получите значение "Additional_information"
		// var additionalInfo int
		// err = db.QueryRow(`SELECT Additional_information FROM "public.user" WHERE tg_id = $1`, tgID).Scan(&additionalInfo)
		// if err != nil {
		// 	return err
		// }
	
		// _, err = db.Exec(`
		// 	INSERT INTO "Additionally" (id_additionally, registration_date, date_of_approval, who_approved, status_admin)
		// 	VALUES ($1, $2, $3, $4, $5)`,
		// 	additionalInfo, additionallyData.RegistrationDate, additionallyData.DateOfApproval, additionallyData.WhoApproved, additionallyData.StatusAdmin)
	
		// if err != nil {
		// 	return err
		// }
	}