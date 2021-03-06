package models

type Todo struct {
	IDModel
	ActivityGroupId		uint		`gorm:"not null" json:"activity_group_id"`
	Title				string		`gorm:"not null;type:varchar(255)" json:"title"`
	IsActive			bool		`gorm:"not null" json:"is_active"`
	Priority			string		`gorm:"not null;type:varchar(16)" json:"priority"`
	Activity			Activity	`gorm:"foreignKey:ActivityGroupId" json:"-"`
	TimestampModel
}

func (t *Todo) TableName() string {
	return "todos"
}