package main

// import package yang dibutuhkan
import(
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// siapkan struct penampung response
type student struct {
	id string
	name string
	age int
	grade int
}


// fungsi untuk konek ke database mysql
func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_golang_basic")

	if err != nil {
		return nil, err
	}

	return db, nil
}

// contoh fungsi untuk membaca data dari database
func sqlQuery() {
	db, err := connect()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM students WHERE age = ?", 27)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer rows.Close()

	var result []student

	for rows.Next() {
		var each = student{}
		var err = rows.Scan(&each.id, &each.name, &each.grade, &each.grade)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Printf("name: %s\n", each.name)
	}
}

// contoh fungsi untuk membaca data dari database yang menghasilkan 1 record saja
func sqlQueryRow() {
	db, err := connect()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	var result = student{}
	var id = "B001"

	err = db.QueryRow("SELECT name, grade FROM students WHERE id = ?", id).Scan(&result.name, &result.grade)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("name: %s\ngrade: %d\n", result.name, result.grade)
}

// contoh penulisan query dengan menggunakan prepare statement
func sqlPrepare() {
	db, err := connect()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	stmt, err := db.Prepare("SELECT name, grade FROM students WHERE id = ?")

    if err != nil {
        fmt.Println(err.Error())
        return
    }

    var result1 = student{}
    stmt.QueryRow("E001").Scan(&result1.name, &result1.grade)
    fmt.Printf("name: %s\ngrade: %d\n", result1.name, result1.grade)

    var result2 = student{}
    stmt.QueryRow("W001").Scan(&result2.name, &result2.grade)
    fmt.Printf("name: %s\ngrade: %d\n", result2.name, result2.grade)
}

// insert, update dan delete dengan menggunakan perintah Exec
func sqlExec() {
    db, err := connect()

    if err != nil {
        fmt.Println(err.Error())
        return
    }

    defer db.Close()

    _, err = db.Exec("INSERT INTO students VALUES (?, ?, ?, ?)", "G001", "Galahad", 29, 2)

    if err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Println("insert success!")

    _, err = db.Exec("UPDATE students SET age = ? WHERE id = ?", 28, "W001")

    if err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Println("update success!")

    _, err = db.Exec("DELETE FROM students WHERE id = ?", "E001")

    if err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Println("delete success!")
}

func main() {
	// sqlQuery()
	// sqlQueryRow()
	// sqlPrepare()
	sqlExec()
}