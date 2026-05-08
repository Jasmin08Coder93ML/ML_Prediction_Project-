import (
    "github.com/Jasmin08Coder93ML/ML_Prediction_Project/vnutrenniy/lesorub"
)

func main() {
    // Loggeri başlatmak
    l := lesorub.NewLogger("student_activity.csv")
    // Synag hökmünde maglumat ýazmak
    l.LogAction(lesorub.StudentLog{
        StudentID: "Jasmin_001",
        Action:    "app_started",
        Timestamp: time.Now(),
    })
}
