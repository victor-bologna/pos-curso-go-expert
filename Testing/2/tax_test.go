package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	tax, err := CalculateTax(1000.00)
	assert.Nil(t, err)
	assert.Equal(t, 10.0, tax)

	_, err = CalculateTax(0)
	assert.Error(t, err)
	assert.Equal(t, 0.0, tax)
	assert.Contains(t, err.Error(), "greater than 0")
}

func TestCalculateTax2AndSave(t *testing.T) {
	repository := &TaxRepositoryMock{}
	repository.On("SaveTax", 10.0).Return(nil)
	repository.On("SaveTax", 0.0).Return(errors.New("Error when saving Tax."))

	err := CalculateTaxAndSave(10000.00, repository)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(0.0, repository)
	assert.Error(t, err, errors.New("Error when saving tax"))

	repository.AssertExpectations(t)
	repository.AssertNumberOfCalls(t, "SaveTax", 2)
}
