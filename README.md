# Real Estate Listings API

A robust REST API built with Go (Gin framework) for managing real estate property listings. This API provides endpoints for creating, retrieving, and managing property listings with features like pagination and realtor authentication.

## 🚀 Features

- Property listing management
- Pagination support
- Image handling for multiple property photos
- Realtor authentication
- Duplicate listing prevention
- Transaction support for data integrity

## 🛠️ Tech Stack

- **Go** - Primary programming language
- **Gin** - Web framework
- **GORM** - ORM for database operations
- **PostgreSQL** - Database (assumed based on the codebase)
- **UUID** - For unique identifier generation

## 📋 API Endpoints

### Listings

#### Get Listings

GET /listings?page=1&limit=10

Query Parameters:

- `page` (optional) - Page number (default: 1)
- `limit` (optional) - Items per page (default: 10)

Response:

```json
{
    "status": "OK",
    "data": [...],
    "total": 100,
    "current_page": 1,
    "total_page": 10,
    "per_page": 10
}
```

#### Create Listing

```http
POST /listings
```

Required fields in request body:

- `title` - Property title
- `address` - Property address
- `price` - Property price
- ... (other fields as defined in CreateListingInput)

## 🏗️ Project Structure

```
src/
├── api/
│   └── listings.go # API handlers
├── config/
│   └── database.go # Database configuration
├── models/
│   └── listing.go # Data models
└── repositories/
    └── listing.repo.go # Database operations
```
