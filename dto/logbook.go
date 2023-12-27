package dto

type Logbook struct {
	Data Weekly      `json:"data"`
	Meta interface{} `json:"meta"`
}

type Weekly struct {
	ID             string        `json:"id"`
	ActivityID     string        `json:"activity_id"`
	UuidAkun       string        `json:"uuid_akun"`
	IDRegPenawaran int           `json:"id_reg_penawaran"`
	StartDate      string        `json:"start_date"`
	EndDate        string        `json:"end_date"`
	Counter        uint          `json:"counter"`
	Status         string        `json:"status"`
	LearnedWeekly  string        `json:"learned_weekly"`
	NoteFromMentor string        `json:"note_from_mentor"`
	ReviewedBy     string        `json:"reviewed_by"`
	ReviewedName   string        `json:"reviewed_name"`
	DailyReport    []DailyReport `json:"daily_report"`
	MentorNotes    []interface{} `json:"mentor_notes"`
	UpdatedAt      string        `json:"updated_at"`
}

type DailyReport struct {
	ID             string `json:"id"`
	WeekyReportID  string `json:"weekly_report_id"`
	UuidAkun       string `json:"uuid_akun"`
	IDRegPenawaran int    `json:"id_reg_penawaran"`
	ReportDate     string `json:"report_date"`
	Status         string `json:"status"`
	FeelingLevel   uint8  `json:"feeling_level"`
	Report         string `json:"report"`
	Counter        uint   `json:"counter"`
	CreatedBy      string `json:"created_by"`
	UpdatedBy      string `json:"updated_by"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}
