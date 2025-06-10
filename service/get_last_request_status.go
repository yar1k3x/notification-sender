package service

import (
	"NotificationSender/db"
	"log"
)

func GetLastStatusChange(requestID int32) (oldStatus string, newStatus string, err error) {
	query := `
		SELECT 
			rs_old.status_name AS old_status_name,
			rs_new.status_name AS new_status_name
		FROM request_status_journal rj
		JOIN request_status rs_old ON rj.old_status = rs_old.id
		JOIN request_status rs_new ON rj.new_status = rs_new.id
		WHERE rj.request_id = ?
		ORDER BY rj.created_at DESC
		LIMIT 1
	`

	err = db.DB.QueryRow(query, requestID).Scan(&oldStatus, &newStatus)
	if err != nil {
		log.Printf("Ошибка при получении названий статусов для request_id %d: %v", requestID, err)
		return "", "", err
	}

	return oldStatus, newStatus, nil
}
