package app

import (
	applicationshistory "github.com/ahmadabdelrazik/masarak/internal/domain/applicationsHistory"
	"github.com/ahmadabdelrazik/masarak/internal/domain/business"
	"github.com/ahmadabdelrazik/masarak/internal/domain/freelancerprofile"
	"github.com/ahmadabdelrazik/masarak/pkg/authuser"
)

type Repositories struct {
	Users              authuser.UserRepository
	Tokens             authuser.TokenRepository
	Businesses         business.Repository
	FreelancerProfile  freelancerprofile.Repository
	ApplicationHistory applicationshistory.Repository
}
