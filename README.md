# Movie Collection API

A RESTful API for managing movie collections built with Go, following clean architecture principles.

## Project Structure

```
.
├── cmd/                    # Application entry points
│   ├── main.go            # Main application file
│   └── initiator/         # Application initialization
│       ├── handlers.go    # Handler initialization
│       ├── initiate.go    # Core initialization logic
│       └── routes.go      # Route configuration
│
├── internal/              # Private application code
│   ├── domain/           # Business entities and interfaces
│   ├── dto/              # Data Transfer Objects
│   ├── handler/          # HTTP request handlers
│   ├── middleware/       # HTTP middleware components
│   ├── repository/       # Data access implementations
│   └── usecase/         # Business logic implementations
│
├── pkg/                  # Public shared packages
│   ├── cloudinary/      # Cloudinary integration
│   ├── db/             # Database configuration
│   ├── response/       # API response utilities
│   └── security/       # Security utilities (JWT, etc.)
│
└── docs/                # Documentation
    └── README.md       # Technical documentation
```

## Prerequisites

- Go 1.16 or higher
- PostgreSQL
- Cloudinary account

## Environment Variables

Create a `.env` file in the project root:

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_user
DB_PASSWORD=your_password
DB_NAME=your_db_name

# JWT Configuration
JWT_SECRET=your_jwt_secret

# Cloudinary Configuration
CLOUDINARY_CLOUD_NAME=your_cloud_name
CLOUDINARY_API_KEY=your_api_key
CLOUDINARY_API_SECRET=your_api_secret
```

## Setup and Installation

1. Clone the repository
```bash
git clone <repository-url>
cd Eskalate-Movie-Project
```

2. Install dependencies
```bash
go mod download
```

3. Set up the database
```bash
# Create PostgreSQL database
createdb your_db_name

# The application will automatically run migrations on startup
```

4. Run the application
```bash
go run cmd/main.go
```

The server will start at `http://localhost:8080`

## API Documentation

Interactive API documentation is available at:
- Swagger UI: `http://localhost:8080/docs`
- OpenAPI Spec: `http://localhost:8080/swagger.yaml`

For detailed API documentation and technical details, please refer to [docs/README.md](docs/README.md).

## Quick API Reference

### Authentication Endpoints
- `POST /signup` - Register a new user
- `POST /login` - Authenticate user and get token

### Movie Endpoints
- `GET /movies` - List all movies (with pagination)
- `GET /movies/:id` - Get movie details
- `POST /movies` - Create a new movie (requires authentication)
- `PUT /movies/:id` - Update a movie (requires authentication)
- `DELETE /movies/:id` - Delete a movie (requires authentication)

## Development

### Code Organization

- **cmd/**: Contains the application's entry point and initialization logic
  - `main.go`: Application entry point
  - `initiator/`: Handles dependency injection and app initialization

- **internal/**: Private application code
  - `domain/`: Business entities and repository interfaces
  - `dto/`: Request/Response data transfer objects
  - `handler/`: HTTP request handlers
  - `middleware/`: HTTP middleware (auth, error handling)
  - `repository/`: Data access layer implementations
  - `templates/` : Swagger Docs
  - `usecase/`: Business logic implementations

- **pkg/**: Shared utilities and external integrations
  - `cloudinary/`: Image upload and management
  - `db/`: Database connection and configuration
  - `response/`: Standardized API response utilities
  - `security/`: Security-related utilities (JWT)
