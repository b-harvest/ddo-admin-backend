package models

import (
	"bharvest-vo/types"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type AccountsBharvestT struct {
	Address      string          `json:"address"`
	Chain        string          `json:"chain"`
	Code         string          `json:"code"`
	Company      string          `json:"company"`
	CreatedAt    string          `json:"createdAt"`
	Id           string          `json:"id"`
	IsStaking    int64           `json:"isStaking"`
	Manage       string          `json:"manage"`
	Memo         string          `json:"memo"`
	Purpose      string          `json:"purpose"`
	Revenue      string          `json:"revenue"`
	TargetTokens string          `json:"targetTokens"`
	TargetTrans  string          `json:"targetTrans"`
	UpdatedAt    string          `json:"updatedAt"`
	RawUsd       sql.NullFloat64 `json:"-"`
	Usd          float64         `json:"usd"`
}

func GetAccountsBharvestT() ([]AccountsBharvestT, error) {
	config := types.GetConfig()
	var (
		DB_HOST = config.Mysql.DbHost
		DB_USER = config.Mysql.DbUser
		DB_PASS = config.Mysql.DbPass
		DB_NAME = config.Mysql.DbName
		//DB_TABLE = types.DB_TABLE_ACCOUNTS_BHARVEST
	)

	db, err := sql.Open("mysql", DB_USER+":"+DB_PASS+"@tcp("+DB_HOST+")/"+DB_NAME)
	//db, err := sql.Open("mysql", DB_USER+":"+DB_PASS+"@tcp(127.0.0.1:3306)/"+DB_NAME)
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil, err
	}
	defer db.Close()

	//results, err := db.Query("SELECT address, chain, code, company, purpose FROM " + DB_TABLE)
	results, err := db.Query(`select a.address, a.chain, a.code, a.company, a.purpose, am.usd 
		FROM accounts_bharvest a
		LEFT JOIN (
			select t.code as code, sum(in_usd) as usd
			from balance_per_address t
			inner join (
				select code, denom, max(time) as MaxDate
				from balance_per_address
				where code is not null and code <> '' and code <> 'TERA007'
				group by code, denom
			) tm on t.code = tm.code and t.denom = tm.denom and t.time = tm.MaxDate
			group by t.code
		) am on a.code = am.code
		order by a.code`)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil, err
	}

	accounts := []AccountsBharvestT{}
	for results.Next() {
		var account AccountsBharvestT
		err = results.Scan(
			&account.Address,
			&account.Chain,
			&account.Code,
			&account.Company,
			&account.Purpose,
			&account.RawUsd,
		)
		if err != nil {
			panic(err.Error())
		}
		if account.RawUsd.Valid {
			account.Usd = account.RawUsd.Float64
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}
