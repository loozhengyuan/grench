// Package health provides the health check interface and service.
package health

import (
	"strconv"
	"time"
)

// Status represents the status of a health check.
type Status int

// Enum types for the status of a health check.
const (
	StatusUnspecified Status = iota // Default null value
	StatusPass                      // Component is healthy
	StatusFail                      // Component is unhealthy
)

// String returns the string value of a Status
func (s Status) String() string {
	switch s {
	case StatusUnspecified:
		return "unspecified"
	case StatusPass:
		return "pass"
	case StatusFail:
		return "fail"
	default:
		return strconv.Itoa(int(s))
	}
}

// Info represents the outcome of a health check.
type Info struct {
	// Name of the component, e.g. database, network, etc.
	Component string

	// Health status of the component.
	Status Status

	// Information on the health status.
	Message string

	// Timestamp when component was checked.
	CheckedAt time.Time

	// Metadata information about the component check.
	Metadata map[string]string
}

// CheckFunc is a functional expression of a check procedure.
type CheckFunc func() (Info, error)

// Checker is the interface of a health check service.
type Checker interface {
	// Register registers one or more CheckFuncs with the health check service.
	Register(...CheckFunc)

	// Check executes and returns the outcome of all registered CheckFuncs.
	Check() ([]Info, error)
}

type service struct {
	checks []CheckFunc
}

var _ Checker = (*service)(nil)

func (s *service) Register(checks ...CheckFunc) {
	for _, c := range checks {
		s.checks = append(s.checks, c)
	}
}

func (s service) Check() ([]Info, error) {
	var output []Info
	for _, check := range s.checks {
		info, err := check()
		if err != nil {
			return nil, err
		}
		output = append(output, info)
	}
	return output, nil
}

// New returns a new health check service.
func New() Checker {
	return &service{}
}
