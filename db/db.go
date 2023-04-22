package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
	"os"
)

func GetPoolFromUrl(dburl string, concurrent int) (*pgxpool.Pool, error) {
	dburl = fmt.Sprintf("%s&pool_max_conns=%d", dburl, concurrent)
	log.Debugf("Using dburl: '%s'", dburl)
	fmt.Printf("Using dburl: '%s'", dburl)
	pool, err := pgxpool.New(context.Background(), dburl)
	return pool, err
}

// GetPool Get a database pool
// postgres://root@127.0.0.1:26257?sslmode=disable&pool_max_conns=2000
func GetPool(dbIp string, concurrency int) *pgxpool.Pool {

	if concurrency == 0 {
		return nil
	}
	dburl := DbUrl(dbIp, concurrency)
	fmt.Println(dburl)
	dbpool, err := pgxpool.New(context.Background(), dburl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to databse : %v\n", err)
		os.Exit(1)
	}
	return dbpool
}

func DbUrl(dbIp string, concurrency int) string {
	return fmt.Sprintf("postgres://root@%s:26257?sslmode=disable&pool_max_conns=%d", dbIp, concurrency)
}
