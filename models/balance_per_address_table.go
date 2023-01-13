package models

import (
	"bharvest-vo/types"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type BalancePerAddressT struct {
	Chain               string          `json:"chain"`
	Code                string          `json:"code"`
	Address             string          `json:"address"`
	Height              int64           `json:"height"`
	Symbol              string          `json:"symbol"`
	Denom               string          `json:"denom"`
	AmountUnit          string          `json:"amountUnit"`
	AmountRaw           sql.NullFloat64 `json:"-"`
	Amount              float64         `json:"amount"`
	Variation           float64         `json:"variation"`
	InUsd               float64         `json:"inUsd"`
	InKrw               float64         `json:"inKrw"`
	DelegationBalance   float64         `json:"delegationBalance"`
	UnDelegationBalance float64         `json:"unDelegationBalance"`
	Time                string          `json:"time"`
}

func GetBalancePerAddress() ([]BalancePerAddressT, error) {
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
	results, err := db.Query(`select t.chain, t.code, t.address, t.height, t.symbol, t.denom, t.amount_unit, 
	t.amount, t.variation, t.in_usd, t.in_krw, t.time 
	from balance_per_address t
	inner join (
		select code, denom, max(time) as MaxDate
		from balance_per_address
		where code is not null and code <> '' and code <> 'TERA007'
		group by code, denom
	) tm on t.code = tm.code and t.denom = tm.denom and t.time = tm.MaxDate
	order by t.code;`)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil, err
	}

	balances := []BalancePerAddressT{}
	for results.Next() {
		var balance BalancePerAddressT
		err = results.Scan(
			&balance.Chain,
			&balance.Code,
			&balance.Address,
			&balance.Height,
			&balance.Symbol,
			&balance.Denom,
			&balance.AmountUnit,
			&balance.Amount,
			&balance.Variation,
			&balance.InUsd,
			&balance.InKrw,
			&balance.Time,
		)
		if err != nil {
			panic(err.Error())
		}
		// if account.RawUsd.Valid {
		// 	account.Usd = account.RawUsd.Float64
		// }
		balances = append(balances, balance)
	}

	return balances, nil
}
