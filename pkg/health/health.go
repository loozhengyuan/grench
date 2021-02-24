// Package health provides the health check interface and service.
package health

import (
	"strconv"
	"sync"
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
	Component string `json:"component,omitempty"`

	// Health status of the component.
	Status Status `json:"status,omitempty"`

	// Information on the health status.
	Message string `json:"message,omitempty"`

	// Timestamp when component was checked.
	CheckedAt time.Time `json:"checked_at,omitempty"`

	// Metadata information about the component check.
	Metadata map[string]string `json:"metadata,omitempty"`
}

// CheckFunc is a functional expression of a check procedure.
type CheckFunc func() Info

// Checker is the interface of a health check service.
type Checker interface {
	// Check executes and returns the outcome of all registered CheckFuncs.
	Check() []Info
}

type service struct {
	checks []CheckFunc
}

var _ Checker = (*service)(nil)

func (s service) Check() []Info {
	status := make([]Info, len(s.checks))
	var wg sync.WaitGroup
	for i, check := range s.checks {
		i, check := i, check // capture range variable
		wg.Add(1)
		go func() {
			defer wg.Done()
			status[i] = check()
		}()
	}
	wg.Wait()
	return status
}

// New returns a new health check service.
func New(checks ...CheckFunc) Checker {
	return &service{
		checks: checks,
	}
}
