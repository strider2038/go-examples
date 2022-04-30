package domain

import "context"

type Company struct {
	ID   int
	Name string
}

type CompanyRepository interface {
	Find(ctx context.Context, search string) ([]*Company, error)
}
