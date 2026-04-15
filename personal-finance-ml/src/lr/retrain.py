def retrain(user_id=1):
    import pandas as pd
    from sklearn.linear_model import LinearRegression
    import joblib
    from src.lr.db import fetch_expense_data

    df = fetch_expense_data(user_id)
    df["expense_date"] = pd.to_datetime(df["expense_date"])

    df["month"] = df["expense_date"].dt.to_period("M")
    monthly = df.groupby("month")["amount"].sum().reset_index()
    monthly["date"] = monthly["month"].dt.to_timestamp()

    # use RAW month index (NOT normalized)
    monthly["month_num"] = range(len(monthly))

    X = monthly[["month_num"]]
    y = monthly["amount"]

    model = LinearRegression()
    model.fit(X, y)

    joblib.dump((model, None), "models/model.pkl")

    print("Model retrained successfully")
    return "Model retrained successfully"

if __name__ == "__main__":
    retrain(user_id=1)