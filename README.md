# Crawfish Crawler

A full-stack app to crawl metadata from websites.

## Tech Stack
- Frontend: Next.js (TypeScript, Tailwind)
- Backend: Go (Gin), MySQL
- Auth: JWT (basic)

## Getting Started

```bash
# Start MySQL container
docker compose up -d db

# Start backend
cd backend && go run ./cmd/main.go

# Start frontend
cd frontend && npm run dev
