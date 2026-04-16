package auth

import (
	"encoding/json"
	"net/http"
)

// 🔹 Redirect user to Google
func GoogleLoginHandler(w http.ResponseWriter, r *http.Request) {
	config := GetGoogleOAuthConfig()

	url := config.AuthCodeURL("state-token")
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// 🔹 Handle Google callback
func GoogleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	config := GetGoogleOAuthConfig()

	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Code not found", http.StatusBadRequest)
		return
	}

	// 🔁 Exchange code for token
	token, err := config.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, "Failed to exchange token", http.StatusInternalServerError)
		return
	}

	// 👤 Get user info from Google
	client := config.Client(r.Context(), token)

	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		http.Error(w, "Failed to get user info", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		http.Error(w, "Failed to decode user info", http.StatusInternalServerError)
		return
	}

	// 📧 Extract email safely
	email, ok := userInfo["email"].(string)
	if !ok {
		http.Error(w, "Email not found", http.StatusInternalServerError)
		return
	}

	println("Google user:", email)

	// 👉 TODO: Check DB or create user
	// For now using dummy userID = 1
	jwtToken, err := GenerateJWT(1)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"message": "Google login successful",
		"token":   jwtToken,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}