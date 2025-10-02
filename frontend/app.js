const API_BASE = 'http://localhost:8080/api';

let currentUser = null;
let allFoods = [];

// Initialize
document.addEventListener('DOMContentLoaded', () => {
    const today = new Date().toISOString().split('T')[0];
    document.getElementById('summaryDate').value = today;
});

// Auth Functions
async function register() {
    const name = document.getElementById('registerName').value;
    const email = document.getElementById('registerEmail').value;
    const password = document.getElementById('registerPassword').value;

    try {
        const response = await fetch(`${API_BASE}/auth/register`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ name, email, password })
        });

        if (response.ok) {
            const user = await response.json();
            currentUser = user;
            showApp();
            alert('Registration successful!');
        } else {
            const error = await response.text();
            alert('Registration failed: ' + error);
        }
    } catch (err) {
        alert('Error: ' + err.message);
    }
}

async function login() {
    const email = document.getElementById('loginEmail').value;
    const password = document.getElementById('loginPassword').value;

    try {
        const response = await fetch(`${API_BASE}/auth/login`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ email, password })
        });

        if (response.ok) {
            const user = await response.json();
            currentUser = user;
            showApp();
            alert('Login successful!');
        } else {
            const error = await response.text();
            alert('Login failed: ' + error);
        }
    } catch (err) {
        alert('Error: ' + err.message);
    }
}

function logout() {
    currentUser = null;
    document.getElementById('authSection').classList.remove('hidden');
    document.getElementById('authSection').classList.add('active');
    document.getElementById('appSection').classList.add('hidden');
    document.getElementById('userDisplay').textContent = 'Not logged in';
}

function showApp() {
    document.getElementById('authSection').classList.add('hidden');
    document.getElementById('authSection').classList.remove('active');
    document.getElementById('appSection').classList.remove('hidden');
    document.getElementById('userDisplay').textContent = `Logged in as ${currentUser.name}`;

    // Load initial data
    loadFoods();
    loadEntries();
    loadDailySummary();
    loadGoals();
}

// Tab Navigation
function showTab(tabName) {
    // Hide all tabs
    document.querySelectorAll('.section').forEach(section => {
        section.classList.remove('active');
    });

    // Remove active class from all tab buttons
    document.querySelectorAll('.tab-btn').forEach(btn => {
        btn.classList.remove('active');
    });

    // Show selected tab
    document.getElementById(tabName + 'Tab').classList.add('active');

    // Add active class to clicked button
    event.target.classList.add('active');
}

// Foods Functions
async function loadFoods() {
    try {
        const response = await fetch(`${API_BASE}/foods`);
        if (response.ok) {
            allFoods = await response.json();
            displayFoods(allFoods);
            populateFoodSelect();
        }
    } catch (err) {
        alert('Error loading foods: ' + err.message);
    }
}

function displayFoods(foods) {
    const tbody = document.getElementById('foodsTable');
    tbody.innerHTML = '';

    foods.forEach(food => {
        const tr = document.createElement('tr');
        tr.innerHTML = `
            <td>${food.name}</td>
            <td>${food.category}</td>
            <td>${food.calories}</td>
            <td>${food.protein}g</td>
            <td>${food.carbs}g</td>
            <td>${food.fats}g</td>
            <td>${food.serving_size}${food.serving_unit}</td>
            <td>
                ${food.user_id ? `<button onclick="deleteFood('${food.id}')" class="danger">Delete</button>` : '-'}
            </td>
        `;
        tbody.appendChild(tr);
    });
}

function searchFoods() {
    const query = document.getElementById('foodSearch').value.toLowerCase();
    const filtered = allFoods.filter(food =>
        food.name.toLowerCase().includes(query) ||
        food.category.toLowerCase().includes(query)
    );
    displayFoods(filtered);
}

function showAddFoodForm() {
    document.getElementById('addFoodForm').classList.remove('hidden');
}

function hideAddFoodForm() {
    document.getElementById('addFoodForm').classList.add('hidden');
}

async function addFood() {
    const food = {
        name: document.getElementById('foodName').value,
        category: document.getElementById('foodCategory').value,
        calories: parseFloat(document.getElementById('foodCalories').value),
        protein: parseFloat(document.getElementById('foodProtein').value),
        carbs: parseFloat(document.getElementById('foodCarbs').value),
        fats: parseFloat(document.getElementById('foodFats').value),
        serving_size: parseFloat(document.getElementById('foodServingSize').value),
        serving_unit: document.getElementById('foodServingUnit').value
    };

    try {
        const response = await fetch(`${API_BASE}/foods`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(food)
        });

        if (response.ok) {
            alert('Food added successfully!');
            hideAddFoodForm();
            loadFoods();
            // Clear form
            document.getElementById('foodName').value = '';
            document.getElementById('foodCategory').value = '';
            document.getElementById('foodCalories').value = '';
            document.getElementById('foodProtein').value = '';
            document.getElementById('foodCarbs').value = '';
            document.getElementById('foodFats').value = '';
            document.getElementById('foodServingSize').value = '';
            document.getElementById('foodServingUnit').value = '';
        } else {
            const error = await response.text();
            alert('Failed to add food: ' + error);
        }
    } catch (err) {
        alert('Error: ' + err.message);
    }
}

async function deleteFood(id) {
    if (!confirm('Are you sure you want to delete this food?')) return;

    try {
        const response = await fetch(`${API_BASE}/foods/${id}`, {
            method: 'DELETE'
        });

        if (response.ok) {
            alert('Food deleted successfully!');
            loadFoods();
        } else {
            const error = await response.text();
            alert('Failed to delete food: ' + error);
        }
    } catch (err) {
        alert('Error: ' + err.message);
    }
}

// Entries Functions
async function loadEntries() {
    try {
        const response = await fetch(`${API_BASE}/entries`);
        if (response.ok) {
            const entries = await response.json();
            displayEntries(entries);
        }
    } catch (err) {
        alert('Error loading entries: ' + err.message);
    }
}

function displayEntries(entries) {
    const tbody = document.getElementById('entriesTable');
    tbody.innerHTML = '';

    entries.sort((a, b) => new Date(b.eaten_at) - new Date(a.eaten_at));

    entries.forEach(entry => {
        const date = new Date(entry.eaten_at);
        const tr = document.createElement('tr');
        tr.innerHTML = `
            <td>${date.toLocaleDateString()}</td>
            <td>${date.toLocaleTimeString()}</td>
            <td>${entry.meal_type}</td>
            <td>${entry.food_name}</td>
            <td>${entry.quantity}</td>
            <td>${entry.calories.toFixed(0)}</td>
            <td>
                <button onclick="deleteEntry('${entry.id}')" class="danger">Delete</button>
            </td>
        `;
        tbody.appendChild(tr);
    });
}

function populateFoodSelect() {
    const select = document.getElementById('entryFoodId');
    select.innerHTML = '<option value="">Select a food...</option>';

    allFoods.forEach(food => {
        const option = document.createElement('option');
        option.value = food.id;
        option.textContent = `${food.name} (${food.calories} cal)`;
        select.appendChild(option);
    });
}

function showAddEntryForm() {
    const now = new Date();
    const localDateTime = new Date(now.getTime() - now.getTimezoneOffset() * 60000)
        .toISOString()
        .slice(0, 16);
    document.getElementById('entryDateTime').value = localDateTime;
    document.getElementById('addEntryForm').classList.remove('hidden');
}

function hideAddEntryForm() {
    document.getElementById('addEntryForm').classList.add('hidden');
}

async function addEntry() {
    const foodId = document.getElementById('entryFoodId').value;
    const quantity = parseFloat(document.getElementById('entryQuantity').value);
    const mealType = document.getElementById('entryMealType').value;
    const dateTime = document.getElementById('entryDateTime').value;

    if (!foodId) {
        alert('Please select a food');
        return;
    }

    const eatenAt = new Date(dateTime).toISOString();

    const entry = {
        food_id: foodId,
        quantity: quantity,
        meal_type: mealType,
        eaten_at: eatenAt
    };

    try {
        const response = await fetch(`${API_BASE}/entries`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(entry)
        });

        if (response.ok) {
            alert('Entry logged successfully!');
            hideAddEntryForm();
            loadEntries();
            loadDailySummary();
        } else {
            const error = await response.text();
            alert('Failed to log entry: ' + error);
        }
    } catch (err) {
        alert('Error: ' + err.message);
    }
}

async function deleteEntry(id) {
    if (!confirm('Are you sure you want to delete this entry?')) return;

    try {
        const response = await fetch(`${API_BASE}/entries/${id}`, {
            method: 'DELETE'
        });

        if (response.ok) {
            alert('Entry deleted successfully!');
            loadEntries();
            loadDailySummary();
        } else {
            const error = await response.text();
            alert('Failed to delete entry: ' + error);
        }
    } catch (err) {
        alert('Error: ' + err.message);
    }
}

// Dashboard Functions
async function loadDailySummary() {
    const date = document.getElementById('summaryDate').value;

    try {
        const response = await fetch(`${API_BASE}/nutrition/daily/${date}`);
        if (response.ok) {
            const summary = await response.json();

            document.getElementById('totalCalories').textContent = summary.calories.toFixed(0);
            document.getElementById('totalProtein').textContent = summary.protein.toFixed(1) + 'g';
            document.getElementById('totalCarbs').textContent = summary.carbs.toFixed(1) + 'g';
            document.getElementById('totalFats').textContent = summary.fats.toFixed(1) + 'g';

            // Display entries for the day
            const tbody = document.getElementById('dailyEntriesTable');
            tbody.innerHTML = '';

            if (summary.entries && summary.entries.length > 0) {
                summary.entries.forEach(entry => {
                    const time = new Date(entry.eaten_at).toLocaleTimeString();
                    const tr = document.createElement('tr');
                    tr.innerHTML = `
                        <td>${time}</td>
                        <td>${entry.meal_type}</td>
                        <td>${entry.food_name}</td>
                        <td>${entry.quantity}</td>
                        <td>${entry.calories.toFixed(0)}</td>
                        <td>${entry.protein.toFixed(1)}g</td>
                    `;
                    tbody.appendChild(tr);
                });
            } else {
                tbody.innerHTML = '<tr><td colspan="6" style="text-align: center;">No entries for this date</td></tr>';
            }
        }
    } catch (err) {
        alert('Error loading daily summary: ' + err.message);
    }
}

// Goals Functions
async function loadGoals() {
    try {
        const response = await fetch(`${API_BASE}/nutrition/goals`);
        if (response.ok) {
            const goals = await response.json();
            document.getElementById('goalCalories').value = goals.daily_calorie_goal;
            document.getElementById('goalProtein').value = goals.daily_protein_goal;
            document.getElementById('goalCarbs').value = goals.daily_carbs_goal;
            document.getElementById('goalFats').value = goals.daily_fats_goal;
        }
    } catch (err) {
        alert('Error loading goals: ' + err.message);
    }
}

async function updateGoals() {
    const goals = {
        daily_calorie_goal: parseFloat(document.getElementById('goalCalories').value),
        daily_protein_goal: parseFloat(document.getElementById('goalProtein').value),
        daily_carbs_goal: parseFloat(document.getElementById('goalCarbs').value),
        daily_fats_goal: parseFloat(document.getElementById('goalFats').value)
    };

    try {
        const response = await fetch(`${API_BASE}/nutrition/goals`, {
            method: 'PUT',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(goals)
        });

        if (response.ok) {
            alert('Goals updated successfully!');
        } else {
            const error = await response.text();
            alert('Failed to update goals: ' + error);
        }
    } catch (err) {
        alert('Error: ' + err.message);
    }
}
