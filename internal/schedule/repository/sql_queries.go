package repository

// This file has for only purpose to serve the repository_test.go file

const (
	createQuery = `INSERT INTO "schedules" (.+) RETURNING`
	getQuery    = `SELECT (.+) FROM "schedules" WHERE "schedules"."id" = \$1 (.+)`
	getAllQuery = `SELECT (.+) FROM "schedules" WHERE (.+)`
	deleteQuery = `UPDATE "schedules" SET "deleted_at"=\$1 WHERE "schedules"."id" = \$2 (.+)`
	// Update use INSERT mainly because of gorm way to work that INSERT instead of UPDATE if record does not exist
	updateQuery  = `INSERT INTO "schedules" (.+) RETURNING`
	getAllByUser = `
	SELECT 
		schedules.* 
	FROM "schedules" 
	LEFT JOIN
		"campus" ON campus.id = schedules.campus
	LEFT JOIN
		"classes" ON classes.id = schedules.class
	LEFT JOIN
		"courses" ON courses.id = schedules.course
	LEFT JOIN
		"schools" ON schools.id = campus.school_id
	LEFT JOIN
		"users" ON users.school_id = schools.id
	WHERE 
		
			(schools.user_id = $1) 
			OR
			(classes.id = users.class_refer AND users.id = $1)
			OR
			(courses.teacher_id =  $1)
		

	`

	getAllByUserUnique = `
	SELECT 
		schedules.* 
	FROM "schedules" 
	LEFT JOIN
		"campus" ON campus.id = schedules.campus
	LEFT JOIN
		"classes" ON classes.id = schedules.class
	LEFT JOIN
		"courses" ON courses.id = schedules.course
	LEFT JOIN
		"schools" ON schools.id = campus.school_id
	LEFT JOIN
		"users" ON users.school_id = schools.id
	WHERE 
		(	
			(schools.user_id = $1) 
			OR
			(classes.id = users.class_refer AND users.id = $1)
			OR
			(courses.teacher_id =  $1)
		)

		
		AND schedules.id = $2 
	LIMIT 1
	`
)
