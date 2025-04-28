package background

import (
	"VoteGolang/internal/database"
	"log"
	"time"
)

func StartElectionAutoCloser() {
	ticker := time.NewTicker(1 * time.Hour)
	go func() {
		for {
			select {
			case <-ticker.C:
				closeExpiredElections()
			}
		}
	}()
}

func closeExpiredElections() {
	_, err := database.DB.Exec(`
		UPDATE elections
		SET is_closed = TRUE
		WHERE election_date < NOW()
		AND is_closed = FALSE
	`)
	if err != nil {
		log.Println("❌ Failed to auto-close elections:", err)
	} else {
		log.Println("✅ Auto-closed expired elections.")
	}
}
