package dasorm

import (
	"fmt"

	"github.com/pkg/errors"
)

func connectMySQL(creds *Config) (*Connection, error) {
	connectionURL := fmt.Sprintf("%s:%s@(%s)/%s?parseTime=true", creds.User, creds.Password, creds.Host, creds.Database)
	db, err := connectURL(mysqlDialect, connectionURL)
	if err != nil {
		return nil, errors.Wrap(err, "connect mysql")
	}
	return &Connection{
		DB:      &DB{DB: db},
		Dialect: &mysql{},
	}, nil
}

type mysql struct{}

func (m *mysql) Name() string {
	return "mysql"
}

func (m *mysql) TranslateSQL(sql string) string {
	return sql
}

func (m *mysql) Create(db DBInterface, model *Model) error {
	return errors.Wrap(genericCreate(db, model), "mysql create")
}

func (m *mysql) CreateMany(db DBInterface, model *Model) error {
	return errors.Wrap(genericCreateMany(db, model), "mysql create")
}

func (m *mysql) Update(db DBInterface, model *Model) error {
	return errors.Wrap(genericUpdate(db, model), "mysql update")
}

func (m *mysql) Destroy(db DBInterface, model *Model) error {
	return errors.Wrap(genericDestroy(db, model), "mysql destroy")
}

func (m *mysql) DestroyMany(db DBInterface, model *Model) error {
	return errors.Wrap(genericDestroyMany(db, model), "mysql destroy many")
}

func (m *mysql) SelectOne(db DBInterface, model *Model, query Query) error {
	return errors.Wrap(genericSelectOne(db, model, query), "mysql select one")
}

func (m *mysql) SelectMany(db DBInterface, models *Model, query Query) error {
	return errors.Wrap(genericSelectMany(db, models, query), "mysql select many")
}

func (m *mysql) SQLView(db DBInterface, model *Model, format map[string]string) error {
	return errors.Wrap(genericSQLView(db, model, format), "mysql sql view")
}

func (m *mysql) CreateUpdate(db DBInterface, model *Model) error {
	return errors.Wrap(genericCreateUpdate(db, model), "mysql create update")
}

func (m *mysql) CreateManyTemp(DBInterface, *Model) error {
	return ErrNotImplemented
}

func (m *mysql) CreateManyUpdate(db DBInterface, model *Model) error {
	return errors.Wrap(genericCreateManyUpdate(db, model), "mysql create update many")
}

func (m *mysql) Truncate(db DBInterface, model *Model) error {
	return errors.Wrap(genericTruncate(db, model), "mysql truncate")
}
