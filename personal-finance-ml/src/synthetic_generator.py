"""
Docstring for personal-finance-ml.data.synthetic_generator

Purpose : 
- Generate realistic, monthly aggregated expense data for multiple users.
- Used for ML training (KMeans, Linear Regression models). 
- Mimics real human spending behaviou using rules + controlled randomness.
"""
import random 
import pandas as pd 
from datetime import datetime

# Configuration : 
NUM_USERS = 500  # Number of synthetic unique users 
MONTHS = 12   # Months of data per user 

OUTPUT_PATH = (
    '/Users/dhrubajitchakravarty/Documents/Project/PBL/personal-finance-ml/data/processed/synthetic_monthly_expenses.csv'
)

# Spending Categories : 
CATEGORIES = [
    "food & drink", 
    "rent", 
    "utilities", 
    "travel",
    "entertainment", 
    "health & fitness", 
    "shopping", 
    "other", 
    "investment"
]

# Helper Functions : 
def generate_income(): 
    """
    Generate a realistic monthly income based on income bands. 
    """

    band = random.choices(
        ["low", "medium", "high"], 
        weights=[0.4, 0.45, 0.15], 
    )[0]

    if band == "low": 
        return random.randint(20000, 35000)
    elif band == "medium": 
        return random.randint(40000, 80000)
    else: 
        return random.randint(90000, 150000)
    
def generate_monthly_expenses(income): 
    """
    Generate category-wise expenses using real-life financial rules. 
    Percentages are based on common budgeting patterns. 
    """

    rent = income * random.uniform(0.25, 0.35) 
    food = income * random.uniform(0.15, 0.25)
    utilities = income * random.uniform(0.05, 0.08)
    investment = income * random.uniform(0.10, 0.25)

    travel = random.choice([
        0, 
        income * random.uniform(0.03, 0.08)
    ])

    entertainment = income * random.uniform(0.03, 0.07)
    health = random.choice([
        0,
        income * random.uniform(0.02, 0.06)
    ])

    shopping = income * random.uniform(0.04, 0.10)

    allocated = (
        rent + food + utilities + travel + entertainment + health + shopping + investment
    )

    # Remaining goes to "other"
    other = max(income-allocated, 0)

    return {
        "food & drink": int(food), 
        "rent": int(rent), 
        "utilities": int(utilities), 
        "travel": int(travel), 
        "entertainment": int(entertainment), 
        "health & fitness": int(health), 
        "shopping": int(shopping), 
        "other": int(other), 
        "investment": int(investment)
    } 

# Data Generation : 
data = [] 
current_year = datetime.now().year 

for user_idx in range(1, NUM_USERS + 1): 
    user_id = f"U{user_idx:04d}"
    income = generate_income()

    for month_idx in range(1, MONTHS + 1): 
        month = f"{current_year}-{month_idx:02d}"
        expenses = generate_monthly_expenses(income)

        total_expense = sum(
            value for key, value in expenses.items() 
            if key != "investment"
        )

        row = {
            "user_id": user_id, 
            "month": month, 
            "income": income, 
            **expenses, 
            "total_expense": total_expense
        }

        data.append(row)

# Create DataFrame : 
df = pd.DataFrame(data)

# Data Validation and Cleaning : 

# Ensure no negative values : 
for col in CATEGORIES + ["income", "total_expense"]: 
    df[col] = df[col].clip(lower=0)

# Save Dataset : 
df.to_csv(OUTPUT_PATH, index=False)

print(f"Synthetic monthly expense dataset created at: {OUTPUT_PATH}")