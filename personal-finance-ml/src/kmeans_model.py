import pandas as pd 
from sklearn.cluster import KMeans
from sklearn.preprocessing import StandardScaler

# Load pre-rocessed data : 
df = pd.read_csv('/Users/dhrubajitchakravarty/Documents/Project/PBL/personal-finance-ml/data/processed/synthetic_monthly_expenses.csv')
# Normalize column names (defensive)
df.columns = df.columns.str.strip().str.lower()

# Desired features for clustering
desired_features = [
    "food & drink", "rent", "utilities", "entertainment",
    "travel", "health & fitness", "shopping", "other",
]

# Keep only features that actually exist in the dataset
features = [f for f in desired_features if f in df.columns]

if not features:
    raise ValueError("No valid feature columns found for clustering.")

X = df[features]

# Scale Data : 
scaler = StandardScaler() 
X_scaled = scaler.fit_transform(X)

# KMeans : 
kmeans = KMeans(n_clusters=3, random_state=42)
df['cluster'] = kmeans.fit_predict(X_scaled)

# Save clustered data : 
df.to_csv(
    "/Users/dhrubajitchakravarty/Documents/Project/PBL/personal-finance-ml/outputs/monthly_expenses_clustered_results.csv",
     index = False
)

print("KMeans clustering completed and results saved.")

# -----------------------------
# STEP 3: CLUSTER INTERPRETATION
# -----------------------------

# Calculate cluster-wise average spending
cluster_summary = df.groupby("cluster")[features].mean()

print("\nCluster-wise Average Spending:")
print(cluster_summary)

# Sort clusters by total_expense (low → high)
# Create total_spending manually from all features : 
cluster_summary["total_spending"] = cluster_summary.sum(axis=1)

# Sort clusters based on total_spending :
sorted_clusters = (
    cluster_summary["total_spending"]
    .sort_values()
    .index
    .tolist()
)

# Assign human-readable labels based on spending level
cluster_labels = {}

if len(sorted_clusters) == 3:
    cluster_labels = {
        sorted_clusters[0]: "Saver",
        sorted_clusters[1]: "Balanced",
        sorted_clusters[2]: "High Spender"
    }
else:
    for i, c in enumerate(sorted_clusters):
        cluster_labels[c] = f"Spender_Group_{i+1}"

print("\nCluster Label Mapping:")
for cluster_id, label in cluster_labels.items():
    print(f"Cluster {cluster_id} → {label}")

# Add readable spender type column
df["spender_type"] = df["cluster"].map(cluster_labels)

# Save final interpreted results
df.to_csv(
    "/Users/dhrubajitchakravarty/Documents/Project/PBL/personal-finance-ml/outputs/monthly_expenses_with_spender_type.csv",
    index=False
)

print("\nCluster interpretation completed and labels assigned.")

# -----------------------------
# STEP 4: USER-LEVEL AGGREGATION
# -----------------------------

# Sanity check: required columns
required_cols = ["user_id", "spender_type"]
missing = [c for c in required_cols if c not in df.columns]
if missing:
    raise ValueError(f"Missing required columns for user aggregation: {missing}")

# Count how many times each spender_type appears per user
user_summary = (
    df.groupby(["user_id", "spender_type"])
      .size()
      .reset_index(name="count")
)

# For each user, pick the spender_type with highest count (majority voting)
final_user_profile = (
    user_summary
    .sort_values(["user_id", "count"], ascending=[True, False])
    .drop_duplicates(subset=["user_id"])
    .rename(columns={"spender_type": "overall_spender_type"})
)

print("\nUser-level Spender Type Summary:")
print(final_user_profile.head())

# Save user-level profile result
final_user_profile.to_csv(
    "/Users/dhrubajitchakravarty/Documents/Project/PBL/personal-finance-ml/outputs/user_level_spender_type_summary.csv",
    index=False
)

print("\nUser-level aggregation completed and results saved.") 

# -----------------------------
# STEP 5 : SAVE MODEL FOR BACKEND USE 
# -----------------------------
import joblib 
import os 

MODEL_DIR = '/Users/dhrubajitchakravarty/Documents/Project/PBL/personal-finance-ml/models/'

# Create directory if it doesn't exist : 
os.makedirs(MODEL_DIR, exist_ok=True)

# Save KMeans model : 
joblib.dump(kmeans, os.path.join(MODEL_DIR, 'kmeans_model.pkl'))
# Save Scaler : 
joblib.dump(scaler, os.path.join(MODEL_DIR, 'scaler.pkl'))

print("\nKMeans model and Scaler saved for backend use.")