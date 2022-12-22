package evidence

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/akralani/jobapplications/database"
)

func getEvidence(evidenceID int) (*Evidence, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	row := database.DbConn.QueryRowContext(ctx, `SELECT 
	evidenceId, 
	date, 
	companyName,
	link,
	jobDescription,
	location,
	jobType,
	field,
	interviewDate,
	interviewDescription,
	accepted 
	FROM evidences 
	WHERE evidenceId = ?`, evidenceID)

	evidence := &Evidence{}
	err := row.Scan(
		&evidence.EvidenceID,
		&evidence.Date,
		&evidence.CompanyName,
		&evidence.Link,
		&evidence.JobDescription,
		&evidence.Location,
		&evidence.JobType,
		&evidence.Field,
		&evidence.InterviewDate,
		&evidence.InterviewDescription,
		&evidence.Accepted,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		log.Println(err)
		return nil, err
	}
	return evidence, nil
}

func GetTopTenEvidences() ([]Evidence, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	results, err := database.DbConn.QueryContext(ctx, `SELECT 
	evidenceId, 
	date, 
	companyName,
	link,
	jobDescription,
	location,
	jobType,
	field,
	interviewDate,
	interviewDescription,
	accepted 
	FROM evidences ORDER BY date DESC LIMIT 10
	`)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer results.Close()
	evidences := make([]Evidence, 0)
	for results.Next() {
		var evidence Evidence
		results.Scan(&evidence.EvidenceID,
			&evidence.Date,
			&evidence.CompanyName,
			&evidence.Link,
			&evidence.JobDescription,
			&evidence.Location,
			&evidence.JobType,
			&evidence.Field,
			&evidence.InterviewDate,
			&evidence.InterviewDescription,
			&evidence.Accepted)

		evidences = append(evidences, evidence)
	}
	return evidences, nil
}

func removeEvidence(evidenceID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := database.DbConn.ExecContext(ctx, `DELETE FROM evidences where evidenceId = ?`, evidenceID)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func getEvidenceList() ([]Evidence, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	results, err := database.DbConn.QueryContext(ctx, `SELECT 
	evidenceId, 
	date, 
	companyName,
	link,
	jobDescription,
	location,
	jobType,
	field,
	interviewDate,
	interviewDescription,
	accepted 
	FROM evidences`)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer results.Close()
	evidences := make([]Evidence, 0)
	for results.Next() {
		var evidence Evidence
		results.Scan(&evidence.EvidenceID,
			&evidence.Date,
			&evidence.CompanyName,
			&evidence.Link,
			&evidence.JobDescription,
			&evidence.Location,
			&evidence.JobType,
			&evidence.Field,
			&evidence.InterviewDate,
			&evidence.InterviewDescription,
			&evidence.Accepted)

		evidences = append(evidences, evidence)
	}
	return evidences, nil
}

func updateEvidence(evidence Evidence) error {
	// if the evidence id is set, update, otherwise add
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if evidence.EvidenceID == nil || *evidence.EvidenceID == 0 {
		return errors.New("evidence has invalid ID")
	}
	_, err := database.DbConn.ExecContext(ctx, `UPDATE evidences SET 
		date=?, 
		companyName=?, 
		link=?, 
		jobDescription=?, 
		location=?, 
		jobType=?, 
		field=?, 
		interviewDate=?, 
		interviewDescription=?, 
		accepted=?
		WHERE evidenceId=?`,
		evidence.Date,
		evidence.CompanyName,
		evidence.Link,
		evidence.JobDescription,
		evidence.Location,
		evidence.JobType,
		evidence.Field,
		evidence.InterviewDate,
		evidence.InterviewDescription,
		evidence.Accepted,
		evidence.EvidenceID)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func insertEvidence(evidence Evidence) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	result, err := database.DbConn.ExecContext(ctx, `INSERT INTO evidences  
	(date, 
	companyName,
	link,
	jobDescription,
	location,
	jobType,
	field,
	interviewDate,
	interviewDescription,
	accepted) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		evidence.Date,
		evidence.CompanyName,
		evidence.Link,
		evidence.JobDescription,
		evidence.Location,
		evidence.JobType,
		evidence.Field,
		evidence.InterviewDate,
		evidence.InterviewDescription,
		evidence.Accepted)
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	return int(insertID), nil
}
