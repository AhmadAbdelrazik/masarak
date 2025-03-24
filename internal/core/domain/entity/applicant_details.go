package entity

import "github.com/Rhymond/go-money"

type ApplicantDetails struct {
	name              string
	email             string
	title             string
	yearsOfExperience int
	hourlyRate        *money.Money
}

func NewApplicantDetails(name, email, title string, yearsOfExperience int, hourlyRate *money.Money) *ApplicantDetails {
	return &ApplicantDetails{
		name:              name,
		email:             email,
		title:             title,
		yearsOfExperience: yearsOfExperience,
		hourlyRate:        hourlyRate,
	}
}
