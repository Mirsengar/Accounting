package db

import (
	"Accounting/helpers"
	"Accounting/models"
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	sql *gorm.DB
}

var logger *log.Logger = helpers.GetLoggerInstace()
func CreateConnection(dbName string) Database {
	db, err := gorm.Open(sqlite.Open(fmt.Sprintf("%s.db", dbName)), &gorm.Config{})
	if err != nil {
		logger.Println("Connection is not setup with ecommerce.db")
		logger.Fatalln(err)
	}
	db.AutoMigrate(&models.Product{}, &models.Category{}, &models.Cart{}, &models.Invoice{}, &models.User{})
	return Database{db}
}

func (db *Database) InsertRow(table string, value interface{}) error {

	var result *gorm.DB
	switch value.(type) {

	case *models.Category:
		result = db.sql.Create(value.(*models.Category))
	case *models.Product:
		result = db.sql.Create(value.(*models.Product))
	case *models.User:
		result = db.sql.Create(value.(*models.User))
	case *models.Cart:
		result = db.sql.Create(value.(*models.Cart))
	case *models.Invoice:
		result = db.sql.Create(value.(*models.Invoice))
	}
	if result.Error != nil {
		return result.Error
	}
	logger.Println("Successfully added entry to", table)
	return nil
}

func (db *Database) Delete(model interface{}, ids []uint) error {

	res := db.sql.Where("id IN ?", ids).Delete(model)
	if res.Error != nil {
		return res.Error
	}
	logger.Println("Affected rows:", res.RowsAffected)
	return nil
}

func (db *Database) DeleteFromCart(ids []uint) error {
	res := db.sql.Where("product_id IN ?", ids).Delete(&models.Cart{})
	if res.Error != nil {
		return res.Error
	}
	logger.Println("Deleted products with ids", ids)
	return nil
}

func (db *Database) SaveBill(bill *models.Invoice) error {
	res := db.sql.Create(bill)
	if res.Error != nil {
		return res.Error
	}
	logger.Println("Saved bill")
	return nil
}

func (db *Database) UpdateCart(cart *models.Cart) error {
	res := db.sql.Create(cart)
	if res.Error != nil {
		return res.Error
	}
	logger.Println("Updated cart products")
	return nil
}

func (db *Database) DeleteFromUserCart(userID uint) error {
	res := db.sql.Where("user_id=?", userID).Delete(&models.Cart{})
	if res.Error != nil {
		return res.Error
	}
	logger.Println("Removed User's cart")
	return nil
}

func (db *Database) GetUserDetails(id uint, user *models.User) {
	res := db.sql.Where("id=?", id).Find(&user)
	if res.Error != nil {
		logger.Fatalln("Failed to fetch user details for id:", id, "\n", res.Error)
	}
	logger.Println("Fetched user details for id:", id)
}

func (db *Database) GetBills(bills *[]models.Invoice) {
	res := db.sql.Find(bills)
	if res.Error != nil {
		logger.Fatalln("Failed to get bills \n", res.Error)
	}
	logger.Println("Fetched bills")
}

func (db *Database) FetchCartProducts(userID uint, products *[]models.Product) {
	cartProducts := []models.Cart{}
	res := db.sql.Where("user_id=?", userID).Find(&cartProducts)
	if res.Error != nil {
		logger.Fatalln("Failed to get cart products \n", res.Error)
	}
	prods := []models.Product{}
	for _, cp := range cartProducts {
		p := models.Product{}
		db.FetchProductDetails(cp.ProductID, &p)
		prods = append(prods, p)
	}
	*products = prods
	logger.Println("Fetched cart products")
}

func (db *Database) FetchProducts(pids []uint, products *[]models.Product) {
	res := db.sql.Where("id IN ?", pids).Find(&products)
	if res.Error != nil {
		logger.Fatalln("Failed to get products \n", res.Error)
	}
	logger.Println("Fetched products")
}

func (db *Database) FetchAllCategories(categories *[]models.Category) {
	res := db.sql.Find(&categories)
	if res.Error != nil {
		logger.Fatalln("Failed to fetch categories:\n", res.Error)
	}
	logger.Println("Fetched categories from `categories` table")
}

func (db *Database) FetchProductsOfCategory(categoryName string, products *[]models.Product) {
	res := db.sql.Find(&products, map[string]interface{}{"category_name": categoryName})
	if res.Error != nil {
		logger.Fatalln("Failed to fetch products for categoryName: ", categoryName, "\n", res.Error)
	}
	logger.Println("Fetched products from category", categoryName)
}

func (db *Database) FetchProductDetails(productID uint, product *models.Product) {
	res := db.sql.First(&product, map[string]interface{}{"id": productID})
	if res.Error != nil {
		logger.Fatalln("Failed to fetch product details for productID: ", productID, "\n", res.Error)
	}
	logger.Println("Product Details fetched for id:", productID)
}

func (db *Database) FetchCartDetails(userID uint, cart *models.Cart) {
	res := db.sql.Where("user_id=?", userID).Find(&cart)
	if res.Error != nil {
		logger.Fatalln("Failed to cart details for a user \n", res.Error)
	}
	logger.Println("Fetched cart details for user:", userID)
}
func (db *Database) FetchParticularProducts(userID uint, pids []uint, cartProducts *[]models.Product) {

	cps := []models.Cart{}
	res := db.sql.Where("user_id=? AND product_id IN ?", userID, pids).Find(&cps)
	if res.Error != nil {
		logger.Fatalln("Failed to get cart products for productIDS", pids, "\n", res.Error)
	}
	prods := []models.Product{}
	for _, cp := range cps {
		p := models.Product{}
		db.FetchProductDetails(cp.ProductID, &p)
		prods = append(prods, p)
	}
	*cartProducts = prods
	logger.Println("Fetched products with ids", pids)
}