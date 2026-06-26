**🎬 Movie Recommendation System**



A production-style **Movie Recommendation System** built using **Go**, **Gin**, **GORM**, and **PostgreSQL** following **Clean Architecture**, **Repository Pattern**, **Dependency Injection**, and **Layered Architecture**.



---



**📌 Project Overview**



This project provides a RESTful backend for managing movies and generating movie recommendations. It demonstrates how a real-world Go backend application is structured using clean architecture and industry best practices.



---



**🚀 Tech Stack**



* Go

* Gin

* GORM

* PostgreSQL



---



**🏗️ Architecture**



```text

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

```



---



**📂 Project Structure**



```text

movie-recommendation/



├── cmd/

├── internal/

├── .env

├── go.mod

└── README.md

```



---



**⚙️ Features**



### Genre Module



* Create Genre

* Get All Genres

* Get Genre By ID

* Update Genre

* Delete Genre



### Movie Module



* Create Movie

* Get All Movies

* Get Movie By ID

* Update Movie

* Delete Movie



### User Module



* Create User

* Get All Users

* Get User By ID

* Update User

* Delete User



### Rating Module



* Create Rating

* Get All Ratings

* Get Rating By ID

* Update Rating

* Delete Rating



### Recommendation APIs



* Get Movies by Genre

* Get Movies Rated by User

* Get Top Rated Movies

* Get Recommended Movies



---



**🗄️ Database Schema**



### Genre



| Column |

| ------ |

| id     |

| name   |



### Movie



| Column       |

| ------------ |

| id           |

| title        |

| description  |

| release_year |

| genre_id     |



### User



| Column |

| ------ |

| id     |

| name   |

| email  |



### Rating



| Column   |

| -------- |

| id       |

| user_id  |

| movie_id |

| rating   |



---



**🔗 Relationships**



```text

Genre

   │

   ▼

Movie

   ▲

   │

Rating

   ▲

   │

User

```



---



**📡 REST API Endpoints**



### Genre APIs



| Method | Endpoint      |

| ------ | ------------- |

| POST   | `/genres`     |

| GET    | `/genres`     |

| GET    | `/genres/:id` |

| PUT    | `/genres/:id` |

| DELETE | `/genres/:id` |



### Movie APIs



| Method | Endpoint      |

| ------ | ------------- |

| POST   | `/movies`     |

| GET    | `/movies`     |

| GET    | `/movies/:id` |

| PUT    | `/movies/:id` |

| DELETE | `/movies/:id` |



### User APIs



| Method | Endpoint     |

| ------ | ------------ |

| POST   | `/users`     |

| GET    | `/users`     |

| GET    | `/users/:id` |

| PUT    | `/users/:id` |

| DELETE | `/users/:id` |



### Rating APIs



| Method | Endpoint       |

| ------ | -------------- |

| POST   | `/ratings`     |

| GET    | `/ratings`     |

| GET    | `/ratings/:id` |

| PUT    | `/ratings/:id` |

| DELETE | `/ratings/:id` |



### Recommendation APIs



| Method | Endpoint                     |

| ------ | ---------------------------- |

| GET    | `/genres/:id/movies`         |

| GET    | `/users/:id/movies`          |

| GET    | `/movies/top-rated`          |

| GET    | `/users/:id/recommendations` |



---



**🛠️ Design Patterns Used**



* Repository Pattern

* Dependency Injection

* Layered Architecture

* DTO Pattern

* Interface-Based Programming



---



**📊 SQL Concepts Used**



* INNER JOIN

* GROUP BY

* AVG()

* ORDER BY

* WHERE

* Subqueries

* Aggregate Functions



---



**▶️ Getting Started**



```bash

git clone https://github.com/<your-username>/movie-recommendation.git



cd movie-recommendation



go mod tidy



go run ./cmd

```



---



**📚 What I Learned**



* Go

* Gin Framework

* GORM

* PostgreSQL

* Clean Architecture

* Repository Pattern

* Dependency Injection

* DTO Design

* SQL JOIN Queries

* Aggregate Queries

* REST API Development





