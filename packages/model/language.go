package model

import (
	"github.com/ug93tad/go-apla/packages/converter"
)

// Language is model
type Language struct {
	tableName  string
	ID         int64  `gorm:"primary_key;not null"`
	AppID      int64  `gorm:"column:app_id;not null"`
	Name       string `gorm:"not null;size:100"`
	Res        string `gorm:"type:jsonb(PostgreSQL)"`
	Conditions string `gorm:"not null"`
}

// SetTablePrefix is setting table prefix
func (l *Language) SetTablePrefix(tablePrefix string) {
	l.tableName = tablePrefix + "_languages"
}

// TableName returns name of table
func (l *Language) TableName() string {
	return l.tableName
}

// GetAll is retrieving all records from database
func (l *Language) GetAll(prefix string) ([]Language, error) {
	result := new([]Language)
	err := DBConn.Table(prefix + "_languages").Order("name").Find(&result).Error
	return *result, err
}

// ToMap is converting model to map
func (l *Language) ToMap() map[string]string {
	result := make(map[string]string, 0)
	result["name"] = l.Name
	result["res"] = l.Res
	result["conditions"] = l.Conditions
	result["app_id"] = converter.Int64ToStr(l.AppID)
	return result
}
