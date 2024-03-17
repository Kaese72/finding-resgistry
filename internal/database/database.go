package database

import "github.com/Kaese72/findings-registry/internal/intermediaries"

type Persistence interface {
	UpdateFinding(intermediaries.Finding) (intermediaries.Finding, error)
	GetFinding(string) (intermediaries.Finding, error)
	GetFindings() ([]intermediaries.Finding, error)
	Purge() error
}
