package dto

type Result struct {
	StartDate     string  `json:"start_date"`
	EndDate       string  `json:"end_date"`
	LearnedWeekly string  `json:"learned_weekly"`
	DailyReport   []Daily `json:"daily_report"`
}

type Daily struct {
	CreatedAt string `json:"created_at"`
	Report    string `json:"report"`
}
