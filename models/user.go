package models

type User struct {
	Id              int `json:"id" db:"id"`
	Score           int `json:"score" db:"score"`
	TranslatedPages int `json:"translated_pages" db:"translated_pages"`
	EditedPages     int `json:"edited_pages" db:"edited_pages"`
	CheckedPages    int `json:"checked_pages" db:"checked_pages"`
	CleanedPages    int `json:"cleaned_pages" db:"cleaned_pages"`
	TypedPages      int `json:"typed_pages" db:"typed_pages"`
}
