package infrastructure

import (
	"database/sql"
	"fmt"

	"example/todo/interfaces/database"

	_ "github.com/go-sql-driver/mysql"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() *SqlHandler {
	//driverName dataSourceName=[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	conn, err := sql.Open("mysql", "user:password@tcp(db-dev:3306)/db_dev")
	if err != nil {
		panic(err.Error())
	}
	SqlHandler := new(SqlHandler)
	SqlHandler.Conn = conn
	return SqlHandler

}

func (handler *SqlHandler) Execute(statement string, args ...interface{}) (database.Result, error) {
	res := SqlResult{}
	result, err := handler.Conn.Exec(statement, args...)
	if err != nil {
		return res, err
	}
	res.Result = result
	return res, nil
}

func (handler *SqlHandler) Query(statement string, args ...interface{}) (database.Row, error) {
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return new(SqlRow), err
	}
	row := new(SqlRow)
	fmt.Println(row)
	row.Rows = rows
	fmt.Println(row.Rows)
	return row, nil
}

type SqlResult struct {
	Result sql.Result
}

func (r SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r SqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

type SqlRow struct {
	Rows *sql.Rows
}

func (r SqlRow) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
	//↑ここのdestを忘れていたせいでScanした値がきちんと返ってきてなかった
}

func (r SqlRow) Next() bool {
	return r.Rows.Next()
}

func (r SqlRow) Close() error {
	return r.Rows.Close()
}
