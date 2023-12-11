package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	COUNT_OF_PRODUCT = `SELECT count FROM product WHERE unique_code = $1`

	PRODUCT_RESERVE = `
		INSERT INTO reservation (product_id, count)
		VALUES ((SELECT id FROM product WHERE unique_code = $1), 1)
		ON CONFLICT (product_id) DO UPDATE SET count = reservation.count + 1;	
		UPDATE product SET count = count - 1 WHERE unique_code = $1;
	`

	RELEASE_RESERVATION = `
		UPDATE reservation SET count = 0 WHERE product_id = (SELECT id FROM product WHERE unique_code = $1);
		UPDATE product SET count = count + 1 WHERE unique_code = $1;
	`

	GET_REMAINING_PRODUCTS = `
		SELECT SUM(count)
		FROM store_product
		WHERE store_id = $1;
	`
)

func ReserveProduct(db *sql.DB, uniqueCodes []string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
		}
	}()

	for _, code := range uniqueCodes {
		var count int
		err := tx.QueryRow(COUNT_OF_PRODUCT, code).Scan(&count)
		if err != nil {
			tx.Rollback()
			return err
		}
		if count <= 0 {
			tx.Rollback()
			return fmt.Errorf("insufficient quantity. Code: %s", code)
		}

		_, err = tx.Exec(PRODUCT_RESERVE, code)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func ReleaseReservation(db *sql.DB, uniqueCodes []string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
		}
	}()

	for _, code := range uniqueCodes {
		_, err := tx.Exec(RELEASE_RESERVATION, code)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func GetRemainingProducts(db *sql.DB, storeID int32) (int, error) {
	var remainingCount int

	err := db.QueryRow(GET_REMAINING_PRODUCTS, storeID).Scan(&remainingCount)

	if err != nil {
		return 0, err
	}

	return remainingCount, nil
}
