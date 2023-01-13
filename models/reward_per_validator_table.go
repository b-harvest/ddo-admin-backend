package models

import (
	"bharvest-vo/types"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type TxRewardPerValidatorResT struct {
	Amount          float64 `json:"amount"`
	AmountUnit      string  `json:"amountUnit"`
	Chain           string  `json:"chain"`
	Denom           string  `json:"denom"`
	Height          int64   `json:"height"`
	InKrw           float64 `json:"inKrw"`
	InUsd           float64 `json:"inUsd"`
	OperatorAddress string  `json:"operatorAddress"`
	Symbol          string  `json:"symbol"`
	Time            string  `json:"time"`
	Type            string  `json:"type"`
	Variation       float64 `json:"variation"`
	VariationValue  float64 `json:"variationValue"`
}

func GetValidators() ([]TxRewardPerValidatorResT, error) {
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

	results, err := db.Query(`select t.amount, t.amount_unit, t.chain, 
	t.denom, t.height, t.in_krw, t.in_usd, t.operator_address, t.symbol, t.time, 
	t.type, t.variation, t.variation_value 
	from reward_per_validator t
	inner join (
		select operator_address, denom, max(time) as MaxDate
		from reward_per_validator
		group by operator_address, denom
	) tm on t.operator_address = tm.operator_address and t.denom = tm.denom and t.time = tm.MaxDate
	order by t.chain;`)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil, err
	}

	items := []TxRewardPerValidatorResT{}
	for results.Next() {
		var item TxRewardPerValidatorResT
		err = results.Scan(
			&item.Amount,
			&item.AmountUnit,
			&item.Chain,
			&item.Denom,
			&item.Height,
			&item.InKrw,
			&item.InUsd,
			&item.OperatorAddress,
			&item.Symbol,
			&item.Time,
			&item.Type,
			&item.Variation,
			&item.VariationValue,
		)
		if err != nil {
			panic(err.Error())
		}
		items = append(items, item)
	}

	return items, nil
}
