# MyJunkPal Frontend

A lightweight, vanilla JavaScript web application for tracking food intake and nutrition.

## Overview

This is a single-page application (SPA) built with pure HTML, CSS, and JavaScript - no frameworks required. It provides a clean, intuitive interface for managing food entries and monitoring daily nutrition.

## Features

### 🔐 Authentication
- User registration
- User login
- Session management (persists while backend is running)

### 📊 Dashboard
- Real-time daily nutrition summary
- Interactive date picker to view historical data
- Stats cards showing:
  - Total calories
  - Total protein
  - Total carbs
  - Total fats
- Detailed meal breakdown table

### 🍔 Foods Database
- Browse system foods (15 pre-populated)
- Search and filter foods by name or category
- Create custom foods with full nutritional info
- Delete custom foods
- View complete nutritional breakdown per serving

### 📝 Food Entries
- Log food consumption with:
  - Food selection from database
  - Quantity (servings)
  - Meal type (breakfast, lunch, dinner, snack)
  - Date and time
- Automatic nutrition calculation
- View all entries sorted by date
- Delete entries

### 🎯 Nutrition Goals
- Set daily goals for:
  - Calories
  - Protein
  - Carbs
  - Fats
- Persistent storage of goals

## Tech Stack

- **HTML5** - Semantic markup
- **CSS3** - Custom styling with flexbox and grid
- **Vanilla JavaScript** - No frameworks or libraries
- **Fetch API** - RESTful API communication

## File Structure

```
frontend/
├── index.html      # Main HTML file with all UI components
├── app.js          # JavaScript logic for API calls and UI interactions
└── README.md       # This file
```

## Getting Started

### Prerequisites

- A modern web browser (Chrome, Firefox, Safari, Edge)
- Backend server running on `http://localhost:8080`

### Installation

No installation required! This is a static web application.

### Running the App

1. **Make sure the backend is running:**
   ```bash
   cd ../backend
   go run main.go
   ```

2. **Open the frontend:**
   - **Option 1:** Double-click `index.html` in your file explorer
   - **Option 2:** Open with a local server (recommended):
     ```bash
     # Using Python 3
     python3 -m http.server 3000

     # Using Node.js (with http-server installed)
     npx http-server -p 3000
     ```
   - **Option 3:** Open directly in browser:
     ```bash
     open index.html  # macOS
     ```

3. **Access the app:**
   - If using a local server: `http://localhost:3000`
   - If opening directly: file path in browser

## Usage Guide

### First Time Setup

1. **Register an account:**
   - Enter your name, email, and password
   - Click "Register"
   - You'll be automatically logged in

2. **Set your nutrition goals:**
   - Navigate to the "Goals" tab
   - Enter your daily targets
   - Click "Save Goals"

### Daily Usage

1. **Log your meals:**
   - Go to "Entries" tab
   - Click "Log Food"
   - Select a food from the dropdown
   - Enter quantity (in servings)
   - Select meal type and time
   - Click "Log Entry"

2. **View your progress:**
   - Dashboard shows today's summary by default
   - Use the date picker to view other days
   - See breakdown by meal and total nutrition

3. **Browse foods:**
   - "Foods" tab shows all available foods
   - Use search bar to find specific items
   - Add custom foods if needed

### Adding Custom Foods

1. Navigate to "Foods" tab
2. Click "Add Custom Food"
3. Fill in all nutritional information:
   - Name and category
   - Calories, protein, carbs, fats
   - Serving size and unit
4. Click "Save Food"

## API Integration

The frontend communicates with the backend REST API at `http://localhost:8080/api`:

### Endpoints Used

- `POST /api/auth/register` - User registration
- `POST /api/auth/login` - User login
- `GET /api/foods` - List all foods
- `POST /api/foods` - Create custom food
- `DELETE /api/foods/{id}` - Delete custom food
- `GET /api/entries` - List all entries
- `POST /api/entries` - Log new entry
- `DELETE /api/entries/{id}` - Delete entry
- `GET /api/nutrition/daily/{date}` - Get daily summary
- `GET /api/nutrition/goals` - Get nutrition goals
- `PUT /api/nutrition/goals` - Update nutrition goals

## Customization

### Changing API Base URL

Edit the `API_BASE` constant in `app.js`:

```javascript
const API_BASE = 'http://your-backend-url:port/api';
```

### Styling

All CSS is embedded in `index.html` in the `<style>` tag. Modify colors, fonts, and layouts directly there.

### Adding New Features

The code is organized by feature:
- Auth functions: `register()`, `login()`, `logout()`
- Foods functions: `loadFoods()`, `addFood()`, `deleteFood()`
- Entries functions: `loadEntries()`, `addEntry()`, `deleteEntry()`
- Dashboard functions: `loadDailySummary()`
- Goals functions: `loadGoals()`, `updateGoals()`

## Browser Compatibility

- ✅ Chrome/Edge 90+
- ✅ Firefox 88+
- ✅ Safari 14+
- ✅ Opera 76+

## Limitations

- Single user session (one user logged in at a time on the backend)
- No offline support
- Session resets when backend restarts
- No persistent authentication tokens

## Future Enhancements

Potential features to add:
- Charts and graphs for nutrition trends
- Weekly/monthly views
- Meal planning
- Barcode scanning integration
- Export data to CSV
- Dark mode toggle
- Mobile responsive design improvements
- Progressive Web App (PWA) support

## Troubleshooting

### "Failed to fetch" errors
- Ensure backend is running on port 8080
- Check browser console for CORS errors
- Verify API_BASE URL is correct

### No foods showing
- Make sure you're logged in
- Check that backend has `data/foods.json` populated
- Look for errors in browser console

### Can't log in
- Verify you've registered an account first
- Check that backend `data/users.json` exists
- Try registering a new account

## License

MIT
