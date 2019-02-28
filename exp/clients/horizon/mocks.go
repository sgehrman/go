package horizonclient

import (
	"github.com/stretchr/testify/mock"
)

// MockClient is a mockable horizon client.
type MockClient struct {
	mock.Mock
}

// AccountDetail is a mocking method
func (m *MockClient) AccountDetail(request AccountRequest) (Account, error) {
	a := m.Called(request)
	return a.Get(0).(Account), a.Error(1)
}

// AccountData is a mocking method
func (m *MockClient) AccountData(request AccountRequest) (AccountData, error) {
	a := m.Called(request)
	return a.Get(0).(AccountData), a.Error(1)
}

// AllEffects is a mocking method
func (m *MockClient) AllEffects(request EffectRequest) (EffectsPage, error) {
	a := m.Called(request)
	return a.Get(0).(EffectsPage), a.Error(1)
}

// LedgerEffects is a mocking method
func (m *MockClient) LedgerEffects(request EffectRequest) (EffectsPage, error) {
	a := m.Called(request)
	return a.Get(0).(EffectsPage), a.Error(1)
}

// OperationEffects is a mocking method
func (m *MockClient) OperationEffects(request EffectRequest) (EffectsPage, error) {
	a := m.Called(request)
	return a.Get(0).(EffectsPage), a.Error(1)
}

// TransactionEffects is a mocking method
func (m *MockClient) TransactionEffects(request EffectRequest) (EffectsPage, error) {
	a := m.Called(request)
	return a.Get(0).(EffectsPage), a.Error(1)
}

// ensure that the MockClient implements ClientInterface
var _ ClientInterface = &MockClient{}
