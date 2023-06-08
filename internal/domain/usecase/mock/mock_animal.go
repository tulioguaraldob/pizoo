// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/usecase/usecase.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	entity "github.com/TulioGuaraldoB/pizoo/internal/domain/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockIAnimalUsecase is a mock of IAnimalUsecase interface.
type MockIAnimalUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockIAnimalUsecaseMockRecorder
}

// MockIAnimalUsecaseMockRecorder is the mock recorder for MockIAnimalUsecase.
type MockIAnimalUsecaseMockRecorder struct {
	mock *MockIAnimalUsecase
}

// NewMockIAnimalUsecase creates a new mock instance.
func NewMockIAnimalUsecase(ctrl *gomock.Controller) *MockIAnimalUsecase {
	mock := &MockIAnimalUsecase{ctrl: ctrl}
	mock.recorder = &MockIAnimalUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAnimalUsecase) EXPECT() *MockIAnimalUsecaseMockRecorder {
	return m.recorder
}

// CreateAnimal mocks base method.
func (m *MockIAnimalUsecase) CreateAnimal(animal *entity.Animal) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAnimal", animal)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAnimal indicates an expected call of CreateAnimal.
func (mr *MockIAnimalUsecaseMockRecorder) CreateAnimal(animal interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAnimal", reflect.TypeOf((*MockIAnimalUsecase)(nil).CreateAnimal), animal)
}

// DeleteAnimal mocks base method.
func (m *MockIAnimalUsecase) DeleteAnimal(animalId uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAnimal", animalId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAnimal indicates an expected call of DeleteAnimal.
func (mr *MockIAnimalUsecaseMockRecorder) DeleteAnimal(animalId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAnimal", reflect.TypeOf((*MockIAnimalUsecase)(nil).DeleteAnimal), animalId)
}

// GetAllAnimals mocks base method.
func (m *MockIAnimalUsecase) GetAllAnimals() ([]entity.Animal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllAnimals")
	ret0, _ := ret[0].([]entity.Animal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllAnimals indicates an expected call of GetAllAnimals.
func (mr *MockIAnimalUsecaseMockRecorder) GetAllAnimals() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllAnimals", reflect.TypeOf((*MockIAnimalUsecase)(nil).GetAllAnimals))
}

// GetAnimalById mocks base method.
func (m *MockIAnimalUsecase) GetAnimalById(animalId uint) (*entity.Animal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAnimalById", animalId)
	ret0, _ := ret[0].(*entity.Animal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAnimalById indicates an expected call of GetAnimalById.
func (mr *MockIAnimalUsecaseMockRecorder) GetAnimalById(animalId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAnimalById", reflect.TypeOf((*MockIAnimalUsecase)(nil).GetAnimalById), animalId)
}