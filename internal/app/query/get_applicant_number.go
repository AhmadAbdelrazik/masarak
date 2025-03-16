package query

import (
	"context"

	"github.com/ahmadabdelrazik/linkedout/internal/domain/applicant"
)

type GetApplicantNumber struct{}

type GetApplicantNumberHandler struct {
	applicantRepo applicant.Repository
}

func NewGetApplicantNumberHandler(applicantRepo applicant.Repository) *GetApplicantNumberHandler {
	return &GetApplicantNumberHandler{
		applicantRepo: applicantRepo,
	}
}

func (h *GetApplicantNumberHandler) Handle(ctx context.Context, cmd GetApplicantNumber) (int, error) {
	applicantsNumber, err := h.applicantRepo.GetApplicantNumbers(ctx)
	if err != nil {
		return 0, err
	}

	return applicantsNumber, nil
}
