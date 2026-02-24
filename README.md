# Restaurant REST API

A REST API built with Go for managing a restaurant's menu and orders.

## Tech Stack

- **Go** with [Gorilla Mux](https://github.com/gorilla/mux) for routing
- **PostgreSQL** with [pgx](https://github.com/jackc/pgx) for database access
- **godotenv** for environment variable management

## Project Structure

```
RestAPI/
├── Main.go                  # Entry point, routes, dependency wiring
├── models/
│   ├── menu.go              # Menu item structs (Burger, Drink, Sides, Dessert)
│   └── order.go             # Order and Item structs
├── handlers/
│   ├── menu.go              # HTTP handlers for menu endpoints
│   └── order.go             # HTTP handlers for order endpoints
├── service/
│   ├── menuService.go       # Menu business logic
│   └── orderService.go      # Order business logic
├── db/
│   ├── db.go                # Database connection pool
│   └── repository/
│       ├── menuRepo.go      # Menu database queries
│       └── orderRepo.go     # Order database queries
└── .env                     # Database connection string (not committed)
```

## Setup

### Prerequisites

- Go 1.25+
- PostgreSQL

### Database

1. Create the database:
   ```sql
   CREATE DATABASE mydb;
   ```

2. Create the tables:
   ```sql
   CREATE TABLE burgers (
       id SERIAL PRIMARY KEY,
       name VARCHAR(100) NOT NULL,
       price DECIMAL(5,2) NOT NULL,
       calories INT NOT NULL,
       is_meal BOOLEAN DEFAULT false
   );

   CREATE TABLE drinks (
       id SERIAL PRIMARY KEY,
       name VARCHAR(100) NOT NULL,
       price DECIMAL(5,2) NOT NULL,
       calories INT NOT NULL,
       is_iced BOOLEAN DEFAULT false
   );

   CREATE TABLE sides (
       id SERIAL PRIMARY KEY,
       name VARCHAR(100) NOT NULL,
       price DECIMAL(5,2) NOT NULL,
       calories INT NOT NULL
   );

   CREATE TABLE desserts (
       id SERIAL PRIMARY KEY,
       name VARCHAR(100) NOT NULL,
       price DECIMAL(5,2) NOT NULL,
       calories INT NOT NULL
   );

   CREATE TABLE orders (
       order_number BIGSERIAL PRIMARY KEY,
       total_price DECIMAL(10,2) NOT NULL,
       date_time TIMESTAMP DEFAULT NOW(),
       status VARCHAR(20) DEFAULT 'pending'
   );
   ```

### Environment

Create a `.env` file in the project root:

```
DATABASE_URL=postgres://postgres:yourpassword@localhost:5432/mydb?sslmode=disable
```

### Run

```bash
go run Main.go
```

Server starts on `http://localhost:8080`.

## API Endpoints

### Menu

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/menu/` | Get full menu (burgers, drinks, sides, desserts) |

### Orders

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/orders/` | Get all orders |
| GET | `/orders/{id}` | Get order by ID |
| POST | `/orders` | Create a new order |
| POST | `/orders/{id}` | Update order total price |
| PATCH | `/orders/{id}` | Change order status |
| DELETE | `/orders/{id}` | Delete an order |

### Example Requests

**Create an order:**
```bash
curl -X POST http://localhost:8080/orders \
  -H "Content-Type: application/json" \
  -d '[{"menu_item": {"name": "Classic Burger", "price": 5.99}, "quantity": 2}]'
```

**Change order status:**
```bash
curl -X PATCH http://localhost:8080/orders/1 \
  -H "Content-Type: application/json" \
  -d '{"status": "completed"}'
```
