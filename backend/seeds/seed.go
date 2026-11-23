package seeds

import (
	"fmt"
	"strconv"
	"strings"

	"ganigeorgiev.com/pocketbase/pocketbase/core"
	"ganigeorgiev.com/pocketbase/pocketbase/tools/filesystem"
	"ganigeorgiev.com/pocketbase/pocketbase/tools/types"
)

var password = "Testing12345"

func Run(app core.App) error {
	return app.RunInTransaction(func(txApp core.App) error {
		// 1. Seed users
		users, err := txApp.FindCollectionByNameOrId("users")
		if err != nil {
			return err
		}

		userIDs := make(map[string]string, 9)

		userEmails := []string{"outsider@demo.test", "adminA@demo.test", "boardA@demo.test", "resA1@demo.test", "resA2@demo.test", "formerA1@demo.test", "disabledA1@demo.test", "mixed@demo.test", "adminB@demo.test"}
		for _, email := range userEmails {
			record := core.NewRecord(users)
			record.Set("email", email)
			record.Set("verified", true)
			record.Set("password", password)
			if err := txApp.Save(record); err != nil {
				return fmt.Errorf("failed creating user %s: %w", email, err)
			}
			userIDs[email] = record.Id
		}

		// 2. Seed associations
		associations, err := txApp.FindCollectionByNameOrId("associations")
		if err != nil {
			return err
		}

		associationIDs := make(map[string]string, 2)

		for _, assoc := range []string{"A", "B"} {
			record := core.NewRecord(associations)
			record.Set("name", fmt.Sprintf("Association %s", assoc))
			record.Set("slug", fmt.Sprintf("association-%s", strings.ToLower(assoc)))
			record.Set("timezone", "America/Chicago")
			record.Set("currency", "usd")
			if err := txApp.Save(record); err != nil {
				return fmt.Errorf("failed creating asoociation %s: %w", assoc, err)
			}
			associationIDs[assoc] = record.Id
		}

		// 3. Seed properties
		properties, err := txApp.FindCollectionByNameOrId("properties")
		if err != nil {
			return err
		}

		propertyIDs := make(map[string]string, 3)

		for _, property := range []string{"P1", "P2", "P3"} {
			record := core.NewRecord(properties)
			if property == "P1" || property == "P2" {
				record.Set("association", associationIDs["A"])
			} else {
				record.Set("association", associationIDs["B"])
			}
			record.Set("address", fmt.Sprintf("%s Main Street", property))
			if err := txApp.Save(record); err != nil {
				return fmt.Errorf("failed creating property %s: %w", property, err)
			}
			propertyIDs[property] = record.Id
		}

		// 4. Seed memberships
		memberships, err := txApp.FindCollectionByNameOrId("memberships")
		if err != nil {
			return err
		}

		membershipIDs := make(map[string][]string, 9)

		for _, email := range userEmails {
			// Outsider and anon is a no-op
			if strings.Contains(strings.ToLower(email), strings.ToLower("outsider")) {
				continue
			}
			// Mixed member first
			if strings.Contains(strings.ToLower(email), strings.ToLower("mixed")) {
				// Mixed user membership in each association
				for _, assoc := range []string{"A", "B"} {
					record := core.NewRecord(memberships)
					record.Set("association", associationIDs[assoc])
					record.Set("user", userIDs[email])
					// Active admin in A
					if assoc == "A" {
						record.Set("role", "admin")
						record.Set("status", "active")
					} else {
						record.Set("role", "resident")
						record.Set("status", "disabled")
					}
					if err := txApp.Save(record); err != nil {
						return fmt.Errorf("failed creating membership for %s: %w", email, err)
					}
					membershipIDs[email] = append(membershipIDs[email], record.Id)
				}
			} else {
				record := core.NewRecord(memberships)
				lowerEmail := strings.ToLower(email)
				record.Set("user", userIDs[email])

				switch {
				case strings.Contains(lowerEmail, "admin"):
					if strings.Contains(email, "A") {
						record.Set("association", associationIDs["A"])
					} else {
						record.Set("association", associationIDs["B"])
					}
					record.Set("user", userIDs[email])
					record.Set("status", "active")
					record.Set("role", "admin")
				case strings.Contains(lowerEmail, "board"):
					record.Set("user", userIDs[email])
					record.Set("association", associationIDs["A"])
					record.Set("status", "active")
					record.Set("role", "board")
				case strings.Contains(lowerEmail, "disabled"):
					record.Set("user", userIDs[email])
					record.Set("association", associationIDs["A"])
					record.Set("status", "disabled")
					record.Set("role", "resident")
				case strings.Contains(lowerEmail, "former"):
					record.Set("user", userIDs[email])
					record.Set("association", associationIDs["A"])
					record.Set("status", "active")
					record.Set("role", "resident")
				default:
					// Resident of single association - both in A
					record.Set("user", userIDs[email])
					record.Set("association", associationIDs["A"])
					record.Set("status", "active")
					record.Set("role", "resident")
				}

				if err := txApp.Save(record); err != nil {
					return fmt.Errorf("failed creating membership for %s: %w", email, err)
				}
				membershipIDs[email] = append(membershipIDs[email], record.Id)
			}
		}

		// 5. Seed residents
		residents, err := txApp.FindCollectionByNameOrId("residents")
		if err != nil {
			return err
		}

		// resA1 and resA2
		for i := 1; i < 3; i++ {
			idStr := strconv.Itoa(i)
			record := core.NewRecord(residents)
			property := "P" + idStr
			email := "resA" + idStr + "@demo.test"
			record.Set("user", userIDs[email])
			record.Set("association", associationIDs["A"])
			record.Set("property", propertyIDs[property])
			record.Set("start_date", types.NowDateTime().AddDate(0, 0, -80))
			record.Set("name", email)
			if err := txApp.Save(record); err != nil {
				return fmt.Errorf("failed creating resident record for %s: %w", email, err)
			}
		}

		// formerA1 - end_date set
		formerA1Record := core.NewRecord(residents)
		formerA1Record.Set("user", userIDs["formerA1@demo.test"])
		formerA1Record.Set("property", propertyIDs["P1"])
		formerA1Record.Set("association", associationIDs["A"])
		formerA1Record.Set("start_date", types.NowDateTime().AddDate(0, 0, -80))
		formerA1Record.Set("end_date", types.NowDateTime().AddDate(0, 0, -20))
		formerA1Record.Set("name", "formerA1@demo.test")
		if err := txApp.Save(formerA1Record); err != nil {
			return fmt.Errorf("failed creating resident record for %s: %w", "formerA1@demo.test", err)
		}

		// disabledA1
		disabledA1Record := core.NewRecord(residents)
		disabledA1Record.Set("user", userIDs["disabledA1@demo.test"])
		disabledA1Record.Set("association", associationIDs["A"])
		disabledA1Record.Set("property", propertyIDs["P1"])
		disabledA1Record.Set("start_date", types.NowDateTime().AddDate(0, 0, -80))
		disabledA1Record.Set("name", "disabledA1@demo.test")
		if err := txApp.Save(disabledA1Record); err != nil {
			return fmt.Errorf("failed creating resident record for %s: %w", "disabledA1@demo.test", err)
		}

		// No user resident - simulates an invite
		noUserRecord := core.NewRecord(residents)
		noUserRecord.Set("association", associationIDs["A"])
		noUserRecord.Set("property", propertyIDs["P2"])
		noUserRecord.Set("name", "Bill")
		if err := txApp.Save(noUserRecord); err != nil {
			return fmt.Errorf("failed creating resident record for %s: %w", "No user record", err)
		}

		// 6. See assessment schedules
		schedules, err := txApp.FindCollectionByNameOrId("assessment_schedules")
		if err != nil {
			return err
		}

		scheduleData := map[string][]map[string]any{
			"A": {
				{
					"name":         "Monthly dues",
					"amount_cents": 25000,
					"frequency":    "monthly",
					"start_date":   "2026-01-01",
					"end_date":     nil,
					"active":       true,
				},
				{
					"name":         "Special assessment: roof repair",
					"amount_cents": 150000,
					"frequency":    "once",
					"start_date":   "2026-03-01",
					"end_date":     nil,
					"active":       true,
				},
				{
					"name":         "Annual landscaping fee",
					"amount_cents": 36000,
					"frequency":    "annual",
					"start_date":   "2026-01-01",
					"end_date":     nil,
					"active":       true,
				},
				{
					"name":         "Old pool levy",
					"amount_cents": 5000,
					"frequency":    "monthly",
					"start_date":   "2024-01-01",
					"end_date":     "2025-12-31",
					"active":       false,
				},
			},
			"B": {
				{
					"name":         "Quarterly dues",
					"amount_cents": 90000,
					"frequency":    "quarterly",
					"start_date":   "2026-01-01",
					"end_date":     nil,
					"active":       true,
				},
				{
					"name":         "Monthly dues (paused)",
					"amount_cents": 30000,
					"frequency":    "monthly",
					"start_date":   "2026-01-01",
					"end_date":     nil,
					"active":       false,
				},
			},
		}

		fmt.Println(scheduleData)

		scheduleIDs := make(map[string]string, 6)

		for assoc, fees := range scheduleData {
			for _, feeData := range fees {
				record := core.NewRecord(schedules)
				for field, value := range feeData {
					record.Set(field, value)
				}
				record.Set("association", associationIDs[assoc])
				if err := txApp.Save(record); err != nil {
					return fmt.Errorf("failed creating assessment schedule record for association %s: %w", assoc, err)
				}
				scheduleIDs[assoc+"/"+record.GetString("name")] = record.Id
			}
		}

		// 7. Seed property notes
		notes, err := txApp.FindCollectionByNameOrId("property_notes")
		if err != nil {
			return err
		}

		noteData := []map[string]any{
			{
				"association": associationIDs["A"],
				"property":    propertyIDs["P1"],
				"author":      userIDs["boardA@demo.test"],
				"note":        "<p>Gate code was changed in January. Resident has been told.</p>",
			},
			{
				// No author - simulates a note whose author account was deleted
				"association": associationIDs["A"],
				"property":    propertyIDs["P1"],
				"note":        "<p>Previous owner left a payment plan arrangement on file.</p>",
			},
			{
				"association": associationIDs["A"],
				"property":    propertyIDs["P2"],
				"author":      userIDs["adminA@demo.test"],
				"note":        "<p>Fence repair approved by the board, pending contractor quote.</p>",
			},
			{
				"association": associationIDs["B"],
				"property":    propertyIDs["P3"],
				"author":      userIDs["adminB@demo.test"],
				"note":        "<p>Owner requested paperless statements.</p>",
			},
		}

		for _, data := range noteData {
			if _, err := seedRecord(txApp, notes, data); err != nil {
				return err
			}
		}

		// 8. Seed announcements
		announcements, err := txApp.FindCollectionByNameOrId("announcements")
		if err != nil {
			return err
		}

		announcementData := []map[string]any{
			{
				"association": associationIDs["A"],
				"author":      userIDs["boardA@demo.test"],
				"title":       "Annual meeting date set",
				"body":        "<p>The annual meeting will be held on the first Saturday of March.</p>",
				"pinned":      true,
			},
			{
				"association": associationIDs["A"],
				"author":      userIDs["adminA@demo.test"],
				"title":       "Pool opens in May",
				"body":        "<p>The pool will open Memorial Day weekend. Keys are available from the office.</p>",
				"pinned":      false,
			},
			{
				"association": associationIDs["B"],
				"author":      userIDs["adminB@demo.test"],
				"title":       "Welcome to Association B",
				"body":        "<p>Dues are billed quarterly. Please set up online payment.</p>",
				"pinned":      false,
			},
		}

		for _, data := range announcementData {
			if _, err := seedRecord(txApp, announcements, data); err != nil {
				return err
			}
		}

		// 9. Seed documents
		documents, err := txApp.FindCollectionByNameOrId("documents")
		if err != nil {
			return err
		}

		documentData := []map[string]any{
			{
				"association": associationIDs["A"],
				"title":       "Bylaws",
				"category":    "bylaws",
				"visibility":  "all",
				"uploaded_by": userIDs["adminA@demo.test"],
				"filename":    "bylaws.pdf",
			},
			{
				"association": associationIDs["A"],
				"title":       "Architectural request form",
				"category":    "forms",
				"visibility":  "all",
				"uploaded_by": userIDs["boardA@demo.test"],
				"filename":    "architectural-request-form.pdf",
			},
			{
				"association": associationIDs["A"],
				"title":       "Board meeting minutes",
				"category":    "minutes",
				"visibility":  "board",
				"uploaded_by": userIDs["boardA@demo.test"],
				"filename":    "board-minutes.pdf",
			},
			{
				"association": associationIDs["A"],
				"title":       "Reserve study",
				"category":    "financial",
				"visibility":  "board",
				"uploaded_by": userIDs["adminA@demo.test"],
				"filename":    "reserve-study.pdf",
			},
			{
				"association": associationIDs["B"],
				"title":       "Welcome packet",
				"category":    "other",
				"visibility":  "all",
				"uploaded_by": userIDs["adminB@demo.test"],
				"filename":    "welcome-packet.pdf",
			},
		}

		for _, data := range documentData {
			file, err := filesystem.NewFileFromBytes(seedPDF, data["filename"].(string))
			if err != nil {
				return fmt.Errorf("failed creating file for document %v: %w", data["title"], err)
			}
			delete(data, "filename")
			data["file"] = file
			if _, err := seedRecord(txApp, documents, data); err != nil {
				return err
			}
		}

		// 10. Seed payments
		payments, err := txApp.FindCollectionByNameOrId("payments")
		if err != nil {
			return err
		}

		paymentSeeds := []struct {
			key  string
			data map[string]any
		}{
			{
				key: "P1-2026-01",
				data: map[string]any{
					"association":           associationIDs["A"],
					"property":              propertyIDs["P1"],
					"payer":                 userIDs["resA1@demo.test"],
					"amount_cents":          25000,
					"fee_cents":             755,
					"method":                "card",
					"status":                "succeeded",
					"stripe_payment_intent": "pi_seed_p1_jan",
					"stripe_event_id":       "evt_seed_p1_jan",
					"received_at":           "2026-01-05",
				},
			},
			{
				key: "P1-2026-02",
				data: map[string]any{
					"association":  associationIDs["A"],
					"property":     propertyIDs["P1"],
					"payer":        userIDs["resA1@demo.test"],
					"amount_cents": 25000,
					"fee_cents":    0,
					"method":       "check",
					"status":       "succeeded",
					"reference":    "Check #1042",
					"received_at":  "2026-02-08",
				},
			},
			{
				key: "P2-2026-01",
				data: map[string]any{
					"association":           associationIDs["A"],
					"property":              propertyIDs["P2"],
					"payer":                 userIDs["resA2@demo.test"],
					"amount_cents":          25000,
					"fee_cents":             200,
					"method":                "ach",
					"status":                "succeeded",
					"stripe_payment_intent": "pi_seed_p2_jan",
					"stripe_event_id":       "evt_seed_p2_jan",
					"received_at":           "2026-01-10",
				},
			},
			{
				// Failed payment - no ledger entry is written for this one
				key: "P2-2026-02-failed",
				data: map[string]any{
					"association":           associationIDs["A"],
					"property":              propertyIDs["P2"],
					"payer":                 userIDs["resA2@demo.test"],
					"amount_cents":          25000,
					"fee_cents":             0,
					"method":                "card",
					"status":                "failed",
					"stripe_payment_intent": "pi_seed_p2_feb_failed",
					"stripe_event_id":       "evt_seed_p2_feb_failed",
					"received_at":           "2026-02-12",
				},
			},
			{
				key: "P3-2026-Q1",
				data: map[string]any{
					"association":           associationIDs["B"],
					"property":              propertyIDs["P3"],
					"payer":                 userIDs["adminB@demo.test"],
					"amount_cents":          90000,
					"fee_cents":             2640,
					"method":                "card",
					"status":                "succeeded",
					"stripe_payment_intent": "pi_seed_p3_q1",
					"stripe_event_id":       "evt_seed_p3_q1",
					"received_at":           "2026-01-20",
				},
			},
		}

		paymentIDs := make(map[string]string, len(paymentSeeds))

		for _, p := range paymentSeeds {
			id, err := seedRecord(txApp, payments, p.data)
			if err != nil {
				return err
			}
			paymentIDs[p.key] = id
		}

		// 11. Seed ledger entries (charges are positive, payments and credits are negative)
		ledger, err := txApp.FindCollectionByNameOrId("ledger_entries")
		if err != nil {
			return err
		}

		propertyAssoc := map[string]string{"P1": "A", "P2": "A", "P3": "B"}

		addEntry := func(property, entryType string, amount int, postedOn, description string, extra map[string]any) (string, error) {
			data := map[string]any{
				"association":  associationIDs[propertyAssoc[property]],
				"property":     propertyIDs[property],
				"type":         entryType,
				"amount_cents": amount,
				"posted_on":    postedOn,
				"description":  description,
			}
			for field, value := range extra {
				data[field] = value
			}
			return seedRecord(txApp, ledger, data)
		}

		monthlyDuesA := scheduleIDs["A/Monthly dues"]
		roofRepairA := scheduleIDs["A/Special assessment: roof repair"]
		quarterlyDuesB := scheduleIDs["B/Quarterly dues"]

		// Monthly dues for P1 and P2, January to March
		chargeIDs := make(map[string]string, 6)

		for _, property := range []string{"P1", "P2"} {
			for _, month := range []string{"01", "02", "03"} {
				period := "2026-" + month
				id, err := addEntry(property, "charge", 25000, period+"-01", "Monthly dues "+period, map[string]any{
					"due_on":   period + "-15",
					"schedule": monthlyDuesA,
					"period":   period,
				})
				if err != nil {
					return err
				}
				chargeIDs[property+"-"+month] = id
			}
		}

		// Special assessment for P1 and P2
		for _, property := range []string{"P1", "P2"} {
			if _, err := addEntry(property, "charge", 150000, "2026-03-01", "Special assessment: roof repair", map[string]any{
				"due_on":   "2026-03-31",
				"schedule": roofRepairA,
				"period":   "2026-03",
			}); err != nil {
				return err
			}
		}

		// Payments received - each pairs a negative ledger entry with its payment record
		paymentEntries := []struct {
			property string
			key      string
			amount   int
			postedOn string
			label    string
		}{
			{"P1", "P1-2026-01", -25000, "2026-01-05", "Payment received (card)"},
			{"P1", "P1-2026-02", -25000, "2026-02-08", "Payment received (check #1042)"},
			{"P2", "P2-2026-01", -25000, "2026-01-10", "Payment received (ACH)"},
		}

		for _, p := range paymentEntries {
			if _, err := addEntry(p.property, "payment", p.amount, p.postedOn, p.label, map[string]any{
				"payment": paymentIDs[p.key],
			}); err != nil {
				return err
			}
		}

		// Late fee on P2 after the failed February payment
		if _, err := addEntry("P2", "late_fee", 2500, "2026-02-16", "Late fee: February dues unpaid", nil); err != nil {
			return err
		}

		// Courtesy credit on P1
		if _, err := addEntry("P1", "credit", -2000, "2026-02-20", "Courtesy credit", nil); err != nil {
			return err
		}

		// Reversal of P2's March dues, billed in error. Points at the original charge.
		if _, err := addEntry("P2", "reversal", -25000, "2026-03-05", "Reversal: March dues billed in error", map[string]any{
			"reverses": chargeIDs["P2-03"],
		}); err != nil {
			return err
		}

		// Association B: quarterly dues, paid by card
		if _, err := addEntry("P3", "charge", 90000, "2026-01-01", "Quarterly dues 2026-Q1", map[string]any{
			"due_on":   "2026-01-31",
			"schedule": quarterlyDuesB,
			"period":   "2026-Q1",
		}); err != nil {
			return err
		}

		if _, err := addEntry("P3", "payment", -90000, "2026-01-20", "Payment received (card)", map[string]any{
			"payment": paymentIDs["P3-2026-Q1"],
		}); err != nil {
			return err
		}

		return nil
	})
}

// seedPDF is a minimal valid PDF so the documents file field accepts it.
var seedPDF = []byte("%PDF-1.4\n1 0 obj<</Type/Catalog/Pages 2 0 R>>endobj\n2 0 obj<</Type/Pages/Kids[3 0 R]/Count 1>>endobj\n3 0 obj<</Type/Page/Parent 2 0 R/MediaBox[0 0 200 200]>>endobj\ntrailer<</Root 1 0 R>>\n%%EOF")

// seedRecord creates and saves a record, returning its id.
// It uses Set (not Load) so each value is cast by its field type.
func seedRecord(txApp core.App, collection *core.Collection, data map[string]any) (string, error) {
	record := core.NewRecord(collection)
	for field, value := range data {
		record.Set(field, value)
	}
	if err := txApp.Save(record); err != nil {
		return "", fmt.Errorf("failed creating %s record: %w", collection.Name, err)
	}
	return record.Id, nil
}
