package db

import (
	"context"
	"errors"
	"fmt"
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
        blood_group VARCHAR(10),
        phone_number VARCHAR(13),
        address TEXT
    );`
	_, err := pgpool.Exec(context.Background(), schema)
	return err
}

type DB interface {
	CreatePatient(*sql.PatientInfo) error
	SearchPatient(name string) (*[]sql.PatientInfo, error)
	ListPatients(limit, page int) (*[]sql.PatientInfo, error)
	Close()
}

func (d *DBImpl) Close() {
	d.Pool.Close()
}

func (d *DBImpl) CreatePatient(patient *sql.PatientInfo) error {
	sql := "INSERT INTO users (first_name, last_name, age, date_of_birth, blood_group, phone_number, address) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	if _, err := d.Pool.Exec(context.Background(), sql, patient.FirstName, patient.LastName, patient.Age, patient.DateOfBirth, patient.BloodGroup, patient.PhoneNumber, patient.Address); err != nil {
		return err
	}
	return nil
}

func (d *DBImpl) SearchPatient(name string) (*[]sql.PatientInfo, error) {
	query := "SELECT * FROM users WHERE first_name = $1"
	// search for all rows matching the name
	rows, err := d.Pool.Query(context.Background(), query, name)
	if err != nil {
		return nil, err
	}
	data, err := pgx.CollectRows(rows, pgx.RowToStructByName[sql.PatientInfo])
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// List all Patient with pagination
func (d *DBImpl) ListPatients(limit, page int) (*[]sql.PatientInfo, error) {
	if limit <= 0 || page <= 0 {
		return nil, errors.New("limit, page must be greater than 0")
	}
	page = page - 1
	query := fmt.Sprintf("SELECT * FROM users limit %d offset %d", limit, page*limit)
	// search for all rows matching the name
	rows, err := d.Pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	data, err := pgx.CollectRows(rows, pgx.RowToStructByName[sql.PatientInfo])
	if err != nil {
		return nil, err
	}
	return &data, nil
}
