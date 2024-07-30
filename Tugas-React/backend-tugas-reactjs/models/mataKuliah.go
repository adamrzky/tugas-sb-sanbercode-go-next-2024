// MataKuliah.go
package models

type MataKuliah struct {
	ID   uint `gorm:"primarykey"`
	Nama string
}

// TableName overrides the table name used by User to `mata_kuliah`
func (MataKuliah) TableName() string {
	return "mata_kuliah"
}
