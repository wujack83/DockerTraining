package database

import (
	"testing"
	"gitlab.com/andersph/docker-master-class/api/internal"
)

func TestDbWriter(t *testing.T) {
	
	// load configuration
    c := internal.LoadConfig()

	if internal.GetEnv("IS_STDOUT", c.Flags.Is_StdOut) != "true" {
		var employee internal.Employee

		employee.Name = "Max Mustermann"
		employee.Address = "01234 Musterhausen, Musterstraße 1a"
		employee.Mail = "max@mustermann.org"
		employee.DateOfBirth = "1989-11-11"
		employee.Department = "Musterdepartment"
		employee.JobTitle = "Musterjob"

		SqlWriter(employee)
	} else {
		t.Errorf("Test not implemented for std_out option")
	}
}
