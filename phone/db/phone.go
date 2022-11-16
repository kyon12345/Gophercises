package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func init() {
	db, err := sql.Open("sqlite3", "./phone.db")
	defer db.Close()
	checkErr(err)

	//create table
	createTableSql := `CREATE TABLE 'Phone'('phone_id' INTEGER PRIMARY KEY AUTOINCREMENT,'phone_number' VARCHAR(64) NOT NULL)`
	_, err = db.Exec(createTableSql)
	checkErr(err)
}

func CreatePhoneNumberBatch(phones []string) error {
	db, err := sql.Open("sqlite3", "./phone.db")
	defer db.Close()
	checkErr(err)

	stmt, err := db.Prepare("INSERT INTO Phone(phone_number) values (?)")

	if err != nil {
		return err
	}

	for i := 0; i < len(phones); i++ {
		if _, err := stmt.Exec(phones[i]); err != nil {
			return err
		}
	}

	return nil
}

func UpdatePhoneNumber(phone string, id int) error {
	db, err := sql.Open("sqlite3", "./phone.db")
	defer db.Close()
	checkErr(err)

	stmt, err := db.Prepare("UPDATE Phone SET phone_numeber = '?' ")

	if err != nil {
		return err
	}

	if _, err := stmt.Exec(phone); err != nil {
		return err
	}
	return nil
}

func DeletePhoneNumber(id int) error {
	db, err := sql.Open("sqlite3", "./phone.db")
	defer db.Close()
	checkErr(err)

	stmt, err := db.Prepare("DELETE Phone WHERE id = ?")
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(id); err != nil {
		return err
	}
	return nil
}

func PrintAll() {
	db, err := sql.Open("sqlite3", "./phone.db")
	defer db.Close()
	checkErr(err)

	row, err := db.Query("SELECT phone_number From Phone")
	defer row.Close()
	checkErr(err)

	var phone_number string
	for row.Next() {
		err := row.Scan(&phone_number)
		checkErr(err)
		fmt.Println(phone_number)
	}
}

func QueryPhoneNumbers(phone string) ([]int, error) {
	db, err := sql.Open("sqlite3", "./phone.db")
	defer db.Close()
	checkErr(err)

	row, err := db.Query(`SELECT id FROM Phone Where phone_number = ` + phone + ``)
	defer row.Close()
	if err != nil {
		return nil, err
	}
	var id int
	res := make([]int, 0)
	for row.Next() {
		err := row.Scan(&id)
		if err != nil {
			return nil, err
		}
		res = append(res, id)
	}

	return res, nil
}

func QueryPhoneNumberFirst(phone string) (string, int, error) {
	db, err := sql.Open("sqlite3", "./phone.db")
	defer db.Close()
	checkErr(err)

	row, err := db.Query(`SELECT * FROM Phone`)
	if err != nil {
		return "", -1, err
	}
	var id int
	var phone_number string
	for row.Next() {
		err := row.Scan(&id, &phone_number)
		if err != nil {
			return "", -1, err
		}
		if phone_number == phone {
			return phone_number, id, nil
		}
	}
	defer row.Close()
	return "", -1, nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
