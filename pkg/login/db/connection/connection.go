package connection

import {
	"gorm.io/gorm"
	"gorm.io/driver/postgres"

}

dsn := "host=localhost user=buddi password=admin dbname=avua-lo port=5432 sslmode=disable TimeZone=Asia/Shanghai"
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})