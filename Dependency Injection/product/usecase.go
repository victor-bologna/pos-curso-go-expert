package product

type ProductUseCase struct {
	ProductRepository ProductRepositoryInterface
}

func NewProductUseCase(ProductRepository ProductRepositoryInterface) *ProductUseCase {
	return &ProductUseCase{ProductRepository: ProductRepository}
}

func (puc *ProductUseCase) GetProductId(id int) (Product, error) {
	return puc.ProductRepository.GetProductId(id)
}
