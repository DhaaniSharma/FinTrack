import plotly.graph_objects as go
from src.lr.predict import predict_range

def generate_plot_html(user_id, start_month, end_month, model_path="models/model.pkl"):

    # USE BACKEND FUNCTION (IMPORTANT)
    data = predict_range(model_path, start_month, end_month, user_id)

    if "error" in data:
        return f"<h2>{data['error']}</h2>"

    history_dates = data["history_dates"]
    actual = data["history_actual"]
    lr = data["lr_line"]

    future_dates = data.get("future_dates", [])
    future = data.get("future_predicted", [])

    fig = go.Figure()

    # ACTUAL
    fig.add_trace(go.Scatter(
        x=history_dates,
        y=actual,
        mode='lines+markers',
        name='Actual',
        line=dict(color='#00E5FF', width=3, dash='dot')
    ))

    # LR
    fig.add_trace(go.Scatter(
        x=history_dates,
        y=lr,
        mode='lines',
        name='Linear Regression',
        line=dict(color='#FF4D4D', width=4)
    ))

    # FUTURE (ONLY IF EXISTS)
    if future_dates:
        fig.add_trace(go.Scatter(
            x=future_dates,
            y=future,
            mode='lines',
            name='Prediction',
            line=dict(color='#FFA500', width=4)
        ))

    fig.update_layout(
        template="plotly_dark",
        title="Expense Prediction",
        xaxis_title="Date",
        yaxis_title="Amount",
        legend=dict(x=0.01, y=0.99),
        margin=dict(l=40, r=40, t=60, b=40),
        xaxis=dict(type="date", tickformat="%b %Y")
    )

    return fig.to_html(full_html=False)