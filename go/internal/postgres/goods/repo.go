package goods

type Repository interface {
	GetByOKPD2(okpd2 int64) (Goods, error)
	Update(inn int64, new *Goods) error
}
