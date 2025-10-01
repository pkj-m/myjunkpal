# MyJunkPal API

A RESTful API for tracking food intake and nutrition, similar to MyFitnessPal.

## Getting Started

### Prerequisites
- Go 1.16 or higher

### Installation

1. Clone the repository
2. Install dependencies:
```bash
go mod tidy
```

3. Run the server:
```bash
go run main.go
```

The server will start on `http://localhost:8080`

## API Documentation

### Authentication

#### Register User
```http
POST /api/auth/register
```

**Request Body:**
```json
{
  "email": "user@example.com",
  "name": "John Doe",
  "password": "password123"
}
```

**Response:** `200 OK`
```json
{
  "id": "uuid",
  "email": "user@example.com",
  "name": "John Doe",
  "password": "",
  "daily_calorie_goal": 2000,
  "daily_protein_goal": 150,
  "daily_carbs_goal": 250,
  "daily_fats_goal": 65,
  "created_at": "2025-10-01T12:00:00Z"
}
```

#### Login
```http
POST /api/auth/login
```

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

**Response:** `200 OK`
```json
{
  "id": "uuid",
  "email": "user@example.com",
  "name": "John Doe",
  "daily_calorie_goal": 2000,
  "daily_protein_goal": 150,
  "daily_carbs_goal": 250,
  "daily_fats_goal": 65,
  "created_at": "2025-10-01T12:00:00Z"
}
```

#### Get Current User
```http
GET /api/users/me
```

**Response:** `200 OK`
```json
{
  "id": "uuid",
  "email": "user@example.com",
  "name": "John Doe",
  "daily_calorie_goal": 2000,
  "daily_protein_goal": 150,
  "daily_carbs_goal": 250,
  "daily_fats_goal": 65,
  "created_at": "2025-10-01T12:00:00Z"
}
```

---

### Foods

#### List Foods
```http
GET /api/foods
```

**Query Parameters:**
- `name` (optional): Filter by food name (case-insensitive)
- `category` (optional): Filter by category

**Response:** `200 OK`
```json
[
  {
    "id": "food-1",
    "user_id": "",
    "name": "Chicken Breast",
    "calories": 165,
    "protein": 31,
    "carbs": 0,
    "fats": 3.6,
    "serving_size": 100,
    "serving_unit": "g",
    "category": "protein",
    "created_at": "2025-10-01T00:00:00Z"
  }
]
```

#### Get Food by ID
```http
GET /api/foods/{id}
```

**Response:** `200 OK`
```json
{
  "id": "food-1",
  "user_id": "",
  "name": "Chicken Breast",
  "calories": 165,
  "protein": 31,
  "carbs": 0,
  "fats": 3.6,
  "serving_size": 100,
  "serving_unit": "g",
  "category": "protein",
  "created_at": "2025-10-01T00:00:00Z"
}
```

#### Create Custom Food
```http
POST /api/foods
```

**Request Body:**
```json
{
  "name": "Custom Protein Shake",
  "calories": 250,
  "protein": 30,
  "carbs": 15,
  "fats": 5,
  "serving_size": 300,
  "serving_unit": "ml",
  "category": "beverage"
}
```

**Response:** `201 Created`
```json
{
  "id": "uuid",
  "user_id": "user-uuid",
  "name": "Custom Protein Shake",
  "calories": 250,
  "protein": 30,
  "carbs": 15,
  "fats": 5,
  "serving_size": 300,
  "serving_unit": "ml",
  "category": "beverage",
  "created_at": "2025-10-01T12:00:00Z"
}
```

#### Update Food
```http
PUT /api/foods/{id}
```

**Request Body:**
```json
{
  "name": "Updated Food Name",
  "calories": 200,
  "protein": 25,
  "carbs": 10,
  "fats": 8,
  "serving_size": 100,
  "serving_unit": "g",
  "category": "protein"
}
```

**Response:** `200 OK`

#### Delete Food
```http
DELETE /api/foods/{id}
```

**Response:** `204 No Content`

---

### Food Entries

#### List Entries
```http
GET /api/entries
```

**Query Parameters:**
- `start_date` (optional): Filter entries from this date (format: YYYY-MM-DD)
- `end_date` (optional): Filter entries until this date (format: YYYY-MM-DD)
- `meal_type` (optional): Filter by meal type (breakfast, lunch, dinner, snack)

**Response:** `200 OK`
```json
[
  {
    "id": "entry-uuid",
    "user_id": "user-uuid",
    "food_id": "food-1",
    "food_name": "Chicken Breast",
    "quantity": 1.5,
    "meal_type": "lunch",
    "eaten_at": "2025-10-01T12:30:00Z",
    "calories": 247.5,
    "protein": 46.5,
    "carbs": 0,
    "fats": 5.4,
    "created_at": "2025-10-01T12:30:00Z"
  }
]
```

#### Get Entry by ID
```http
GET /api/entries/{id}
```

**Response:** `200 OK`

#### Create Entry
```http
POST /api/entries
```

**Request Body:**
```json
{
  "food_id": "food-1",
  "quantity": 1.5,
  "meal_type": "lunch",
  "eaten_at": "2025-10-01T12:30:00Z"
}
```

**Response:** `201 Created`
```json
{
  "id": "entry-uuid",
  "user_id": "user-uuid",
  "food_id": "food-1",
  "food_name": "Chicken Breast",
  "quantity": 1.5,
  "meal_type": "lunch",
  "eaten_at": "2025-10-01T12:30:00Z",
  "calories": 247.5,
  "protein": 46.5,
  "carbs": 0,
  "fats": 5.4,
  "created_at": "2025-10-01T12:30:00Z"
}
```

#### Update Entry
```http
PUT /api/entries/{id}
```

**Request Body:**
```json
{
  "quantity": 2.0,
  "meal_type": "dinner",
  "eaten_at": "2025-10-01T18:30:00Z"
}
```

**Response:** `200 OK`

#### Delete Entry
```http
DELETE /api/entries/{id}
```

**Response:** `204 No Content`

---

### Nutrition Summary

#### Get Daily Summary
```http
GET /api/nutrition/daily/{date}
```

**Path Parameters:**
- `date`: Date in YYYY-MM-DD format (e.g., 2025-10-01)

**Response:** `200 OK`
```json
{
  "date": "2025-10-01",
  "calories": 1850.5,
  "protein": 142.3,
  "carbs": 185.2,
  "fats": 62.1,
  "entries": [
    {
      "id": "entry-1",
      "user_id": "user-uuid",
      "food_id": "food-1",
      "food_name": "Chicken Breast",
      "quantity": 1.5,
      "meal_type": "lunch",
      "eaten_at": "2025-10-01T12:30:00Z",
      "calories": 247.5,
      "protein": 46.5,
      "carbs": 0,
      "fats": 5.4,
      "created_at": "2025-10-01T12:30:00Z"
    }
  ]
}
```

#### Get Weekly Summary
```http
GET /api/nutrition/weekly
```

**Query Parameters:**
- `start_date` (optional): Starting date for the 7-day period (format: YYYY-MM-DD). Defaults to today.

**Response:** `200 OK`
```json
[
  {
    "date": "2025-10-01",
    "calories": 1850.5,
    "protein": 142.3,
    "carbs": 185.2,
    "fats": 62.1,
    "entries": [...]
  },
  {
    "date": "2025-09-30",
    "calories": 2100.0,
    "protein": 155.0,
    "carbs": 220.0,
    "fats": 70.0,
    "entries": [...]
  }
]
```

#### Get Nutrition Goals
```http
GET /api/nutrition/goals
```

**Response:** `200 OK`
```json
{
  "daily_calorie_goal": 2000,
  "daily_protein_goal": 150,
  "daily_carbs_goal": 250,
  "daily_fats_goal": 65
}
```

#### Update Nutrition Goals
```http
PUT /api/nutrition/goals
```

**Request Body:**
```json
{
  "daily_calorie_goal": 2200,
  "daily_protein_goal": 165,
  "daily_carbs_goal": 275,
  "daily_fats_goal": 73
}
```

**Response:** `200 OK`

---

## Pre-populated Foods

The system comes with 15 pre-populated foods:

| ID | Name | Calories | Protein | Carbs | Fats | Serving Size | Category |
|----|------|----------|---------|-------|------|--------------|----------|
| food-1 | Chicken Breast | 165 | 31g | 0g | 3.6g | 100g | protein |
| food-2 | White Rice | 130 | 2.7g | 28g | 0.3g | 100g | grain |
| food-3 | Broccoli | 55 | 3.7g | 11g | 0.6g | 100g | vegetable |
| food-4 | Banana | 89 | 1.1g | 23g | 0.3g | 100g | fruit |
| food-5 | Eggs | 155 | 13g | 1.1g | 11g | 100g | protein |
| food-6 | Salmon | 208 | 20g | 0g | 13g | 100g | protein |
| food-7 | Oatmeal | 389 | 17g | 66g | 7g | 100g | grain |
| food-8 | Almonds | 579 | 21g | 22g | 50g | 100g | nuts |
| food-9 | Sweet Potato | 86 | 1.6g | 20g | 0.1g | 100g | vegetable |
| food-10 | Greek Yogurt | 59 | 10g | 3.6g | 0.4g | 100g | dairy |
| food-11 | Apple | 52 | 0.3g | 14g | 0.2g | 100g | fruit |
| food-12 | Avocado | 160 | 2g | 9g | 15g | 100g | fruit |
| food-13 | Brown Rice | 111 | 2.6g | 23g | 0.9g | 100g | grain |
| food-14 | Spinach | 23 | 2.9g | 3.6g | 0.4g | 100g | vegetable |
| food-15 | Peanut Butter | 588 | 25g | 20g | 50g | 100g | nuts |

---

## Error Responses

### 400 Bad Request
```json
{
  "error": "Invalid request body"
}
```

### 401 Unauthorized
```json
{
  "error": "Unauthorized"
}
```

### 403 Forbidden
```json
{
  "error": "Forbidden"
}
```

### 404 Not Found
```json
{
  "error": "Resource not found"
}
```

### 409 Conflict
```json
{
  "error": "User already exists"
}
```

### 500 Internal Server Error
```json
{
  "error": "Internal server error"
}
```

---

## Data Storage

Data is persisted to JSON files in the `data/` directory:
- `users.json` - User accounts
- `foods.json` - Food database (system + custom foods)
- `entries.json` - Food intake entries

---

## Authentication

The API uses in-memory session management. After logging in, the user session is maintained in memory until the server restarts. For MVP purposes, only one user can be logged in at a time.

---

## Example Usage

### Complete Workflow

1. **Register a new user:**
```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"john@example.com","name":"John Doe","password":"pass123"}'
```

2. **Login (if not already logged in):**
```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"john@example.com","password":"pass123"}'
```

3. **Browse available foods:**
```bash
curl http://localhost:8080/api/foods
```

4. **Log breakfast:**
```bash
curl -X POST http://localhost:8080/api/entries \
  -H "Content-Type: application/json" \
  -d '{
    "food_id": "food-7",
    "quantity": 0.5,
    "meal_type": "breakfast",
    "eaten_at": "2025-10-01T08:00:00Z"
  }'
```

5. **Check daily nutrition:**
```bash
curl http://localhost:8080/api/nutrition/daily/2025-10-01
```

6. **Set nutrition goals:**
```bash
curl -X PUT http://localhost:8080/api/nutrition/goals \
  -H "Content-Type: application/json" \
  -d '{
    "daily_calorie_goal": 2500,
    "daily_protein_goal": 180,
    "daily_carbs_goal": 300,
    "daily_fats_goal": 80
  }'
```

---

## License

MIT
