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

func main() {
	companyRepository := mock.NewCompanyRepository(
		&domain.Company{ID: 1, Name: "Acme"},
		&domain.Company{ID: 2, Name: "Rosnano"},
		&domain.Company{ID: 3, Name: "Hamashi"},
	)
	productRepository := mock.NewProductRepository(
		&domain.Product{ID: 1, Name: "Car", CompanyID: 1},
		&domain.Product{ID: 2, Name: "Truck", CompanyID: 1},
		&domain.Product{ID: 3, Name: "Rocket", CompanyID: 2},
	)

	start := time.Now()
	defer func() {
		log.Println("time elapsed:", time.Since(start))
	}()

	ctx, waiter := async.NewCancelingWaiter(context.Background())
	findCompanies := async.Go(ctx, "acme", companyRepository.Find)
	findProducts := async.Go(ctx, "car", productRepository.Find)
	doError := async.Go(ctx, 0, func(ctx context.Context, in int) (int, error) {
		<-time.After(70 * time.Millisecond)
		return 0, fmt.Errorf("await error")
	})

	err := waiter.Await(findCompanies, findProducts, doError)
	if err != nil {
		log.Println("await error:", err)

		time.Sleep(5 * time.Millisecond)
		return
	}

	companies, _ := findCompanies.Get()
	products, _ := findProducts.Get()

	for i, company := range companies {
		log.Println("company", i, "is", company.Name)
	}
	for i, product := range products {
		log.Println("product", i, "is", product.Name)
	}
}
