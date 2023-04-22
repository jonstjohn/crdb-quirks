package quirks

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jonstjohn/crdb-quirks/generator"
	"sync"
	"time"
)

const GarbageCreateDatabaseSql = "CREATE DATABASE IF NOT EXISTS garbage"
const GarbageCreateTableSql = `
CREATE TABLE IF NOT EXISTS garbage.workflow (
  id uuid primary key,
  description string,
  expires timestamp)
`
const GarbageUpsertSql = "UPSERT INTO garbage.workflow (id, description, expires) VALUES ($1, $2, now() + INTERVAL '1m')"

func RunGarbageWorkload(pool *pgxpool.Pool, upsertWorkers int) {

	// Setup schema
	setupGarbageSchema(pool)

	// Upsert garbage
	generator := generator.New(0)

	var wg sync.WaitGroup
	for i := 0; i < upsertWorkers; i++ {

		wg.Add(1)

		id, err := generator.UUID()
		if err != nil {
			panic(err)
		}
		go func(id string, description string) {
			for {
				upsertGarbage(pool, id, id)
				time.Sleep(100 * time.Millisecond)
			}
			wg.Done()
		}(id, id)
	}

	// Read garbage

	wg.Wait()

}

func setupGarbageSchema(pool *pgxpool.Pool) {
	// Create database
	pool.Exec(context.Background(), GarbageCreateDatabaseSql)

	// Create table
	pool.Exec(context.Background(), GarbageCreateTableSql)
}

func upsertGarbage(pool *pgxpool.Pool, id string, description string) {
	exec, err := pool.Exec(context.Background(), GarbageUpsertSql, id, description)
	if err != nil {
		panic(err)
	}
	if exec.RowsAffected() != 1 {
		panic(fmt.Errorf("upsert failed"))
	}
}
