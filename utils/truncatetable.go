package utils

// import (
// 	"database/sql"
// 	"fmt"
// )

// type TruncateTableExecutor struct {
// 	db *sql.DB
// }

// func InitTruncateTableExecutor(db *sql.DB) TruncateTableExecutor {
// 	return TruncateTableExecutor{db}
// }

// func (executor *TruncateTableExecutor) TruncateTable(tableNames []string) {
// 	var err error
// 	tx, err := executor.db.Begin()

// 	if err != nil {
// 		panic(err)
// 	} else {
// 		for _, name := range tableNames {
// 			query := fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE;", name)
// 			_, err := tx.Exec(query)
// 			if err != nil {
// 				panic(err)
// 			}
// 		}
// 	}
// 	if err == nil {
// 		tx.Commit()
// 	} else {
// 		tx.Rollback()
// 		panic(err)
// 	}
// }
