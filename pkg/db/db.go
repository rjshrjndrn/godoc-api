package db

import (
	"context"
	"godoc/internal/datastructures/sql"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// New Generate db connection and return it
func New(connUrl string) (DB, error) {
	config, err := pgxpool.ParseConfig(connUrl)
	if err != nil {
		log.Fatalf("Unable to parse connection string: %v", err)
	}
	// https://github.com/jackc/pgx/issues/1847#issuecomment-2219737645
	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeCacheDescribe
	dbpool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}
	if err = initDb(dbpool); err != nil {
		return nil, err
	}
	return &DBImpl{Pool: dbpool}, nil
}

type DBImpl struct {
	Pool *pgxpool.Pool
}

func initDb(pgpool *pgxpool.Pool) error {
	schema := `
    CREATE TABLE IF NOT EXISTS users (
        user_id BIGSERIAL PRIMARY KEY,
        first_name VARCHAR(50),
        last_name VARCHAR(50),
        age INT,
        date_of_birth DATE,
        blood_group VARCHAR(5),
        phone_number BIGINT,
        address TEXT
    );`
	_, err := pgpool.Exec(context.Background(), schema)
	return err
}

type DB interface {
	CreatePatient(*sql.CreatePatient) error
	Close()
}

func (d *DBImpl) Close() {
	d.Pool.Close()
}

func (d *DBImpl) CreatePatient(patient *sql.CreatePatient) error {
	sql := "INSERT INTO users (first_name, last_name, age, date_of_birth, blood_group, phone_number, address) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	if _, err := d.Pool.Exec(context.Background(), sql, patient.FirstName, patient.LastName, patient.Age, patient.DateOfBirth, patient.BloodGroup, patient.PhoneNumber, patient.Address); err != nil {
		return err
	}
	return nil
}
