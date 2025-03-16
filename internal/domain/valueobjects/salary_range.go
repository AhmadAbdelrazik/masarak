package valueobjects

import (
	"github.com/Rhymond/go-money"
	"github.com/pkg/errors"
)

type SalaryRange struct {
	from *money.Money
	to   *money.Money
}

func (s *SalaryRange) GetLowerRange() *money.Money {
	return s.from
}

func (s *SalaryRange) GetUpperRange() *money.Money {
	return s.to
}

// NewSalaryRange - salary values are in currency smallest unit
// 4 dollars --> 400 Cents --> 400
func NewSalaryRange(from, to int64, currency string) (*SalaryRange, error) {
	if from <= 0 || to <= 0 {
		return nil, errors.New("salary range must be positive value")
	}
	switch currency {
	case money.EGP:
	case money.SAR:
	case money.USD:
	case money.EUR:
	case money.AED:
	default:
		return nil, errors.New("unsupported currency")
	}

	return &SalaryRange{
		from: money.New(from, currency),
		to:   money.New(to, currency),
	}, nil
}
