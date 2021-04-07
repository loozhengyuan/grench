package test

import (
	"testing"
)

// Runner represents a Runner interface.
type Runner interface {
	// Run executes the test suite.
	Run(t *testing.T)
}

// Suite represents a test suite.
type Suite struct {
	SetUpSuiteFunc    func(t *testing.T)
	SetUpTestFunc     func(t *testing.T)
	RunFunc           func(t *testing.T)
	TearDownSuiteFunc func(t *testing.T)
	TearDownTestFunc  func(t *testing.T)
}

var _ Runner = (*Suite)(nil)

// Run executes the test suite.
func (s *Suite) Run(t *testing.T) {
	s.SetUpSuiteFunc(t)
	s.SetUpTestFunc(t)
	s.RunFunc(t)
	s.TearDownTestFunc(t)
	s.TearDownSuiteFunc(t)
}
