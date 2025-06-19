# Movie Collection API - Technical Documentation

## Project Overview

The Movie Collection API is a comprehensive RESTful service designed for managing personal movie collections. Built with Go and following Clean Architecture principles, it provides a robust foundation for movie enthusiasts to catalog, search, and manage their film collections with secure user authentication and cloud-based image storage.

## Architecture

### Clean Architecture Implementation

The project follows Clean Architecture principles with clear separation of concerns:

```
┌─────────────────────────────────────────────────────────────┐
│                    External Interfaces                      │
│  (HTTP Handlers, Database, External Services)              │
├─────────────────────────────────────────────────────────────┤
│                    Use Cases Layer                          │
│  (Business Logic, Application Rules)                       │
├─────────────────────────────────────────────────────────────┤
│                    Domain Layer                             │
│  (Entities, Business Rules, Value Objects)                 │
└─────────────────────────────────────────────────────────────┘
```

### Layer Responsibilities

1. **Domain Layer** (`internal/domain/`)
   - Contains business entities (User, Movie)
   - Defines core business rules and validation
   - Independent of external frameworks

2. **Use Case Layer** (`internal/usecase/`)
   - Implements application-specific business logic
   - Orchestrates domain entities and repositories
   - Handles transaction boundaries

3. **Interface Layer** (`internal/handler/`, `internal/repository/`)
   - HTTP handlers for API endpoints
   - Repository implementations for data access
   - Adapters for external services

4. **Infrastructure Layer** (`pkg/`)
   - Database connections and configurations
   - External service integrations (Cloudinary)
   - Security utilities and middleware

## Core Features

### User Management
- **Registration**: Secure user signup with email/username validation
- **Authentication**: JWT-based login system with password hashing
- **Authorization**: Role-based access control for movie operations

### Movie Collection Management
- **CRUD Operations**: Full create, read, update, delete functionality
- **Image Storage**: Cloudinary integration for movie poster uploads
- **Search & Filter**: Advanced search capabilities with pagination
- **Data Validation**: Comprehensive input validation and sanitization

### Technical Features
- **Pagination**: Efficient data retrieval with configurable page sizes
- **Error Handling**: Centralized error management with appropriate HTTP status codes
- **Input Validation**: Request validation using struct tags and custom validators
- **Security**: JWT tokens, password hashing, and middleware-based protection

## Data Models

### User Entity
```go
type User struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    Username  string    `json:"username" gorm:"unique;not null"`
    Email     string    `json:"email" gorm:"unique;not null"`
    Password  string    `json:"-" gorm:"not null"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    Movies    []Movie   `json:"movies,omitempty" gorm:"foreignKey:UserID"`
}
```

### Movie Entity
```go
type Movie struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    Title       string    `json:"title" gorm:"not null"`
    Description string    `json:"description"`
    Genres      string    `json:"genres"`
    Actors      string    `json:"actors"`
    PosterURL   string    `json:"poster_url"`
    UserID      uint      `json:"user_id" gorm:"not null"`
    User        User      `json:"user,omitempty"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
```

## API Design Principles

### RESTful Endpoints
- Follows REST conventions for resource management
- Consistent URL patterns and HTTP methods
- Proper status code usage for different scenarios

### Response Format
All API responses follow a consistent structure:
```json
{
    "success": true,
    "data": { ... },
    "message": "Operation successful"
}
```

### Error Handling
Centralized error handling with structured error responses:
```json
{
    "success": false,
    "error": "Detailed error message",
    "code": "ERROR_CODE"
}
```

## Security Implementation

### Authentication Flow
1. User registration with password hashing using bcrypt
2. Login with JWT token generation
3. Token-based authentication for protected endpoints
4. Automatic token validation and user context injection

### Authorization
- Owner-based access control for movie operations
- Middleware-based route protection
- User context validation in business logic

### Data Protection
- Password hashing with bcrypt
- Input sanitization and validation
- SQL injection prevention through GORM
- Secure file upload handling

## External Integrations

### Cloudinary Integration
- Secure image upload and storage
- Automatic image optimization
- CDN-based delivery for better performance
- Folder organization for better management

### Database Integration
- PostgreSQL with GORM ORM
- Automatic schema migration
- Connection pooling and optimization
- Transaction management

## Performance Considerations

### Database Optimization
- Proper indexing on frequently queried fields
- Efficient pagination implementation
- Optimized query patterns

### Caching Strategy
- JWT token caching for authentication
- Database connection pooling
- Static asset optimization through Cloudinary CDN

## Scalability Features

### Horizontal Scaling
- Stateless API design
- Database connection pooling
- External service integration for file storage

### Code Organization
- Modular architecture for easy feature addition
- Dependency injection pattern
- Clear separation of concerns

## Testing Strategy

### Unit Testing
- Domain logic testing
- Use case testing with mocked dependencies
- Repository testing with test database

### Integration Testing
- API endpoint testing
- Database integration testing
- External service integration testing

## Monitoring and Logging

### Error Tracking
- Centralized error handling
- Structured logging
- Error categorization and reporting

### Performance Monitoring
- Request/response logging
- Database query monitoring
- External service call tracking

## Future Enhancements

### Planned Features
- Movie rating and review system
- Advanced search with filters
- Movie recommendation engine
- Social features (sharing, following)
- Mobile app support
- Real-time notifications

### Technical Improvements
- GraphQL API support
- Microservices architecture
- Event-driven architecture
- Advanced caching with Redis
- API versioning
- Rate limiting and throttling 