package suppliers

type Repository interface {
	GetByINN(inn int64) (Supplier, error)
	Update(inn int64, new *Supplier) error
}
