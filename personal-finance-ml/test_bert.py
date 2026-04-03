from transformers import pipeline

classifier = pipeline("zero-shot-classification", model="typeform/distilbert-base-uncased-mnli")

labels_underscore = [
    "food_and_drink",
    "rent",
    "utilities",
    "entertainment",
    "travel",
    "health_and_fitness",
    "shopping",
    "other"
]

labels_clean = [
    "food and drink",
    "rent",
    "utilities",
    "entertainment",
    "travel",
    "health and fitness",
    "shopping",
    "other"
]

text = "coffee from starbucks"

print(f"\n--- Testing Underscore Labels ---")
print(classifier(text, labels_underscore))

print(f"\n--- Testing Clean Labels ---")
print(classifier(text, labels_clean))
