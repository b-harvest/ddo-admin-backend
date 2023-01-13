package models

import (
	"bharvest-vo/types"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type TxCryptoFlowT struct {
	Chain       string  `json:"chain"`
	WalletCode  string  `json:"walletCode"`
	Address     string  `json:"address"`
	Height      int64   `json:"height"`
	TxHash      string  `json:"txHash"`
	Action      string  `json:"action"`
	FromAddress string  `json:"fromAddress"`
	ToAddress   string  `json:"toAddress"`
	Amount      float64 `json:"amount"`
	Denom       string  `json:"denom"`
	InOut       int64   `json:"inOut"`
	FeeAmount   float64 `json:"feeAmount"`
	FeeDenom    string  `json:"feeDenom"`
	Note        string  `json:"note"`
	Timestamp   string  `json:"timestamp"`
}

func GetAllTxCryptoFlow() ([]TxCryptoFlowT, error) {
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
	results, err := db.Query(`select chain, wallet_code, address, height, txhash, action, 
	from_address, to_address, amount, denom, in_out, fee_amount, fee_denom, note, timestamp 
	from tx_crypto_flow order by timestamp DESC`)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil, err
	}

	txs := []TxCryptoFlowT{}
	for results.Next() {
		var tx TxCryptoFlowT
		err = results.Scan(
			&tx.Chain,
			&tx.WalletCode,
			&tx.Address,
			&tx.Height,
			&tx.TxHash,
			&tx.Action,
			&tx.FromAddress,
			&tx.ToAddress,
			&tx.Amount,
			&tx.Denom,
			&tx.InOut,
			&tx.FeeAmount,
			&tx.FeeDenom,
			&tx.Note,
			&tx.Timestamp,
		)
		if err != nil {
			panic(err.Error())
		}
		// if account.RawUsd.Valid {
		// 	account.Usd = account.RawUsd.Float64
		// }
		txs = append(txs, tx)
	}

	return txs, nil
}
