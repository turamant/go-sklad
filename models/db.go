package models

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	*sql.DB
}

func NewDB(dsn string) (*DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Получение имени базы данных из командной строки
	dbName := flag.String("db-name", "sklad2", "Name of the database")
	flag.Parse()

	// Создание базы данных
	if err := createDatabase(db, *dbName); err != nil {
		return nil, err
	}

	// Выбор базы данных
	if _, err := db.Exec(fmt.Sprintf("USE %s", *dbName)); err != nil {
		return nil, fmt.Errorf("failed to select database: %v", err)
	}

	// Создание таблиц
	if err := createTables(db); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func createDatabase(db *sql.DB, dbName string) error {
	// Создание базы данных
	_, err := db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName))
	if err != nil {
		return fmt.Errorf("failed to create database: %v", err)
	}
	fmt.Printf("Database '%s' created\n", dbName)

	return nil
}

func createTables(db *sql.DB) error {
	// Создание таблицы 'products'
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS products (
			id INT AUTO_INCREMENT PRIMARY KEY,
			article_number VARCHAR(255) UNIQUE NOT NULL,
			name VARCHAR(255) NOT NULL,
			description TEXT NOT NULL,
			purchase_price DECIMAL(10,2) NOT NULL,
			sell_price DECIMAL(10,2) NOT NULL,
			quantity INT NOT NULL,
			user_id INT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
		)
	`)

	if err != nil {
		return fmt.Errorf("failed to create table 'products': %v", err)
	}
	fmt.Println("Table 'products' created")

	// Создание таблицы 'suppliers'
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS suppliers (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			contact_info VARCHAR(255) NOT NULL,
			user_id INT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create table 'suppliers': %v", err)
	}
	fmt.Println("Table 'suppliers' created")

	// Создание таблицы 'arrivals'
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS arrivals (
			id INT AUTO_INCREMENT PRIMARY KEY,
			product_id INT NOT NULL,
			supplier_id INT NOT NULL,
			quantity INT NOT NULL,
			purchase_price DECIMAL(10,2) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create table 'arrivals': %v", err)
	}
	fmt.Println("Table 'arrivals' created")

	// Создание таблицы 'expenses'
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS expenses (
			id INT AUTO_INCREMENT PRIMARY KEY,
			product_id INT NOT NULL,
			customer_id INT NOT NULL,
			quantity INT NOT NULL,
			sell_price DECIMAL(10,2) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create table 'expenses': %v", err)
	}
	fmt.Println("Table 'expenses' created")

	// Создание таблицы 'customers'
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS customers (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			contact_info VARCHAR(255) NOT NULL,
			user_id INT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create table 'customers': %v", err)
	}
	fmt.Println("Table 'customers' created")

	// Создание таблицы 'discounts'
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS discounts (
		id INT AUTO_INCREMENT PRIMARY KEY,
		product_id INT NOT NULL,
		discount_percentage DECIMAL(5,2) NOT NULL,
		start_date DATE NOT NULL,
		end_date DATE NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	)
	`)
	if err != nil {
		return fmt.Errorf("failed to create table 'discounts': %v", err)
	}
	fmt.Println("Table 'discounts' created")

	return nil
}
