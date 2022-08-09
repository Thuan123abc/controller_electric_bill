package database

type Customer struct {
	IDCus   int64 `gorm:"primaryKey"`
	NameCus string
	Addr    string
	IDMeter int64 `gorm:"foreignKey:IDMeter"`
}
type ElectricBill struct {
	Customer
	IDBill    int64 `gorm:"primaryKey"`
	IDMeter   int64
	OldNumber int64
	NewNumber int64
}

func (c Customer) TableName() string {
	return "customer"
}
func (b ElectricBill) TableName() string {
	return "bill"
}
