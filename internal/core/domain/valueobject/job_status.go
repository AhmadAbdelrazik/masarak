package valueobject

type JobStatus struct {
	status string
}

func NewJobStatus(status string) *JobStatus {
	return &JobStatus{
		status: status,
	}
}

func (s *JobStatus) Status() string {
	return s.status
}

func (s *JobStatus) IsAvailable() bool {
	return s.status == "available"
}
