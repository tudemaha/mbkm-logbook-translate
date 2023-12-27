package controller

import (
	"log"

	"github.com/tudemaha/mbkm-logbook-translate/internal/translation/dto"
	"github.com/tudemaha/mbkm-logbook-translate/pkg"
)

func GetTranslationPerWeek(sourceLang, targetLang string, week int) dto.Result {
	var result dto.Result
	var dailyReports []dto.Daily

	weeklyLogbook, err := pkg.GetLogbook(week)
	if err != nil {
		log.Fatalf("%v", err)
	}
	result.StartDate = weeklyLogbook.Data.StartDate
	result.EndDate = weeklyLogbook.Data.EndDate

	var reports []string
	reports = append(reports, weeklyLogbook.Data.LearnedWeekly)
	for _, daily := range weeklyLogbook.Data.DailyReport {
		dailyReports = append(dailyReports, dto.Daily{CreatedAt: daily.CreatedAt})
		reports = append(reports, daily.Report)
	}

	translatedReports, err := pkg.GetTranslatedReport(sourceLang, targetLang, reports)
	if err != nil {
		log.Fatalf("%v", err)
	}
	for i, transltedReport := range translatedReports {
		if i == 0 {
			result.LearnedWeekly = transltedReport
			continue
		}

		dailyReports[i-1].Report = transltedReport
	}

	result.DailyReport = dailyReports
	return result
}
