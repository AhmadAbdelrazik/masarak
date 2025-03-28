package valueobject

import (
	"errors"
	"strings"
)

type JobStatus struct {
	status string
}

var (
	JobStatusClosed   = &JobStatus{status: "closed"}
	JobStatusOpen     = &JobStatus{status: "open"}
	JobStatusArchived = &JobStatus{status: "archived"}

	ErrInvalidJobStatus = errors.New("invalid job status")
)

func NewJobStatus(status string) (*JobStatus, error) {
	switch strings.ToLower(status) {
	case "open", "available":
		return JobStatusOpen, nil
	case "closed":
		return JobStatusClosed, nil
	case "archived":
		return JobStatusArchived, nil
	default:
		return nil, ErrInvalidJobStatus
	}
}

// Status - return the job status ("open", "closed", "archived")
func (s *JobStatus) Status() string {
	return s.status
}
