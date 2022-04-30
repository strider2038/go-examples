package domain

import "context"

type Product struct {
	ID        int
	Name      string
	CompanyID int
}

type ProductRepository interface {
	Find(ctx context.Context, search string) ([]*Product, error)
}
