package main

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

func main() {
	c := sq.Case().When("i = 1", "c.name").Else("c.type")
	s := sq.Select("DISTINCT ON (u.id) u.*", "json_agg(u.id)").
		From("users AS u").
		Where("json_agg(u.id = $1)", 1).
		Where(c).
		OrderBy("u.name ASC")

	sql, data := s.MustSql()
	fmt.Println(sql)
	fmt.Println(data)
}
