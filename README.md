# crossword
This project is a crossword puzzle generator built using Go, following Domain-Driven Design (DDD) principles and a microservices architecture.

## Project Structure
- **cmd/api**: Contains the entry point of the application.
- **internal/app**: The core application logic, including API handlers, middleware, services, and domain models.
- **pkg/models**: Contains data models for puzzles and users.
- **scripts**: Contains scripts for database migrations.
- **config.yml**: Configuration settings for the application.
- **db_*.txt**: Definitions for the database structures.

## Setup Instructions
1. **Clone the repository**:
   ```
   git clone <repository-url>
   cd crossword
   ```

2. **Install dependencies**:
   ```
   go mod tidy
   ```

3. **Run database migrations**:
   ```
   ./scripts/migrate.sh
   ```

4. **Start the application**:
   ```
   go run cmd/api/main.go
   ```

## API Endpoints
- **POST /generate**: Generate a new crossword puzzle.
- **GET /puzzles/{id}**: Retrieve a specific puzzle by ID.
- **POST /register**: Register a new user.
- **POST /login**: Log in a user.

## Configuration
The application configuration can be found in `config.yml`. Modify the settings as needed for your environment.

## License
This project is licensed under the MIT License. See the LICENSE file for more details.