package sqlite

const DEVICE_TABLE_QUERY = `
CREATE TABLE IF NOT EXISTS devices (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) NOT NULL,
    status BOOLEAN NOT NULL,
    active_from DATETIME,
	active_until DATETIME,
	room_id INTEGER,
	energy_consuming_per_hour REAL,
	FOREIGN KEY (room_id) REFERENCES rooms(id)
);`
