// CSV faýlyny açýarys (eger ýok bolsa döredýär, bar bolsa soňuna goşýar)
file, err := os.OpenFile(l.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
if err != nil {
    return err
}
defer file.Close()

writer := csv.NewWriter(file)
defer writer.Flush()

// Okuwçy maglumatlaryny setir (string slice) görnüşine öwürýäris
row := []string{
    log.StudentID,
    log.Action,
    strconv.Itoa(log.Score),
    strconv.Itoa(log.Duration),
    log.Timestamp.Format(time.RFC3339),
}

// Faýla ýazýarys
return writer.Write(row)
