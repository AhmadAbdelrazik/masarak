package app

import (
	"context"

	"github.com/ahmadabdelrazik/masarak/pkg/filters"
)

type SearchFreelancerProfiles struct {
	Name               string
	Title              string
	Skills             []string
	YearsOfExperience  int
	HourlyRateAmount   int
	HourlyRateCurrency string

	// Filters has defaults of page number 1 and page size 20. the sort
	// safe list for should contain (name, title, years_of_experience,
	// hourly_rate_amount)
	Filters filters.Filter
}

// SearchFreelancerProfilesHandler finds freelancer profiles matching the
// search criteria. the nil value for any parameter means to select all except
// the numerical parameters (yearsOfExperience and hourlyRateAmount) with
// default value of -1.
func (q *Queries) SearchFreelancerProfilesHandler(
	ctx context.Context,
	cmd SearchFreelancerProfiles,
) ([]FreelancerProfile, filters.Metadata, error) {
	if cmd.Skills == nil {
		cmd.Skills = []string{}
	}

	profiles, meta, err := q.repo.FreelancerProfile.Search(
		ctx,
		cmd.Name,
		cmd.Title,
		cmd.Skills,
		cmd.YearsOfExperience,
		cmd.HourlyRateAmount,
		cmd.HourlyRateCurrency,
		cmd.Filters,
	)
	if err != nil {
		return nil, filters.Metadata{}, err
	}

	profilesDTO := make([]FreelancerProfile, 0, len(profiles))

	for _, profile := range profiles {
		p := toFreelancerProfile(&profile)
		profilesDTO = append(profilesDTO, p)
	}

	return profilesDTO, meta, nil
}
