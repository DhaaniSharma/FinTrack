import pandas as pd 

# Load the raw dataset : 
df = pd.read_csv('/Users/dhrubajitchakravarty/Documents/Project/PBL/personal-finance-ml/data/raw/Personal_Finance_Dataset.csv')
df.columns = df.columns.str.strip().str.lower()
# print("Columns after normalization:", df.columns.tolist())

# Ensure user_id exists (single-user dataset fallback)
if "user_id" not in df.columns:
    df["user_id"] = 1

# Basic data cleaning : 
df = df.dropna()
df["category"] = df["category"].str.lower()

# Create month column : 
df["month"] = pd.to_datetime(df["date"]).dt.to_period("M").astype(str)

# Monthly Aggregation : 
monthly = df.pivot_table(
    index = ["user_id", "month"], 
    columns = "category",  
    values = "amount", 
    aggfunc = "sum", 
    fill_value = 0
)

# Total expense : 
monthly["total_expense"] = monthly.sum(axis=1)

monthly.reset_index().to_csv(
    '/Users/dhrubajitchakravarty/Documents/Project/PBL/personal-finance-ml/data/processed/monthly_expenses.csv', 
    index = False
)

print("Monthly ML dataset created successfully.")