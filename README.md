#API Documentation

https://.postman.co/workspace/My-Workspace~bcea8f84-f843-4381-8975-fb6a17378208/request/36342630-fbb4e508-7bbc-4f2b-8c3b-d9100e9eb79a?action=share&creator=36342630&ctx=documentation&active-environment=36342630-8a999a61-3234-4855-950c-111836a5f1c5

# Task Manager API

A simple task management API built with **Go**, **PostgreSQL**, and **Docker**.  
This project allows users to create, update, delete, and manage tasks with authentication via JWT.

## 📂 Project Structure

.
├── cmd/ # Main application entry point
│ └── main.go
├── config/ # Configuration loading (Viper)
├── models/ # Database models
├── handlers/ # HTTP request handlers
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── .env

## 🚀 Features

- JWT-based authentication  
- CRUD operations for tasks  
- PostgreSQL as the database  
- Environment-based config loading (local & container)  
- Dockerized for easy setup  

## ⚙️ Requirements

- [Go 1.22+](https://golang.org/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## 🛠️ Environment Variables

Create a `.env` file in the **root** directory:

```env
# App
PORT=:8080
JWT_SECRET_KEY=your-secret-key

# Database
POSTGRES_USER=postgres
POSTGRES_PASSWORD=your-password
POSTGRES_DB=taskmanager
DB_PORT=5432

# Local DB connection
DB_URL=postgres://postgres:your-password@localhost:5432/taskmanager

cp .env.example .env
# Edit .env with your real credentials

🏃 Running Locally

# Install dependencies
go mod tidy

# Start PostgreSQL locally (if installed)
# Make sure DB_URL in .env points to your local database

# Run the app
go run cmd/main.go

🐳 Running with Docker

# Build and start containers
docker-compose up --build

# Stop containers
docker-compose down

📡 API Endpoints (Example)
Method	Endpoint	Description	Auth Required
POST	/register	Register new user	❌
POST	/login	    Login & get JWT	    ❌
GET	    /tasks	    Get all tasks	    ✅
POST	/tasks	    Create a new task	✅
PUT	    /tasks/{id}	Update a task	    ✅
DELETE	/tasks/{id}	Delete a task	    ✅

📝 License
This project is for personal use. Feel free to fork and modify


