package mock

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/strider2038/go-examples/concurrency/async-await/domain"
)

type ProductRepository struct {
	products map[int]*domain.Product
}

func NewProductRepository(products ...*domain.Product) *ProductRepository {
	repository := &ProductRepository{products: map[int]*domain.Product{}}

	for _, product := range products {
		repository.products[product.ID] = product
	}

	return repository
}

func (repository *ProductRepository) Find(ctx context.Context, search string) ([]*domain.Product, error) {
	select {
	case <-ctx.Done():
		log.Println("product search canceled")

		return nil, ctx.Err()
	case <-time.After(100 * time.Millisecond):
	}

	products := make([]*domain.Product, 0)

	for _, product := range repository.products {
		if strings.Contains(strings.ToLower(product.Name), strings.ToLower(search)) {
			products = append(products, product)
		}
	}

	log.Println("product search done")

	return products, nil
}
