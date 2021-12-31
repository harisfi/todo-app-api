package models

type Activity struct {
    IDModel
    Email       string      `gorm:"not null;type:varchar(64)" json:"email"`
    Title       string      `gorm:"not null;type:varchar(16)" json:"title"`
    TimestampModel
}

func (a *Activity) TableName() string {
    return "activities"
}