package tax

import "errors"

type Repository interface {
	SaveTax(amount float64) error
}

func CalculateTaxAndSave(amount float64, repository Repository) error {
	result := CalculateTax2(amount)
	return repository.SaveTax(result)
}

func CalculateTax2(amount float64) float64 {
	if amount <= 0.0 {
		return 0.0
	}
	if amount >= 10000 && amount < 20000 {
		return 10.0
	}
	return 5.0
}

func CalculateTax(amount float64) (float64, error) {
	if amount <= 0.0 {
		return 0.0, errors.New("amount must be greater than 0")
	}
	if amount >= 10000 && amount < 20000 {
		return 10.0, nil
	}
	return 5.0, nil
}
