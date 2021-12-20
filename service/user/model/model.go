package model

type Model struct {
	ID uint `gorm:"primaryKey"`
	CreatedAt int64 `gorm:"autoUpdateTime;not null"`
	UpdatedAt int64 `gorm:"autoCreateTime;not null"`
}
