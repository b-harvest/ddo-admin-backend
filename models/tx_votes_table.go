package models

import (
	"bharvest-vo/types"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type TxVotesT struct {
	Address    string  `json:"address"`
	Chain      string  `json:"chain"`
	FeeAmount  float64 `json:"feeAmount"`
	FeeDenom   string  `json:"feeDenom"`
	Height     int64   `json:"height"`
	Note       string  `json:"note"`
	Option     string  `json:"option"`
	ProposalId int64   `json:"proposalId"`
	Timestamp  string  `json:"timestamp"`
	TxHash     string  `json:"txHash"`
	WalletCode string  `json:"walletCode"`
}

func GetVotes() ([]TxVotesT, error) {
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
	address, chain, fee_amount, fee_denom, height, note, t.option, proposal_id, 
	timestamp, txhash, wallet_code from tx_votes t 
	order by timestamp DESC`)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil, err
	}

	items := []TxVotesT{}
	for results.Next() {
		var item TxVotesT
		err = results.Scan(
			&item.Address,
			&item.Chain,
			&item.FeeAmount,
			&item.FeeDenom,
			&item.Height,
			&item.Note,
			&item.Option,
			&item.ProposalId,
			&item.Timestamp,
			&item.TxHash,
			&item.WalletCode,
		)
		if err != nil {
			panic(err.Error())
		}
		items = append(items, item)
	}

	return items, nil
}
