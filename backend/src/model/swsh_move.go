package model

type SwShMove struct {
	Id             int    `gorm:"primaryKey"`
	Name           string `gorm:"column:name;not null"`
	TypeId         int    `gorm:"column:type_id;not null"`
	MoveCategoryId int    `gorm:"column:move_category_id;not null"`
	Power          *int   `gorm:"column:power;"`
	Accuracy       *int   `gorm:"column:accuracy;"`
}

type SwShMoveRelation struct {
	SwShMove
	Type         Type         `gorm:"foreignKey:TypeId;referneces:Id"`
	MoveCategory MoveCategory `gorm:"foreignKey:MoveCategoryId;references:Id"`
}

func (m SwShMoveRelation) TableName() string {
	return "swsh_moves"
}
