# 🎬 Movie Recommendation System


A production-style **Movie Recommendation System** built using **Go**, **Gin**, **GORM**, and **PostgreSQL** following **Clean Architecture**, **Repository Pattern**, **Dependency Injection**, and **Layered Architecture**.

---

# 📌 Project Overview

This project provides a RESTful backend for managing movies, genres, users, and ratings while generating personalized movie recommendations.

The project demonstrates how enterprise-level Go backend applications are structured using clean architecture and best development practices.

---

# ✨ Features

* 🎭 Genre Management (CRUD)
* 🎬 Movie Management (CRUD)
* 👤 User Management (CRUD)
* ⭐ Rating Management (CRUD)
* 📂 Movies by Genre
* 🎥 Movies Rated by User
* 🏆 Top Rated Movies
* 💡 Personalized Movie Recommendations

---

# 🚀 Tech Stack

| Technology | Description          |
| ---------- | -------------------- |
| Go         | Programming Language |
| Gin        | HTTP Web Framework   |
| GORM       | ORM                  |
| PostgreSQL | Relational Database  |

---

# 🏗️ Architecture

The project follows **Clean Architecture** with strict separation of concerns.

```text
                 HTTP Request
                      │
                      ▼
                  Route Layer
                      │
                      ▼
                 Handler Layer
                      │
                      ▼
              Service Interface
                      │
                      ▼
          Service Implementation
                      │
                      ▼
            Repository Interface
                      │
                      ▼
       Repository Implementation
                      │
                      ▼
                     GORM
                      │
                      ▼
                 PostgreSQL
```

Each layer has a single responsibility, making the application scalable, testable, and maintainable.

---

# 📂 Project Structure

```text
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
│   │   ├── interfaces/
│   │   └── implementations/
│   ├── routes/
│   └── services/
│       ├── interfaces/
│       └── implementations/
│
├── .env
├── go.mod
└── README.md
```

---

# ⚙️ Modules

## 🎭 Genre Module

* Create Genre
* Get All Genres
* Get Genre By ID
* Update Genre
* Delete Genre

---

## 🎬 Movie Module

* Create Movie
* Get All Movies
* Get Movie By ID
* Update Movie
* Delete Movie

---

## 👤 User Module

* Create User
* Get All Users
* Get User By ID
* Update User
* Delete User

---

## ⭐ Rating Module

* Create Rating
* Get All Ratings
* Get Rating By ID
* Update Rating
* Delete Rating

---

## 💡 Recommendation Module

* Get Movies by Genre
* Get Movies Rated by User
* Get Top Rated Movies
* Get Recommended Movies

---

# 🗄️ Database Schema

## Genre

| Column | Type    |
| ------ | ------- |
| id     | Integer |
| name   | String  |

---

## Movie

| Column       | Type    |
| ------------ | ------- |
| id           | Integer |
| title        | String  |
| description  | Text    |
| release_year | Integer |
| genre_id     | Integer |

---

## User

| Column | Type    |
| ------ | ------- |
| id     | Integer |
| name   | String  |
| email  | String  |

---

## Rating

| Column   | Type    |
| -------- | ------- |
| id       | Integer |
| user_id  | Integer |
| movie_id | Integer |
| rating   | Integer |

---

# 🔗 Entity Relationship

```text
           Genre
             │
        1 ───┼──── *
             │
           Movie
             │
        1 ───┼──── *
             │
          Rating
             │
        * ───┼──── 1
             │
            User
```

### Relationships

* One Genre ➜ Many Movies
* One Movie ➜ Many Ratings
* One User ➜ Many Ratings

---

# 📡 REST API Endpoints

## Genre APIs

| Method | Endpoint      |
| ------ | ------------- |
| POST   | `/genres`     |
| GET    | `/genres`     |
| GET    | `/genres/:id` |
| PUT    | `/genres/:id` |
| DELETE | `/genres/:id` |

---

## Movie APIs

| Method | Endpoint      |
| ------ | ------------- |
| POST   | `/movies`     |
| GET    | `/movies`     |
| GET    | `/movies/:id` |
| PUT    | `/movies/:id` |
| DELETE | `/movies/:id` |

---

## User APIs

| Method | Endpoint     |
| ------ | ------------ |
| POST   | `/users`     |
| GET    | `/users`     |
| GET    | `/users/:id` |
| PUT    | `/users/:id` |
| DELETE | `/users/:id` |

---

## Rating APIs

| Method | Endpoint       |
| ------ | -------------- |
| POST   | `/ratings`     |
| GET    | `/ratings`     |
| GET    | `/ratings/:id` |
| PUT    | `/ratings/:id` |
| DELETE | `/ratings/:id` |

---

## Recommendation APIs

| Method | Endpoint                     | Description          |
| ------ | ---------------------------- | -------------------- |
| GET    | `/genres/:id/movies`         | Get movies by genre  |
| GET    | `/users/:id/movies`          | Movies rated by user |
| GET    | `/movies/top-rated`          | Top-rated movies     |
| GET    | `/users/:id/recommendations` | Recommended movies   |

---

# 🧠 Recommendation Logic

### 🎭 Movies by Genre

Returns all movies that belong to the selected genre.

---

### ⭐ Movies Rated by User

Returns every movie rated by the user along with the rating provided.

---

### 🏆 Top Rated Movies

Uses SQL aggregation to calculate:

* Average Rating (`AVG`)
* `GROUP BY`
* `ORDER BY DESC`

Returns movies sorted by highest average rating.

---

### 💡 Recommended Movies

Recommendation logic:

* Excludes movies already rated by the user.
* Calculates average ratings of remaining movies.
* Returns highest-rated unseen movies.

---

# 🛠️ Design Patterns

* ✅ Clean Architecture
* ✅ Repository Pattern
* ✅ Dependency Injection
* ✅ DTO Pattern
* ✅ Interface-Based Programming
* ✅ Layered Architecture

---

# 📊 SQL Concepts Used

* INNER JOIN
* LEFT JOIN
* GROUP BY
* AVG()
* ORDER BY
* WHERE
* Aggregate Functions
* Subqueries

---

# ▶️ Getting Started

## 1️⃣ Clone the Repository

```bash
git clone https://github.com/satishreddykarri/Movie_Recommendation.git
```

---

## 2️⃣ Navigate to the Project

```bash
cd movie-recommendation
```

---

## 3️⃣ Install Dependencies

```bash
go mod tidy
```

---

## 4️⃣ Configure Environment Variables

Create a `.env` file in the project root.

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=movie_recommendation
DB_SSLMODE=disable
```

---

## 5️⃣ Run the Application

```bash
go run ./cmd
```

The server will start on:

```text
http://localhost:8080
```

---

# 📚 Learning Outcomes

Through this project, I gained practical experience in:

* Go Programming
* Gin Framework
* GORM ORM
* PostgreSQL
* Clean Architecture
* Repository Pattern
* Dependency Injection
* DTO Design
* REST API Development
* SQL Joins & Aggregation
* Layered Backend Development
