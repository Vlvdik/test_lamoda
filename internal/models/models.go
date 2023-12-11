package models

type Store struct {
	ID        int    `db:"id"`
	Name      string `db:"name"`
	Available bool   `db:"available"`
}

type Product struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`
	Size       string `db:"size"`
	UniqueCode string `db:"unique_code"`
	Count      int    `db:"count"`
}

type StoreProduct struct {
	ID        int `db:"id"`
	StoreID   int `db:"store_id"`
	ProductID int `db:"product_id"`
	Count     int `db:"count"`
}

type Reservation struct {
	ID        int `db:"id"`
	ProductID int `db:"product_id"`
	Count     int `db:"count"`
}
