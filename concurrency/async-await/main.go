package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/strider2038/go-examples/concurrency/async-await/async"
	"github.com/strider2038/go-examples/concurrency/async-await/domain"
	"github.com/strider2038/go-examples/concurrency/async-await/mock"
)

type UseCase struct {
	companies domain.CompanyRepository
	products  domain.ProductRepository
}

type Response struct {
	Companies []*domain.Company
	Products  []*domain.Product
}

func NewUseCase() *UseCase {
	return &UseCase{
		companies: mock.NewCompanyRepository(
			&domain.Company{ID: 1, Name: "Acme"},
			&domain.Company{ID: 2, Name: "Rosnano"},
			&domain.Company{ID: 3, Name: "Hamashi"},
		),
		products: mock.NewProductRepository(
			&domain.Product{ID: 1, Name: "Car", CompanyID: 1},
			&domain.Product{ID: 2, Name: "Truck", CompanyID: 1},
			&domain.Product{ID: 3, Name: "Rocket", CompanyID: 2},
		),
	}
}

func (useCase *UseCase) Handle(ctx context.Context) (*Response, error) {
	ctx, waiter := async.NewCancelingWaiter(context.Background())
	companies := async.Go(ctx, useCase.companies.Find, "acme")
	products := async.Go(ctx, useCase.products.Find, "car")
	errorFunc := func(ctx context.Context, in int) (int, error) {
		<-time.After(70 * time.Millisecond)
		return 0, fmt.Errorf("await error")
	}
	doError := async.Go(ctx, errorFunc, 0)

	err := waiter.Await(companies, products, doError)
	if err != nil {
		return nil, fmt.Errorf("await error: %w", err)
	}

	response := &Response{}
	response.Companies, _ = companies.Get()
	response.Products, _ = products.Get()

	return response, nil
}

func main() {
	useCase := NewUseCase()

	start := time.Now()
	defer func() {
		log.Println("time elapsed:", time.Since(start))
	}()

	response, err := useCase.Handle(context.Background())
	if err != nil {
		log.Println(err)
		time.Sleep(5 * time.Millisecond)

		return
	}

	for i, company := range response.Companies {
		log.Println("company", i, "is", company.Name)
	}
	for i, product := range response.Products {
		log.Println("product", i, "is", product.Name)
	}
}
