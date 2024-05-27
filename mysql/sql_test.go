package mysql

import (
	"fmt"
	"os"
	"testing"
)

func TestSql(t *testing.T) {
	db, err := InitDb()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	defer db.Close()
	smtp, err := db.Prepare("SELECT id FROM users WHERE name = ?")
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	row, err := smtp.Query("root")
	defer row.Close()

	var id int
	for row.Next() {
		err := row.Scan(&id)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			return
		}
	}
	fmt.Println(id)
}
