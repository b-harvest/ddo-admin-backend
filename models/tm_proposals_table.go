package models

import (
	"bharvest-vo/types"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type TmProposalsT struct {
	Chain           string `json:"chain"`
	Id              int64  `json:"id"`
	Status          string `json:"propStatus"`
	Time            string `json:"time"`
	Title           string `json:"title"`
	Voted           int64  `json:"voted"`
	VotingEndTime   string `json:"votingEndTime"`
	VotingStartTime string `json:"votingStartTime"`
}

func GetTmProposals() ([]TmProposalsT, error) {
	config := types.GetConfig()
	var (
		DB_HOST = config.Mysql.DbHost
		DB_USER = config.Mysql.DbUser
		DB_PASS = config.Mysql.DbPass
		DB_NAME = config.Mysql.DbName
	)

	db, err := sql.Open("mysql", DB_USER+":"+DB_PASS+"@tcp("+DB_HOST+")/"+DB_NAME)
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil, err
	}
	defer db.Close()

	//results, err := db.Query("SELECT address, chain, code, company, purpose FROM " + DB_TABLE)
	results, err := db.Query(`select 
	chain, id, status, time, title, voted, voting_end_time, voting_start_time  
	from tm_proposals
	order by time DESC`)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil, err
	}

	items := []TmProposalsT{}
	for results.Next() {
		var item TmProposalsT
		err = results.Scan(
			&item.Chain,
			&item.Id,
			&item.Status,
			&item.Time,
			&item.Title,
			&item.Voted,
			&item.VotingEndTime,
			&item.VotingStartTime,
		)
		if err != nil {
			panic(err.Error())
		}
		items = append(items, item)
	}

	return items, nil
}
