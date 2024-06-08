package sqlite

const ROOM_TABLE_QUERY = `
CREATE TABLE IF NOT EXISTS rooms (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) NOT NULL,
    temperature REAL,
    humidity REAL
);`
