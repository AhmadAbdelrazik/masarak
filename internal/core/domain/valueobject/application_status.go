package valueobject

type ApplicationStatus struct {
	status string
}

func NewApplicationStatus(status string) *ApplicationStatus {
	return &ApplicationStatus{
		status: status,
	}
}

func (s *ApplicationStatus) Status() string {
	return s.status
}

func (s *ApplicationStatus) IsAccepted() bool {
	return s.status == "accepted"
}

func (s *ApplicationStatus) IsPending() bool {
	return s.status == "pending"
}

func (s *ApplicationStatus) IsRejected() bool {
	return s.status == "rejected"
}
