package gateway

import (
	"context"
	"errors"
	"sync"

	"github.com/morning-night-dream/platform-app/internal/domain/model"
	"github.com/morning-night-dream/platform-app/internal/domain/repository"
)

var _ repository.APICode = (*APICode)(nil)

type APICode struct {
	lock  sync.Mutex
	Codes map[model.CodeID]model.Code
}

func NewAPICode() *APICode {
	return &APICode{
		Codes: make(map[model.CodeID]model.Code),
	}
}

func (ac *APICode) Save(ctx context.Context, Code model.Code) error {
	ac.lock.Lock()
	defer ac.lock.Unlock()

	ac.Codes[Code.CodeID] = Code

	return nil
}

func (ac *APICode) Find(ctx context.Context, CodeID model.CodeID) (model.Code, error) {
	ac.lock.Lock()
	defer ac.lock.Unlock()

	if val, ok := ac.Codes[CodeID]; ok {
		return val, nil
	}

	return model.Code{}, errors.New("code not found")
}

func (ac *APICode) Delete(ctx context.Context, CodeID model.CodeID) error {
	ac.lock.Lock()
	defer ac.lock.Unlock()

	delete(ac.Codes, CodeID)

	return nil
}
