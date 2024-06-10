package sqlite

const DEVICE_TABLE_QUERY = `
CREATE TABLE IF NOT EXISTS devices (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) UNIQUE NOT NULL,
    is_active BOOLEAN NOT NULL,
    updated_at DATETIME,
	room_id INTEGER,
	energy_consuming_per_hour REAL NOT NULL,
	FOREIGN KEY (room_id) REFERENCES rooms(id)
);`
