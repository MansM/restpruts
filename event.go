package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Event type
type Event struct {
	gorm.Model
	id   int    `type:int;primary_key;default:unique_rowid()"`
	Name string `json:"name"`
}

var events []Event
