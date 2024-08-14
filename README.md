# Smart Home IoT Backend

This repository hosts the backend for a smart home IoT system built with Go. The backend manages smart devices and rooms, providing APIs to interact with them. It also includes server configuration for production, and simulates device activities.

### Accessing the Front-End

The front-end for this project is developed in Vue.js and is available in a separate repository. You can find it here: [Smart Home IoT Front-End](https://github.com/ElyarSadig/smart-home-iot).


## Table of Contents

- [Smart Home IoT Backend](#smart-home-iot-backend)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Installation](#installation)
  - [Usage](#usage)
    - [Running the Server](#running-the-server)
    - [Accessing the Front-End](#accessing-the-front-end)
  - [API Endpoints](#api-endpoints)
  - [Configuration](#configuration)
  - [Project Structure](#project-structure)
  - [Contributing](#contributing)

## Features

- **Smart Device Management**: Manage devices in various rooms.
- **Device Simulation**: Simulate air conditioner and humidifier operations with two worker goroutines.
- **Graceful Shutdown**: Handles server shutdown cleanly.
- **Embedded Database**: Uses an embedded SQLite database for storage.
- **CORS Support**: Configurable CORS support for API requests.

## Installation

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/yourusername/smart-home-iot-backend.git

   cd smart-home-iot-backend
   ```

2. **Install Dependencies**:

   ```bash
   go mod tidy
   ```

3. **Build the Project**:

   ```bash
   go build -o smart-home-iot-backend main.go
   ```

4. **Run the Server**:
   ```bash
   ./smart-home-iot-backend
   ```

## Usage

### Running the Server

Start the server with:

```bash
./smart-home-iot-backend
```

The server will be accessible at `http://localhost:8080`.

## API Endpoints

- **`GET /device/:id`**: Fetch device details by ID.
- **`PUT /device`**: Update device activity.
- **`GET /room/:id`**: Fetch room details by ID.
- **`PUT /room`**: Update room status.
- **`GET /room/energy-saving/:id`**: Get energy-saving details for a room.

There is also Postman documentation in docs folder.

## Configuration

The application uses an embedded `config.yaml` for configuration, including database settings and server details. Example configuration:

```yaml
database:
  file_path: "./data/smart_home.db"
  max_open_conns: 10
  max_idle_conns: 5
  conn_max_lifetime: 30s

server:
  port: 8080
  read_timeout: 10s
  write_timeout: 10s
  idle_timeout: 30s
  cors_allowed_origins:
    - "*"
```

## Project Structure

```
smart-home-iot-backend/
├── config/           # Configuration files
├── internal/         # Internal application packages
│   ├── domain/       # Domain logic and entities
│   ├── infrastructure/  # Infrastructure-specific code
│   └── service/      # Business logic services
├── main.go           # Entry point for the application
└── README.md         # Project documentation
```

## Contributing

Contributions are welcome! Please fork this repository and submit a pull request for review.

