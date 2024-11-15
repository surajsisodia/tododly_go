package utils

import (
	"database/sql"
	"fmt"
	"strings"
)

func ConvertRowsToMap(rows *sql.Rows) []map[any]interface{} {

	columns, err := rows.Columns()
	if err != nil {
		return nil
	}

	columnType, err := rows.ColumnTypes()
	if err != nil {
		panic(err)
	}

	fmt.Println(columnType[0].Name())

	var results []map[any]interface{}

	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePointers := make([]interface{}, len(columns))
		for i := range values {
			valuePointers[i] = &values[i]
		}

		err := rows.Scan(valuePointers...)
		if err != nil {
			return nil
		}

		rowMap := make(map[any]interface{})
		for i, colName := range columns {
			// fmt.Printf("Column Type: %s", values[i])
			// fmt.Print("Values: %s", values[i])

			var v interface{}
			b, ok := values[i].([]byte)
			if ok {
				v = b
			} else {
				v = values[i]
			}
			rowMap[strings.ToLower(colName)] = v
		}

		results = append(results, rowMap)
	}

	return results
}
