package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"ganigeorgiev.com/pocketbase/pocketbase"
	"ganigeorgiev.com/pocketbase/pocketbase/core"
	"ganigeorgiev.com/pocketbase/pocketbase/plugins/migratecmd"
	"ganigeorgiev.com/pocketbase/pocketbase/tools/osutils"
	"github.com/spf13/cobra"

	_ "homebase/migrations"
	"homebase/seeds"
)

func main() {
	app := pocketbase.New()

	var skipConfirm, includeSystem bool

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		// enable auto creation of migration files when making collection changes in the Dashboard
		// (the IsProbablyGoRun check is to enable it only during development)
		Automigrate: osutils.IsProbablyGoRun(),
	})

	// Register custom `seed` command
	app.RootCmd.AddCommand(&cobra.Command{
		Use:   "seed",
		Short: "Populate the database with initial seed data",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("Seeding database...")
			if err := seeds.Run(app); err != nil {
				return err
			}
			log.Println("Database seeded successfully!")
			return nil
		},
	})

	app.RootCmd.AddCommand(&cobra.Command{
		Use:   "truncate-all",
		Short: "Deletes ALL records from every base/auth collection (schemas are kept)",
		Run: func(cmd *cobra.Command, args []string) {
			if err := truncateAll(app, skipConfirm, includeSystem); err != nil {
				log.Fatal(err)
			}
		},
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

func truncateAll(app core.App, skipConfirm, includeSystem bool) error {
	all, err := app.FindAllCollections()
	if err != nil {
		return fmt.Errorf("failed to list collections: %w", err)
	}

	var targets []*core.Collection
	for _, c := range all {
		if c.IsView() { // views have no rows of their own
			continue
		}
		if c.System && !includeSystem {
			continue
		}
		targets = append(targets, c)
	}

	if len(targets) == 0 {
		fmt.Println("Nothing to truncate.")
		return nil
	}

	fmt.Println("The following collections will be truncated:")
	for _, c := range targets {
		n, _ := app.CountRecords(c.Name)
		fmt.Printf("  - %-30s (%d records)\n", c.Name, n)
	}

	if !skipConfirm {
		fmt.Print("\nThis cannot be undone. Type 'yes' to continue: ")
		line, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		if strings.TrimSpace(line) != "yes" {
			fmt.Println("Aborted.")
			return nil
		}
	}

	// Truncating can fail while other records still reference the collection
	// (e.g. required relations without cascade delete), so retry the failures
	// in multiple passes until no more progress is made.
	pending := targets
	for pass := 1; len(pending) > 0; pass++ {
		var failed []*core.Collection
		var lastErr error

		for _, c := range pending {
			if err := app.TruncateCollection(c); err != nil {
				failed = append(failed, c)
				lastErr = fmt.Errorf("%s: %w", c.Name, err)
				continue
			}
			fmt.Printf("truncated %s\n", c.Name)
		}

		if len(failed) == len(pending) {
			return fmt.Errorf("could not truncate %d collection(s), last error: %w", len(failed), lastErr)
		}
		pending = failed
	}

	fmt.Println("Done.")
	return nil
}
