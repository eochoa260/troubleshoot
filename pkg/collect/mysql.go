package collect

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	troubleshootv1beta2 "github.com/replicatedhq/troubleshoot/pkg/apis/troubleshoot/v1beta2"
)

func Mysql(c *Collector, databaseCollector *troubleshootv1beta2.Database) (CollectorResult, error) {
	databaseConnection := DatabaseConnection{}

	db, err := sql.Open("mysql", databaseCollector.URI)
	defer db.Close()
	if err != nil {
		databaseConnection.Error = err.Error()
	} else {
		query := `select version()`
		row := db.QueryRow(query)

		version := ""
		if err := row.Scan(&version); err != nil {
			databaseConnection.Error = err.Error()
		} else {
			databaseConnection.IsConnected = true
			databaseConnection.Version = version
		}

		rows, err := db.Query("SHOW VARIABLES")
		defer rows.Close()

		if err != nil {
			databaseConnection.Error = err.Error()
		}

		if err != nil {
			databaseConnection.Error = err.Error()
		}

		variables := map[string]string{}
		for rows.Next() {
			variable := MySQLVariable{}
			err = rows.Scan(&variable.Key, &variable.Value)
			if err != nil {
				databaseConnection.Error = err.Error()
				break
			}
			variables[variable.Key] = variable.Value
		}
		databaseConnection.Variables = variables
	}

	b, err := json.Marshal(databaseConnection)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal database connection")
	}

	collectorName := databaseCollector.CollectorName
	if collectorName == "" {
		collectorName = "mysql"
	}

	output := NewResult()
	output.SaveResult(c.BundlePath, fmt.Sprintf("mysql/%s.json", collectorName), bytes.NewBuffer(b))

	return output, nil
}
