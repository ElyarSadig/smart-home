package iot

import (
	"context"
	"database/sql"
	"log"
	"math/rand"
	"time"
)

const MAX_HUMIDITY = 50
const MIN_HUMIDITY = 20

func AirHumidifeirWorker(ctx context.Context, db *sql.DB) {
	ch := make(chan bool)
	go getAirHumidifierActivityStatus(ctx, db, ch)
	go updateHumidity(ctx, db, ch)
}

func updateHumidity(ctx context.Context, db *sql.DB, ch <-chan bool) {
	for {
		isActive := <-ch
		humidity := getRoomHumidity(ctx, db)
		if isActive && humidity < MAX_HUMIDITY {
			humidity += rand.Float64()
		} else if !isActive && humidity > MIN_HUMIDITY {
			humidity -= rand.Float64()
		}
		updateRoomHumidity(ctx, db, humidity)
	}
}

func getRoomHumidity(ctx context.Context, db *sql.DB) float64 {
	query := `SELECT humidity from rooms WHERE id = 1`
	row := db.QueryRowContext(ctx, query)
	var humidity float64
	err := row.Scan(&humidity)
	if err != nil {
		log.Printf("Error in scanning: %v", err)
	}
	return humidity
}

func updateRoomHumidity(ctx context.Context, db *sql.DB, humidity float64) {
	query := `UPDATE rooms SET humidity = ? WHERE id = 1`
	_, err := db.ExecContext(ctx, query, humidity)
	if err != nil {
		log.Printf("Error in updating: %v", err)
	}
}

func getAirHumidifierActivityStatus(ctx context.Context, db *sql.DB, ch chan<- bool) {
	for {
		time.Sleep(time.Second * 1)
		query := `SELECT is_active FROM devices WHERE id = 5`
		row := db.QueryRowContext(ctx, query)
		var isActive bool
		err := row.Scan(&isActive)
		if err != nil {
			log.Printf("Error in sanning: %v", err)
		}
		ch <- isActive
	}
}
