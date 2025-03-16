package query

import (
	"context"

	"github.com/ahmadabdelrazik/linkedout/internal/domain/applicant"
)

type GetApplicant struct {
	Email string
}

type GetApplicantHandler struct {
	applicantRepo applicant.Repository
}

func NewGetApplicantHandler(applicantRepo applicant.Repository) *GetApplicantHandler {
	return &GetApplicantHandler{
		applicantRepo: applicantRepo,
	}
}

func (h *GetApplicantHandler) Handle(ctx context.Context, cmd GetApplicant) (*applicant.Applicant, error) {
	a, err := h.applicantRepo.GetByEmail(ctx, cmd.Email)
	if err != nil {
		return nil, err
	}

	return a, nil
}
