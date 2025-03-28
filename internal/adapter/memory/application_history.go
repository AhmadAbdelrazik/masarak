package memory

import (
	"context"

	applicationshistory "github.com/ahmadabdelrazik/masarak/internal/core/domain/applicationsHistory"
)

type InMemoryApplicationHistoryRepository struct {
	memory *Memory
}

func NewInMemoryApplicationHistoryRepository(mem *Memory) *InMemoryApplicationHistoryRepository {
	return &InMemoryApplicationHistoryRepository{
		memory: mem,
	}
}

func (r *InMemoryApplicationHistoryRepository) GetByEmail(
	ctx context.Context,
	email string,
) (*applicationshistory.ApplicationHistory, error) {
	r.memory.Lock()
	defer r.memory.Unlock()

	for _, ah := range r.memory.applicationsHistory {
		if ah.Email() == email {
			return ah, nil
		}
	}

	ah := applicationshistory.New(email)

	return ah, nil
}

func (r *InMemoryApplicationHistoryRepository) Save(
	ctx context.Context,
	applicationHistory *applicationshistory.ApplicationHistory,
) error {
	r.memory.Lock()
	defer r.memory.Unlock()

	for i, ah := range r.memory.applicationsHistory {
		if ah.Email() == applicationHistory.Email() {
			r.memory.applicationsHistory[i] = applicationHistory
			return nil
		}
	}

	r.memory.applicationsHistory = append(r.memory.applicationsHistory, applicationHistory)

	return nil
}
