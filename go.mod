module homecloud

go 1.18

replace homecloud/core => ./core

replace homecloud/accounting => ./accounting

replace homecloud/dashboard => ./dashboard

require homecloud/core v0.0.0

require homecloud/accounting v0.0.0

require homecloud/dashboard v0.0.0

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/mattn/go-sqlite3 v1.14.11 // indirect
	gorm.io/driver/sqlite v1.2.6 // indirect
	gorm.io/gorm v1.22.5 // indirect
)
