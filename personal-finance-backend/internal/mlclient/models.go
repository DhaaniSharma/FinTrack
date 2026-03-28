package mlclient

type ExpenseInput struct {
	FoodAndDrink     float64 `json:"food_and_drink"`
	Rent             float64 `json:"rent"`
	Utilities        float64 `json:"utilities"`
	Entertainment    float64 `json:"entertainment"`
	Travel           float64 `json:"travel"`
	HealthAndFitness float64 `json:"health_and_fitness"`
	Shopping         float64 `json:"shopping"`
	Other            float64 `json:"other"`
}

type PredictionData struct {
	Cluster     int    `json:"cluster"`
	SpenderType string `json:"spender_type"`
}

type MLResponse struct {
	Status string          `json:"status"`
	Data   *PredictionData `json:"data"`
	Error  any             `json:"error"`
}
