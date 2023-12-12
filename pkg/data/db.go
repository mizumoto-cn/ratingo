package data

import (
	"github.com/mizumoto-cn/ratingo/pkg/data/ent"

	_ "github.com/go-sql-driver/mysql"
)

func NewClient() (*ent.Client, error) {
	return ent.Open("mysql", "root:root@tcp(localhost:3306)/rating?parseTime=True")
}
