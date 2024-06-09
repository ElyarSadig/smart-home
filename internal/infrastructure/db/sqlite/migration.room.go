package sqlite

const ROOM_TABLE_QUERY = `
CREATE TABLE IF NOT EXISTS rooms (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) UNIQUE NOT NULL,
    temperature REAL,
    humidity REAL,
    is_front_door_locked BOOLEAN,
    is_camera_active BOOLEAN,
    is_alarm_active BOOLEAN
);`
