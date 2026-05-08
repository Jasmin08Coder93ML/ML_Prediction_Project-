package logger

import (
    "fmt"
    "time"
    // Bu ýoly öz GitHub adyňa görä düzet:
    logger "github.com/Jasmin08Coder93ML/ML_Prediction_Project/vnutrenniy/lesorub" 
)
import (
    "fmt"
    // ... başga paketleriň ...
    "github.com/Jasmin08Coder93ML/ML_Prediction_Project/vnutrenniy/logger" // Seniň loggeriň ýoly
)
	"encoding/csv"
	"os"
	"strconv"
	"sync"
	"time"
)

type StudentLog struct {
	StudentID string
	Action    string
	Score     int
	Duration  int
	Timestamp time.Time
}

type Logger struct {
	FilePath string
	mu       sync.Mutex
}
func main(myLogger := ...) {
    // Loggeri başladýarys we faýl adyny berýäris
    myLogger := logger.NewLogger("students_activity.csv")
    
    fmt.Println("Backend ulgamy we Logger işläp başlady...")
    // ... galan kodlaryň ...
}
// Mysal üçin, bir test tamamlananda:
err := myLogger.LogAction(logger.StudentLog{
    StudentID: "Okuwcy_001",
    Action:    "test_completed",
    Score:     95,
    Duration:  120, // sekuntda
    Timestamp: time.Now(),
})

if err != nil {
    fmt.Println("Log ýazylanda ýalňyşlyk:", err)
}
func main() {
    // 1. Loggeri başladýarys (faýlyň adyny görkezýäris)
    appLogger := logger.NewLogger("student_activity.csv")

    // 2. Synag hökmünde maglumat ýazýarys
    err := appLogger.LogAction(logger.StudentLog{
        StudentID: "Jasmin_001",
        Action:    "login_success",
        Score:     100,
        Duration:  0,
        Timestamp: time.Now(),
    })

    if err != nil {
        fmt.Println("Ýalňyşlyk ýüze çykdy:", err)
        return
    }

    fmt.Println("Logger üstünlikli işledi we maglumat ýazyldy!")
}
func NewLogger(path string) *Logger {
	return &Logger{FilePath: path}
}

func (l *Logger) LogAction(log StudentLog) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	// 1. Faýlyň barlygyny barlaýarys (Header goşmak üçin)
	_, err := os.Stat(l.FilePath)
	isNew := os.IsNotExist(err)

	file, err := os.OpenFile(l.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// 2. EGER FAÝL TÄZE BOLSA, SÜTÜN ATLARYNY ÝAZÝARYS
	if isNew {
		header := []string{"StudentID", "Action", "Score", "Duration", "Timestamp"}
		writer.Write(header)
	}

	// 3. Maglumatlary ýazýarys
	row := []string{
		log.StudentID,
		log.Action,
		strconv.Itoa(log.Score),
		strconv.Itoa(log.Duration),
		log.Timestamp.Format(time.RFC3339),
	}

	return writer.Write(row)
}
