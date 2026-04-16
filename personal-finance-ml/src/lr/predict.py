def predict_range(model_path, start_month, end_month, user_id=1):
    import pandas as pd
    import joblib
    from src.lr.db import fetch_expense_data

    # load model
    model, _ = joblib.load(model_path)

    # fetch data
    df = fetch_expense_data(user_id)
    df["expense_date"] = pd.to_datetime(df["expense_date"])

    if df.empty:
        return {"error": "No data found"}

    # monthly aggregation
    df["month"] = df["expense_date"].dt.to_period("M")
    monthly = df.groupby("month")["amount"].sum().reset_index()
    monthly["date"] = monthly["month"].dt.to_timestamp()
    total_months = len(monthly)

    # create feature 
    monthly["month_num"] = range(len(monthly))

    # LR prediction on history
    lr_preds = model.predict(monthly[["month_num"]])

    # user wants range like 1 → 6 (recent months)
    history_months_needed = min(end_month, total_months)

    monthly = monthly.tail(history_months_needed)
    lr_preds = lr_preds[-history_months_needed:]

    # CHECK IF FUTURE IS NEEDED
    last_available_date = monthly["date"].iloc[-1]

    need_future = end_month > total_months

    # FUTURE PREDICTION (ONLY IF NEEDED)
    if need_future:
        future_steps = max(0, end_month - total_months)

        future_dates = pd.date_range(
            start=last_available_date + pd.DateOffset(months=1),
            periods=future_steps,
            freq="MS"
        )

        future_months = list(range(total_months, total_months + future_steps))

        future_df = pd.DataFrame({"month_num": future_months})

        future_preds = model.predict(future_df[["month_num"]])

        # CONNECT WITH LR LINE (NO GAP)
        offset = lr_preds[-1] - future_preds[0]
        future_preds = future_preds + offset

    else:
        future_dates = []
        future_preds = []

    # FINAL RESPONSE
    return {
        "history_dates": monthly["date"].dt.strftime("%Y-%m-%d").tolist(),
        "history_actual": monthly["amount"].tolist(),
        "lr_line": [float(x) for x in lr_preds],
        "future_dates": future_dates.strftime("%Y-%m-%d").tolist() if need_future else [],
        "future_predicted": [float(x) for x in future_preds] if need_future else [],
        "last_date": (
            future_dates[-1].strftime("%Y-%m-%d")
            if need_future else monthly["date"].iloc[-1].strftime("%Y-%m-%d")
        ),
        "last_expense": (
            float(future_preds[-1])
            if need_future else float(monthly["amount"].iloc[-1])
        )
    }