package mock

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/strider2038/go-examples/concurrency/async-await/domain"
)

type CompanyRepository struct {
	companies map[int]*domain.Company
}

func NewCompanyRepository(companies ...*domain.Company) *CompanyRepository {
	repository := &CompanyRepository{companies: map[int]*domain.Company{}}

	for _, company := range companies {
		repository.companies[company.ID] = company
	}

	return repository
}

func (repository *CompanyRepository) Find(ctx context.Context, search string) ([]*domain.Company, error) {
	select {
	case <-ctx.Done():
		log.Println("company search canceled")

		return nil, ctx.Err()
	case <-time.After(50 * time.Millisecond):
	}

	companies := make([]*domain.Company, 0)

	for _, company := range repository.companies {
		if strings.Contains(strings.ToLower(company.Name), strings.ToLower(search)) {
			companies = append(companies, company)
		}
	}

	log.Println("company search done")

	return companies, nil
}
