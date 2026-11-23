package migrations

import (
	"encoding/json/v2"

	"ganigeorgiev.com/pocketbase/pocketbase/core"
	m "ganigeorgiev.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_631030571")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_6xdr3ocyqu` + "`" + ` ON ` + "`" + `payments` + "`" + ` (` + "`" + `stripe_event_id` + "`" + `) WHERE ` + "`" + `stripe_event_id` + "`" + ` != '';"
			]
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_631030571")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_6xdr3ocyqu` + "`" + ` ON ` + "`" + `payments` + "`" + ` (` + "`" + `stripe_event_id` + "`" + `) WHERE 'stripe_event_id' != '';"
			]
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
