package models

type MLOverride struct {
	ID                int    `json:"id"`
	UserID            int    `json:"user_id"`
	RawText           string `json:"raw_text"`
	CorrectedCategory string `json:"corrected_category"`
}
