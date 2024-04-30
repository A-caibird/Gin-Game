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
	smtp, err := db.Prepare("SELECT active FROM users WHERE name = ?")
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	row, err := smtp.Query("root")
	defer row.Close()

	var active int
	for row.Next() {
		err := row.Scan(&active)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			return
		}
	}
	fmt.Println(active)
}
