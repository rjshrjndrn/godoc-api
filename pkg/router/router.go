package router

import (
	"context"
	"godoc/internal/datastructures/sql"
	"godoc/pkg/config"
	"godoc/pkg/db"

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
	sql.CreatePatient
}

type addPatient struct {
	Body PayloadCreatePatient
}

type onboardResponse struct {
	Success bool
}

// Listener handle Post request and print addpatient
func (r *RouterImplimentation) AddPatient(ctx context.Context, input *addPatient) (*onboardResponse, error) {
	patientInfo := input.Body
	if err := r.DB.CreatePatient(patientInfo.CreatePatient); err != nil {
		spew.Dump(patientInfo)
		return nil, err
	}
	return &onboardResponse{Success: true}, nil
}
