package models

import (
	"database/sql"
	"time"
)

type ReviewLog struct {
	ID int						`json:"id"`
	ReviewScheduleID string		`json:"review_schedule_id"`
	Rating int 					`json:"rating"`
	ReviewDate time.Time		`json:"review_date"`
	ElapsedDays int				`json:"elapsed_days"`
	ScheduledDays int			`json:"scheduled_days"`
	State int 					`json:"state"`
}

type ReviewLogStore struct {
	db *sql.DB
}

func NewReviewLogStore(db *sql.DB) *ReviewLogStore {
	return &ReviewLogStore{db: db}
}

func (s *ReviewLogStore) CreateReviewLog(log *ReviewLog) error {
	query := `
		INSERT INTO review_logs 
		(review_schedule_id, rating, review_date, elapsed_days, scheduled_days, state) 
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	err := s.db.QueryRow(
		query,
		log.ReviewScheduleID,
		log.Rating,
		log.ReviewDate,
		log.ElapsedDays,
		log.ScheduledDays,
		log.State,
	).Scan(&log.ID)

	return err
}

func (s *ReviewLogStore) GetReviewLogsByUserID(userID, limit, offset int) ([]ReviewLog, error) {
    query := `
        SELECT r.id, r.review_schedule_id, r.rating, r.review_date, 
                r.elapsed_days, r.scheduled_days, r.state
        FROM review_logs r
        JOIN review_schedules sched ON r.review_schedule_id = sched.id
        JOIN submissions sub ON sched.submission_id = sub.id
        WHERE sub.user_id = $1
        ORDER BY r.review_date DESC
        LIMIT $2 OFFSET $3
    `
    
    rows, err := s.db.Query(query, userID, limit, offset)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var logs []ReviewLog
    for rows.Next() {
        var log ReviewLog
        if err := rows.Scan(
            &log.ID, 
            &log.ReviewScheduleID, 
            &log.Rating, 
            &log.ReviewDate, 
            &log.ElapsedDays, 
            &log.ScheduledDays, 
            &log.State,
        ); err != nil {
            return nil, err
        }
        logs = append(logs, log)
    }
    
    if err := rows.Err(); err != nil {
        return nil, err
    }
    
    return logs, nil
}

func (s *ReviewLogStore) GetReviewLogsCountByUserID(userID int) (int, error) {
    query := `
        SELECT COUNT(*)
        FROM review_logs r
        JOIN review_schedules sched ON r.review_schedule_id = sched.id
        JOIN submissions sub ON sched.submission_id = sub.id
        WHERE sub.user_id = $1
    `
    
    var count int
    err := s.db.QueryRow(query, userID).Scan(&count)
    if err != nil {
        return 0, err
    }
    
    return count, nil
}