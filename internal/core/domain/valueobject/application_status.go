package valueobject

import (
	"errors"
	"strings"
)

type ApplicationStatus struct {
	status string
}

var (
	applicationStatusPending  = &ApplicationStatus{status: "pending"}
	applicationStatusAccepted = &ApplicationStatus{status: "accepted"}
	applicationStatusRejected = &ApplicationStatus{status: "rejected"}

	ErrInvalidApplicationStatus = errors.New("invalid application status")
)

func NewApplicationStatus(status string) (*ApplicationStatus, error) {
	switch strings.ToLower(status) {
	case "pending":
		return applicationStatusPending, nil
	case "accepted", "accept":
		return applicationStatusAccepted, nil
	case "rejected", "reject":
		return applicationStatusRejected, nil
	default:
		return nil, ErrInvalidApplicationStatus
	}
}

func (s *ApplicationStatus) Status() string {
	return s.status
}

func (s *ApplicationStatus) IsPending() bool {
	return s == applicationStatusPending
}

func (s *ApplicationStatus) IsAccepted() bool {
	return s == applicationStatusAccepted
}

func (s *ApplicationStatus) IsRejected() bool {
	return s == applicationStatusRejected
}
