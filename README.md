# FastCICD - Go HTTP Server with PostgreSQL

A sample Go HTTP server with PostgreSQL backend, Docker Compose deployment, and GitHub Actions CI/CD pipeline for VPS deployment.

## Features

- Go HTTP server with standard library (`net/http`)
- PostgreSQL integration using `pgx`
- Docker Compose for local development
- GitHub Actions workflow for CI/CD
- Health check endpoint
- REST API for greetings

## Project Structure

```
fastcicd/
├── main.go                    # Main application entry point
├── Dockerfile                 # Multi-stage Docker build
├── docker-compose.yml         # Local development with PostgreSQL
├── .github/workflows/deploy.yml # CI/CD pipeline
├── internal/
│   ├── database/             # Database connection and migrations
│   ├── handlers/             # HTTP handlers
│   └── models/               # Data models
├── .env.example              # Environment variables template
└── README.md                 # This file
```

## Local Development

### Prerequisites

- Go 1.25+
- Docker and Docker Compose

### Quick Start

1. Clone the repository
2. Copy `.env.example` to `.env`:
   ```bash
   cp .env.example .env
   ```
3. Start services with Docker Compose:
   ```bash
   docker compose up -d
   ```
4. Access the application:
   - Web server: http://localhost:8080
   - API endpoints:
     - `GET /` - Hello world with DB connection status
     - `GET /api/greetings` - List all greetings from database
     - `POST /api/greetings/add` - Add new greeting (JSON: `{"message": "text"}`)
     - `GET /health` - Health check endpoint

### Manual Setup (without Docker)

1. Start PostgreSQL:
   ```bash
   docker run -d -p 5432:5432 \
     -e POSTGRES_USER=postgres \
     -e POSTGRES_PASSWORD=postgres \
     -e POSTGRES_DB=fastcicd \
     postgres:16-alpine
   ```

2. Run the application:
   ```bash
   go run main.go
   ```

## API Examples

```bash
# Get all greetings
curl http://localhost:8080/api/greetings

# Add new greeting
curl -X POST http://localhost:8080/api/greetings/add \
  -H "Content-Type: application/json" \
  -d '{"message": "Hello from curl"}'

# Health check
curl http://localhost:8080/health
```

## Deployment to VPS

### GitHub Secrets Setup

Configure the following secrets in your GitHub repository (Settings > Secrets and variables > Actions):

1. `VPS_HOST` - Your VPS IP address or hostname
2. `VPS_USER` - SSH username (e.g., `root` or `deploy`)
3. `VPS_SSH_KEY` - Private SSH key for authentication
4. `VPS_DEPLOY_PATH` - Deployment directory on VPS (e.g., `/opt/fastcicd`)
5. `DOCKER_USERNAME` - Docker Hub username (optional, for image push)
6. `DOCKER_TOKEN` - Docker Hub access token (optional)

### VPS Setup

1. Install Docker and Docker Compose on your VPS
2. Clone the repository to the deployment path:
   ```bash
   git clone https://github.com/yourusername/fastcicd.git /opt/fastcicd
   ```
3. Create `.env` file with production database credentials

### CI/CD Pipeline

The GitHub Actions workflow (`/.github/workflows/deploy.yml`) includes:

1. **Test**: Runs Go tests with PostgreSQL service
2. **Build**: Builds Docker image and pushes to Docker Hub (optional)
3. **Deploy**: SSH into VPS, pull latest changes, and restart containers

The workflow triggers on:
- Push to `main` branch
- Closed pull requests to `main` branch

### Manual Deployment

```bash
# On your VPS
cd /opt/fastcicd
git pull origin main
docker compose down
docker compose pull
docker compose up -d --build
```

## Configuration

Environment variables (set in `.env` file or directly):

| Variable | Default | Description |
|----------|---------|-------------|
| `DB_HOST` | `localhost` | PostgreSQL host |
| `DB_PORT` | `5432` | PostgreSQL port |
| `DB_USER` | `postgres` | Database user |
| `DB_PASSWORD` | `postgres` | Database password |
| `DB_NAME` | `fastcicd` | Database name |
| `DB_SSL_MODE` | `disable` | SSL mode for PostgreSQL |
| `SERVER_PORT` | `8080` | HTTP server port |

## Database Schema

The application creates a `greetings` table on startup:

```sql
CREATE TABLE greetings (
    id SERIAL PRIMARY KEY,
    message TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## License

MIT