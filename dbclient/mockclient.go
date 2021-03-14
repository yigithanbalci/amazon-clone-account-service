package dbclient

import (
	"github.com/stretchr/testify/mock"
	"github.com/yigithanbalci/amazon-clone-account-service/model"
)

// mocking using stretchr/testify
type MockBoltClient struct {
	mock.Mock
}

func (m *MockBoltClient) QueryAccount(accountID string) (model.Account, error) {
	args := m.Mock.Called(accountID)
	return args.Get(0).(model.Account), args.Error(1)
}

func (m *MockBoltClient) OpenBoltDb() {
	// does nothing
}

func (m *MockBoltClient) Seed() {
	// does nothing
}
