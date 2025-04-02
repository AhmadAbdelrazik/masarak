package business

import (
	"errors"
	"slices"
)

var (
	ErrEmployeeNotFound      = errors.New("employee not found")
	ErrEmployeeAlreadyExists = errors.New("employee already exists")
)

func (b *Business) AddEmployee(email string) error {
	for _, e := range b.employeeEmails {
		if e == email {
			return ErrEmployeeAlreadyExists
		}
	}

	b.employeeEmails = append(b.employeeEmails, email)

	return nil
}

func (b *Business) RemoveEmployee(email string) error {
	for i, e := range b.employeeEmails {
		if e == email {
			b.employeeEmails = slices.Delete(b.employeeEmails, i, i+1)
			return nil
		}
	}

	return ErrEmployeeNotFound
}

func (b *Business) IsEmployee(email string) bool {
	if b.ownerEmail == email {
		return true
	}

	for _, e := range b.employeeEmails {
		if e == email {
			return true
		}
	}

	return false
}
