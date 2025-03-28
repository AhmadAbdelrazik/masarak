package business

import (
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/entity"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/entity/job"
)

func (b *Business) GetJobsByApplicantEmail(email string) ([]job.Job, []entity.Application) {
	jobs := make([]job.Job, 0, len(b.jobs))
	applications := make([]entity.Application, 0, len(b.jobs))

	for _, j := range b.jobs {
		if application, err := j.GetApplicationByEmail(email); err == nil {
			jobs = append(jobs, *j)
			applications = append(applications, application)
		}
	}

	return jobs, applications
}
