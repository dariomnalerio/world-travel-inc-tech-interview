package testing

import "github.com/stretchr/testify/mock"

type MockDogBuilder struct {
	mock *MockDogRepository
}

func NewDogMockBuilder() *MockDogBuilder {
	return &MockDogBuilder{
		mock: &MockDogRepository{},
	}
}

// WithRandomPicture sets up the mock to return a random dog picture.
func (b *MockDogBuilder) WithSuccessfulRandomPicture(url string) *MockDogBuilder {
	b.mock.On("GetRandomPicture").Return(url, nil).Once()
	return b
}

func (b *MockDogBuilder) WithFailedRandomPicture(err error) *MockDogBuilder {
	b.mock.On("GetRandomPicture").Return("", err).Once()
	return b
}

func (b *MockDogBuilder) Build() *MockDogRepository {
	return b.mock
}

func (b *MockDogBuilder) AssertExpectations(t mock.TestingT) {
	b.mock.AssertExpectations(t)
}
