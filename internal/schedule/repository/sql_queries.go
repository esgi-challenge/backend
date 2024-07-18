package repository

const (
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
