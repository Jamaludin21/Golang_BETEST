package query_member

import (
	"context"
	"fmt"
	"log"
	"restapi/config"
	"restapi/models"
)

const (
	table = "member"
)

// GetAll
func GetAll(ctx context.Context) ([]models.Member, error) {

	var members []models.Member

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By id ASC", table)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var member models.Member

		if err = rowQuery.Scan(&member.USERNAME,
			&member.GENDER,
			&member.SKINTYPE,
			&member.SKINCOLOR,
		); err != nil {
			return nil, err
		}

		members = append(members, member)
	}

	return members, nil
}

// func Insert(c echo.Context, mbr models.Member) error {
// 	db, err := config.MySQL()

// 	if err != nil {
// 		log.Fatal("Can't connect to MySQL", err)
// 	}

// 	queryText := fmt.Sprintf("INSERT INTO %v (username,gender,skintype,skincolor) values(%v,'%v',%v,'%v')", table,
// 		mbr.USERNAME,
// 		mbr.GENDER,
// 		mbr.SKINTYPE,
// 		mbr.SKINCOLOR)

// 	_, err = db.ExecContext(c, queryText)

// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
