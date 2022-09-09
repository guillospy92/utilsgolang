package celeritas

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func (a *Accelerator) Migrate(dns string) error {
	m, err := migrate.New("file://"+a.RootPath+"/migrations", dns)

	fmt.Println("file://"+a.RootPath+"/migrations", dns)

	if err != nil {
		fmt.Println("error aqui", err)
		return err
	}

	defer func(m *migrate.Migrate) {
		err, _ := m.Close()
		if err != nil {
			a.ErrorLog.Println("error logs")
		}
	}(m)

	if err = m.Up(); err != nil {
		return err
	}

	return nil
}

func (a *Accelerator) MigrateDownAll(dns string) error {
	m, err := migrate.New("file://"+a.RootPath+"/migrations", dns)
	if err != nil {
		return err
	}

	defer func(m *migrate.Migrate) {
		err, _ := m.Close()
		if err != nil {
			a.ErrorLog.Println("error logs")
		}
	}(m)

	if err := m.Down(); err != nil {
		return err
	}

	return nil
}

func (a *Accelerator) Steps(n int, dns string) error {
	m, err := migrate.New("file://"+a.RootPath+"/migrations", dns)
	if err != nil {
		return err
	}

	defer func(m *migrate.Migrate) {
		err, _ := m.Close()
		if err != nil {
			a.ErrorLog.Println("error logs")
		}
	}(m)

	if err := m.Steps(n); err != nil {
		return err
	}

	return nil
}

func (a *Accelerator) MigrateForce(dns string) error {
	m, err := migrate.New("file://"+a.RootPath+"/migrations", dns)
	if err != nil {
		return err
	}

	defer func(m *migrate.Migrate) {
		err, _ := m.Close()
		if err != nil {
			a.ErrorLog.Println("error logs")
		}
	}(m)

	if err := m.Force(-1); err != nil {
		return err
	}

	return nil
}
