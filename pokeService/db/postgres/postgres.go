package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Enviy/pokedexProject/pokeService/entity"
	"github.com/Enviy/pokedexProject/pokeService/utils/goutils"

	// _ blank import temporary
	_ "github.com/lib/pq"
)

// For better testing
var (
	openDBConnection = sql.Open
)

// Postgres is the interface to the PostgreSQL DB
type Postgres interface {
}

type postgres struct {
	secrets entity.Secrets
	db      *sql.DB
}

// New provides a new Postgres Interface to Fx
func New(secrets entity.Secrets) (Postgres, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		secrets.PostgresInfo.Host,
		secrets.PostgresInfo.Port,
		secrets.PostgresInfo.User,
		secrets.PostgresInfo.Password,
		secrets.PostgresInfo.Name,
	)

	// Create the connection to postgresql server
	db, err := openDBConnection("postgres", psqlInfo)
	if err != nil {
		return nil, goutils.ErrWrap(err)
	}

	return &postgres{
		secrets: secrets,
		db:      db,
	}, nil
}
