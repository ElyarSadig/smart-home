package iot

import (
	"context"
	"database/sql"
	"math/rand"
	"time"

	"log"
)

const MAX_TEMPERATURE = 45
const MIN_TEMPERATURE = 19

func AirConditionerWorker(ctx context.Context, db *sql.DB) {
	ch := make(chan bool)
	go updateTemperature(ctx, db, ch)
	go getAirConditionerActivityStatus(ctx, db, ch)
}

func updateTemperature(ctx context.Context, db *sql.DB, ch <-chan bool) {
	for {
		isActive := <-ch
		temperature := getRoomTemperature(ctx, db)
		if isActive && temperature > MIN_TEMPERATURE {
			temperature -= rand.Float64()
		} else if !isActive && temperature < MAX_TEMPERATURE {
			temperature += rand.Float64()
		}
		updateRoomTemperature(ctx, db, temperature)
	}
}

func getRoomTemperature(ctx context.Context, db *sql.DB) float64 {
	query := `SELECT temperature from rooms WHERE id = 1`
	row := db.QueryRowContext(ctx, query)
	var temperature float64
	err := row.Scan(&temperature)
	if err != nil {
		log.Printf("Error in scanning: %v", err)
	}
	return temperature
}

func updateRoomTemperature(ctx context.Context, db *sql.DB, temperature float64) {
	query := `UPDATE rooms SET temperature = ? WHERE id = 1`
	_, err := db.ExecContext(ctx, query, temperature)
	if err != nil {
		log.Printf("Error in updating: %v", err)
	}
}

func getAirConditionerActivityStatus(ctx context.Context, db *sql.DB, ch chan<- bool) {
	for {
		time.Sleep(time.Second * 1)
		query := `SELECT is_active FROM devices WHERE id = 3`
		row := db.QueryRowContext(ctx, query)
		var isActive bool
		err := row.Scan(&isActive)
		if err != nil {
			log.Printf("Error in sanning: %v", err)
		}
		ch <- isActive
	}
}
