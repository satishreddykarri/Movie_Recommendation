🎬 Movie Recommendation System

A production-style Movie Recommendation System built using Go, Gin, GORM, and PostgreSQL following Clean Architecture, Repository Pattern, Dependency Injection, and Layered Architecture.

📌 Project Overview

This project allows users to:

Manage Genres
Manage Movies
Manage Users
Rate Movies
View Movies by Genre
View Movies Rated by a User
View Top Rated Movies
Get Movie Recommendations

The project is designed to demonstrate how real-world backend applications are structured using Go.

🚀 Tech Stack
Go
Gin
GORM
PostgreSQL
🏗️ Architecture

The project follows a layered architecture.

Route
    ↓
Handler
    ↓
Service Interface
    ↓
Service Implementation
    ↓
Repository Interface
    ↓
Repository Implementation
    ↓
GORM
    ↓
PostgreSQL

Each layer has a single responsibility.

📂 Project Structure
movie-recommendation/

├── cmd/
│   └── main.go
│
├── internal/
│   ├── app/
│   ├── config/
│   ├── database/
│   ├── dto/
│   ├── handler/
│   ├── middleware/
│   ├── models/
│   ├── repositories/
│   │   ├── implementations/
│   │   └── interfaces/
│   ├── routes/
│   └── services/
│       ├── implementations/
│       └── interfaces/
│
├── .env
├── go.mod
└── README.md
⚙️ Features
Genre Module
Create Genre
Get All Genres
Get Genre By ID
Update Genre
Delete Genre
Movie Module
Create Movie
Get All Movies
Get Movie By ID
Update Movie
Delete Movie
User Module
Create User
Get All Users
Get User By ID
Update User
Delete User
Rating Module
Create Rating
Get All Ratings
Get Rating By ID
Update Rating
Delete Rating
Recommendation APIs
Get Movies By Genre
Get Movies Rated By User
Get Top Rated Movies
Get Recommended Movies
🗄️ Database Design
Genre
Column
id
name
Movie
Column
id
title
description
release_year
genre_id
User
Column
id
name
email
Rating
Column
id
user_id
movie_id
rating
🔗 Relationships
Genre
   │
   │ 1
   ▼
Movie
   ▲
   │
Rating
   ▲
   │
User
One Genre → Many Movies
One User → Many Ratings
One Movie → Many Ratings
📡 REST APIs
Genre
Method	Endpoint
POST	/genres
GET	/genres
GET	/genres/:id
PUT	/genres/:id
DELETE	/genres/:id
Movie
Method	Endpoint
POST	/movies
GET	/movies
GET	/movies/:id
PUT	/movies/:id
DELETE	/movies/:id
User
Method	Endpoint
POST	/users
GET	/users
GET	/users/:id
PUT	/users/:id
DELETE	/users/:id
Rating
Method	Endpoint
POST	/ratings
GET	/ratings
GET	/ratings/:id
PUT	/ratings/:id
DELETE	/ratings/:id
Recommendation APIs
Method	Endpoint
GET	/genres/:id/movies
GET	/users/:id/movies
GET	/movies/top-rated
GET	/users/:id/recommendations
🧠 Recommendation Logic
Movies By Genre

Returns all movies that belong to a selected genre.

Movies Rated By User

Returns all movies rated by a specific user along with the rating.

Top Rated Movies

Calculates the average rating for each movie and returns them in descending order.

Recommended Movies

Returns movies that:

Have not been rated by the selected user.
Are ordered by the highest average rating.
🛠️ Design Patterns Used
Repository Pattern
Dependency Injection
Layered Architecture
DTO Pattern
Interface-based Programming
📊 SQL Concepts Used
INNER JOIN
GROUP BY
AVG()
ORDER BY
WHERE
Subqueries
Aggregate Functions
▶️ Getting Started
Clone the Repository
git clone <repository-url>
Navigate to Project
cd movie-recommendation
Install Dependencies
go mod tidy
Configure Environment

Create a .env file.

Example:

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=movie_recommendation
DB_SSLMODE=disable
Run the Application
go run ./cmd
📚 What I Learned

Through this project I gained hands-on experience with:

Go
Gin Framework
GORM
PostgreSQL
Clean Architecture
Repository Pattern
Dependency Injection
DTOs
SQL JOINs
Aggregate Queries
Layered Backend Design
REST API Development
