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

func (s *JobStatus) Status() string {
	return s.status
}

func (s *JobStatus) IsOpen() bool {
	return s == JobStatusOpen
}

func (s *JobStatus) IsClosed() bool {
	return s == JobStatusClosed
}

func (s *JobStatus) IsArchived() bool {
	return s == JobStatusArchived
}
