package main

import (github.com/Jasmin08Coder93ML/ML_Prediction_Project/vnutrenniy/loggerlogger.NewLogger("logs.csv") 
	"encoding/csv" // CSV formatynda ýazmak üçin
	"os"           // Faýl ulgamy bilen işlemek üçin
	"strconv"      // Sanlary tekst öwürmek üçin
	"time"         // Wagty hasaba almak üçin
)

// StudentLog - Okuwçynyň hereketini aňladýan struktura
type StudentLog struct {
	StudentID string
	Action    string
	Score     int
	Duration  int
	Timestamp time.Time
}

// Logger - Faýl ýoluny saklaýan struktura
type Logger struct {
	FilePath string
}

// NewLogger - Täze logger döretmek üçin kömekçi funksiýa
func NewLogger(path string) *Logger {
	return &Logger{FilePath: path}
}

// LogAction - Maglumaty CSV faýlyna ýazýan esasy funksiýa
func (l *Logger) LogAction(log StudentLog) error {
	// Faýly açýarys: goşmak (APPEND), ýok bolsa döretmek (CREATE), diňe ýazmak (WRONLY)
	// 0644 - faýlyň rugsat hukuklary (read/write)
	file, err := os.OpenFile(l.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	// Funksiýa tamamlananda faýly ýapmagy üpjün edýäris
	defer file.Close()

	writer := csv.NewWriter(file)
	// Bu örän möhüm: buffer-däki maglumatlary hakyky faýla boşadýar
	defer writer.Flush()

	// Maglumatlary sütünlere (string slice) öwürýäris
	row := []string{
		log.StudentID,
		log.Action,
		strconv.Itoa(log.Score),    // int -> string
		strconv.Itoa(log.Duration), // int -> string
		log.Timestamp.Format(time.RFC3339), // Wagty standart formatda ýazýarys
	}

	// Setiri CSV faýlyna ýazýarys
	return writer.Write(row)
}

// Taslamany barlap görmek üçin main funksiýasy (Nusga)
func main() {
	// Loggeri döredýäris. Faýl ady: students_logs.csv
	myLogger := NewLogger("students_logs.csv")

	// Barlag üçin bir maglumat döredeliň
	testData := StudentLog{
		StudentID: "User_123",
		Action:    "test_completed",
		Score:     85,
		Duration:  300,
		Timestamp: time.Now(),
	}

	// CSV-ä ýazmaga synanyşýarys
	err := myLogger.LogAction(testData)
	if err != nil {
		// Ýalňyşlyk bolsa terminalda görkezýär
		panic(err)
	}
}
