package app

import (
	"errors"

	applicationshistory "github.com/ahmadabdelrazik/masarak/internal/domain/applicationsHistory"
	"github.com/ahmadabdelrazik/masarak/internal/domain/business"
	"github.com/ahmadabdelrazik/masarak/internal/domain/freelancerprofile"
	"github.com/ahmadabdelrazik/masarak/pkg/authuser"
)

var ErrDatabaseError = errors.New("database error")

type Repositories struct {
	Users              authuser.UserRepository
	Tokens             authuser.TokenRepository
	Businesses         business.Repository
	FreelancerProfile  freelancerprofile.Repository
	ApplicationHistory applicationshistory.Repository
}
