package migrations

import (
	"encoding/json/v2"

	"ganigeorgiev.com/pocketbase/pocketbase/core"
	m "ganigeorgiev.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_3289777763")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_xpn97uw2hx` + "`" + ` ON ` + "`" + `ledger_entries` + "`" + ` (\n  ` + "`" + `schedule` + "`" + `,\n  ` + "`" + `period` + "`" + `,\n  ` + "`" + `property` + "`" + `\n) WHERE ` + "`" + `schedule` + "`" + ` != '';",
				"CREATE INDEX ` + "`" + `idx_acvbprphw5` + "`" + ` ON ` + "`" + `ledger_entries` + "`" + ` (\n  ` + "`" + `property` + "`" + `,\n  ` + "`" + `posted_on` + "`" + `\n)"
			]
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_3289777763")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_xpn97uw2hx` + "`" + ` ON ` + "`" + `ledger_entries` + "`" + ` (\n  ` + "`" + `schedule` + "`" + `,\n  ` + "`" + `period` + "`" + `,\n  ` + "`" + `property` + "`" + `\n) WHERE 'schedule' != '';",
				"CREATE INDEX ` + "`" + `idx_acvbprphw5` + "`" + ` ON ` + "`" + `ledger_entries` + "`" + ` (\n  ` + "`" + `property` + "`" + `,\n  ` + "`" + `posted_on` + "`" + `\n)"
			]
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
