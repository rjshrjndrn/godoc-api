package router

import (
	"context"
	"godoc/internal/datastructures/sql"
	"godoc/pkg/config"
	"godoc/pkg/db"
	"log"

	"github.com/davecgh/go-spew/spew"
)

type RouterImplimentation struct {
	Env *config.EnvConfig
	DB  db.DB
}

func NewRouter() (*RouterImplimentation, error) {
	env, err := config.ParseConfig()
	if err != nil {
		return nil, err
	}
	pgpool, err := db.New(env.Database.DBConnUrl)
	if err != nil {
		return nil, err
	}
	return &RouterImplimentation{
		Env: env,
		DB:  pgpool,
	}, nil
}

type PayloadCreatePatient struct {
	sql.PatientInfo
}

type addPatient struct {
	Body PayloadCreatePatient
}

type getPatient struct {
	FirstName string `path:"firstName" doc:"first name of patient"`
}

type generalResponse struct {
	Success bool
}

// Listener handle Post request and print addpatient
func (r *RouterImplimentation) AddPatient(ctx context.Context, input *addPatient) (*generalResponse, error) {
	patientInfo := input.Body
	if err := r.DB.CreatePatient(&patientInfo.PatientInfo); err != nil {
		spew.Dump(patientInfo)
		return nil, err
	}
	return &generalResponse{Success: true}, nil
}

type getPatientResponse struct {
	Body struct {
		Patients *[]sql.PatientInfo `json:"patients"`
	}
}

func (r *RouterImplimentation) GetPatient(ctx context.Context, input *getPatient) (*getPatientResponse, error) {
	log.Print(input.FirstName)
	data, err := r.DB.SearchPatient(input.FirstName)
	if err != nil {
		return nil, err
	}
	pi := &getPatientResponse{
		Body: struct {
			Patients *[]sql.PatientInfo `json:"patients"`
		}{Patients: data},
	}
	return pi, nil
}
