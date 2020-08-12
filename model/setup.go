package model

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres
)

// DB - gorm DB
var DB *gorm.DB

// SetupDB is database setup
func SetupDB() {
	DSN := os.Getenv("DSN")

	var err error
	DB, err = gorm.Open("postgres", DSN)

	if err != nil {
		log.Fatal(err)
	}

	// Query log
	DB.LogMode(true)

	fmt.Println("connected to database ...")

	DB.SingularTable(true)

	// adding default prefix to table name
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "dp_" + defaultTableName
	}

	// db migrations
	DB.AutoMigrate(
		&Currency{},
		&Payment{},
		&Membership{},
		&Plan{},
		&User{},
		&Product{},
		&Tag{},
		&Catalog{},
		&Cart{},
		&Order{},
		&OrderItem{},
		&Dataset{},
		&Format{},
		&DatasetFormat{},
		&Medium{},
	)

	// Adding foreignKey
	DB.Model(&Payment{}).AddForeignKey("currency_id", "dp_currency(id)", "RESTRICT", "RESTRICT")
	DB.Model(&Membership{}).AddForeignKey("user_id", "dp_user(id)", "RESTRICT", "RESTRICT")
	DB.Model(&Membership{}).AddForeignKey("plan_id", "dp_plan(id)", "RESTRICT", "RESTRICT")
	DB.Model(&Membership{}).AddForeignKey("payment_id", "dp_payment(id)", "RESTRICT", "RESTRICT")
	DB.Model(&Product{}).AddForeignKey("currency_id", "dp_currency(id)", "RESTRICT", "RESTRICT")
	DB.Model(&Product{}).AddForeignKey("featured_medium_id", "dp_medium(id)", "RESTRICT", "RESTRICT")
	DB.Model(&Cart{}).AddForeignKey("user_id", "dp_user(id)", "RESTRICT", "RESTRICT")
	DB.Table("dp_cart_item").AddForeignKey("cart_id", "dp_cart(id)", "RESTRICT", "RESTRICT")
	DB.Table("dp_cart_item").AddForeignKey("product_id", "dp_product(id)", "RESTRICT", "RESTRICT")
	DB.Model(&Order{}).AddForeignKey("payment_id", "dp_payment(id)", "RESTRICT", "RESTRICT")
	DB.Model(&Order{}).AddForeignKey("cart_id", "dp_cart(id)", "RESTRICT", "RESTRICT")
	DB.Model(&OrderItem{}).AddForeignKey("product_id", "dp_product(id)", "RESTRICT", "RESTRICT")
	DB.Model(&DatasetFormat{}).AddForeignKey("format_id", "dp_format(id)", "RESTRICT", "RESTRICT")
	DB.Model(&Catalog{}).AddForeignKey("featured_medium_id", "dp_medium(id)", "RESTRICT", "RESTRICT")
	DB.Model(&Catalog{}).AddForeignKey("currency_id", "dp_currency(id)", "RESTRICT", "RESTRICT")
	DB.Model(&Dataset{}).AddForeignKey("featured_medium_id", "dp_medium(id)", "RESTRICT", "RESTRICT")
	DB.Model(&Dataset{}).AddForeignKey("currency_id", "dp_currency(id)", "RESTRICT", "RESTRICT")
}
