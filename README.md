# 🍔 MyJunkPal

A full-stack food tracking application similar to MyFitnessPal, built with Go backend and vanilla JavaScript frontend.

## Overview

MyJunkPal is a lightweight food and nutrition tracking system that allows users to:
- Track daily food intake
- Monitor macronutrients (calories, protein, carbs, fats)
- Set and manage nutrition goals
- Build custom food databases
- View daily and weekly nutrition summaries

## Project Structure

```
myjunkpal/
├── backend/                 # Go REST API server
│   ├── main.go             # Server entry point
│   ├── handlers/           # HTTP request handlers
│   │   ├── auth.go
│   │   ├── foods.go
│   │   ├── entries.go
│   │   └── nutrition.go
│   ├── models/             # Data models
│   │   ├── user.go
│   │   ├── food.go
│   │   └── entry.go
│   ├── storage/            # JSON file storage layer
│   │   └── json_store.go
│   ├── middleware/         # Authentication middleware
│   │   └── auth.go
│   ├── data/               # JSON data files
│   │   ├── users.json
│   │   ├── foods.json
│   │   └── entries.json
│   ├── go.mod
│   ├── go.sum
│   └── README.md
├── frontend/               # Vanilla JS web app
│   ├── index.html
│   ├── app.js
│   └── README.md
└── README.md              # This file
```

## Features

### Backend (Go)
- ✅ RESTful API with JSON responses
- ✅ User authentication (register/login)
- ✅ CRUD operations for foods and entries
- ✅ Daily and weekly nutrition summaries
- ✅ Custom nutrition goals
- ✅ JSON file-based persistence
- ✅ CORS enabled for frontend integration
- ✅ In-memory session management

### Frontend (Vanilla JavaScript)
- ✅ Single-page application (no frameworks)
- ✅ User authentication UI
- ✅ Interactive dashboard with daily stats
- ✅ Food database browser with search
- ✅ Food entry logging interface
- ✅ Nutrition goals management
- ✅ Responsive design with clean UI

## Tech Stack

### Backend
- **Language:** Go 1.16+
- **Router:** gorilla/mux
- **CORS:** rs/cors
- **UUID:** google/uuid
- **Storage:** JSON files

### Frontend
- **HTML5** - Semantic markup
- **CSS3** - Modern styling with flexbox/grid
- **JavaScript (ES6+)** - Vanilla JS, no frameworks
- **Fetch API** - RESTful communication

## Quick Start

### Prerequisites

- Go 1.16 or higher
- Modern web browser

### Installation

1. **Clone the repository:**
   ```bash
   git clone <repository-url>
   cd myjunkpal
   ```

2. **Install backend dependencies:**
   ```bash
   cd backend
   go mod tidy
   ```

### Running the Application

1. **Start the backend server:**
   ```bash
   cd backend
   go run main.go
   ```
   Server will start on `http://localhost:8080`

2. **Open the frontend:**
   ```bash
   cd frontend
   open index.html
   ```
   Or use a local server:
   ```bash
   python3 -m http.server 3000
   # Then open http://localhost:3000
   ```

3. **Register and start tracking!**
   - Create a new account
   - Browse the 15 pre-populated foods
   - Log your meals
   - Track your progress

## API Documentation

Full API documentation is available in `backend/README.md`.

### Key Endpoints

#### Authentication
- `POST /api/auth/register` - Register new user
- `POST /api/auth/login` - Login
- `GET /api/users/me` - Get current user

#### Foods
- `GET /api/foods` - List foods (with filters)
- `POST /api/foods` - Create custom food
- `PUT /api/foods/{id}` - Update food
- `DELETE /api/foods/{id}` - Delete food

#### Entries
- `GET /api/entries` - List entries (with filters)
- `POST /api/entries` - Log food entry
- `PUT /api/entries/{id}` - Update entry
- `DELETE /api/entries/{id}` - Delete entry

#### Nutrition
- `GET /api/nutrition/daily/{date}` - Daily summary
- `GET /api/nutrition/weekly` - Weekly summary
- `GET /api/nutrition/goals` - Get goals
- `PUT /api/nutrition/goals` - Update goals

## Usage Example

### 1. Register a User
```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "name": "John Doe",
    "password": "password123"
  }'
```

### 2. Log a Food Entry
```bash
curl -X POST http://localhost:8080/api/entries \
  -H "Content-Type: application/json" \
  -d '{
    "food_id": "food-1",
    "quantity": 1.5,
    "meal_type": "lunch",
    "eaten_at": "2025-10-01T12:30:00Z"
  }'
```

### 3. Get Daily Summary
```bash
curl http://localhost:8080/api/nutrition/daily/2025-10-01
```

## Pre-populated Foods

The system comes with 15 nutritious foods:

| Food | Category | Calories | Protein |
|------|----------|----------|---------|
| Chicken Breast | protein | 165 | 31g |
| White Rice | grain | 130 | 2.7g |
| Broccoli | vegetable | 55 | 3.7g |
| Banana | fruit | 89 | 1.1g |
| Eggs | protein | 155 | 13g |
| Salmon | protein | 208 | 20g |
| Oatmeal | grain | 389 | 17g |
| Almonds | nuts | 579 | 21g |
| Sweet Potato | vegetable | 86 | 1.6g |
| Greek Yogurt | dairy | 59 | 10g |

*See `backend/data/foods.json` for complete list*

## Architecture

### Data Flow

```
User → Frontend (HTML/JS) → Backend API (Go) → JSON Storage
                            ↓
                     Authentication
                     (In-Memory Session)
```

### Storage

- **Users:** `backend/data/users.json`
- **Foods:** `backend/data/foods.json`
- **Entries:** `backend/data/entries.json`

Data is persisted to disk as JSON files, making it easy to inspect and debug.

### Authentication

Simple in-memory session management:
- Single user can be logged in at a time (MVP)
- Session resets on server restart
- No JWT tokens or cookies (for simplicity)

## Development

### Backend Development

```bash
cd backend

# Run server
go run main.go

# Run tests (if added)
go test ./...

# Build binary
go build -o myjunkpal
```

### Frontend Development

Simply edit `frontend/index.html` or `frontend/app.js` and refresh your browser.

For better development experience, use a local server with live reload.

### Adding New Features

#### Add a New API Endpoint

1. Create handler function in appropriate `handlers/*.go` file
2. Register route in `backend/main.go`
3. Update frontend `app.js` with new API call
4. Add UI components in `index.html`

## Configuration

### Backend Port

Change port in `backend/main.go`:
```go
http.ListenAndServe(":8080", corsHandler.Handler(r))
```

### Frontend API URL

Change API base URL in `frontend/app.js`:
```javascript
const API_BASE = 'http://localhost:8080/api';
```

### Data Directory

Change data storage location in `backend/main.go`:
```go
store := storage.NewJSONStore("./data")
```

## Limitations (MVP)

- Single user login at a time
- In-memory session (no persistence)
- No password hashing
- JSON file storage (not suitable for production)
- No data validation
- No rate limiting
- No authentication tokens

## Future Roadmap

### Backend Enhancements
- [ ] PostgreSQL/MySQL database
- [ ] JWT token authentication
- [ ] Password hashing (bcrypt)
- [ ] Input validation
- [ ] Rate limiting
- [ ] Unit tests
- [ ] Docker deployment
- [ ] API versioning

### Frontend Enhancements
- [ ] Charts and graphs (nutrition trends)
- [ ] Mobile responsive design
- [ ] Progressive Web App (PWA)
- [ ] Meal planning
- [ ] Recipe calculator
- [ ] Export data to CSV/PDF
- [ ] Dark mode
- [ ] Barcode scanning

### Features
- [ ] Multi-user support
- [ ] Social features (share meals)
- [ ] Exercise tracking
- [ ] Weight tracking
- [ ] Photo uploads for meals
- [ ] Meal templates/favorites
- [ ] Notifications and reminders

## Contributing

This is a learning/demo project. Feel free to fork and extend!

## Troubleshooting

### Backend won't start
- Check if port 8080 is already in use
- Ensure `data/` directory exists
- Run `go mod tidy` to install dependencies

### Frontend can't connect
- Verify backend is running on port 8080
- Check browser console for CORS errors
- Ensure API_BASE URL matches backend

### Data not persisting
- Check file permissions on `data/` directory
- Verify JSON files are not corrupted
- Restart backend server

## License

MIT License - feel free to use this project for learning or as a starting point for your own application.

## Acknowledgments

Built as a demonstration of:
- RESTful API design
- Go web development
- Vanilla JavaScript SPA
- Full-stack application architecture

---

**Note:** This is an MVP built for demonstration purposes. For production use, implement proper security measures, use a real database, and add comprehensive error handling.
