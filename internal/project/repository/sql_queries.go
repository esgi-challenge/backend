package repository

// This file has for only purpose to serve the repository_test.go file

const (
	createQuery = `INSERT INTO "projects" (.+) RETURNING`
	getQuery    = `SELECT (.+) FROM "projects" WHERE "projects"."id" = \$1 (.+)`
	getAllQuery = `SELECT (.+) FROM "projects" WHERE (.+)`
	deleteQuery = `UPDATE "projects" SET "deleted_at"=\$1 WHERE "projects"."id" = \$2 (.+)`
	// Update use INSERT mainly because of gorm way to work that INSERT instead of UPDATE if record does not exist
	updateQuery           = `INSERT INTO "projects" (.+) RETURNING`
	checkIfExistsMultiple = `SELECT * FROM "projects" 
	LEFT JOIN classes ON projects.class_id = classes.id 
	LEFT JOIN paths ON classes.path_id = paths.id 
	LEFT JOIN schools ON paths.school_id = schools.id 
	LEFT JOIN users ON users.school_id = schools.id  
	WHERE 
			(users.school_id = $1 
		OR 
			schools.user_id = $1 
		OR
			users.class_refer = classes.id
		)
	AND projects.deleted_at IS NULL`

	checkIfExists     = `SELECT projects.* FROM "projects" LEFT JOIN classes ON projects.class_id = classes.id LEFT JOIN paths ON classes.path_id = paths.id LEFT JOIN schools ON paths.school_id = schools.id LEFT JOIN users ON users.school_id = schools.id  WHERE (users.school_id = $1 OR schools.user_id = $1) AND projects.id = $2 AND projects.deleted_at IS NULL`
	checkIfJoined     = `SELECT projects.* FROM "project_students" LEFT JOIN projects ON project_students.project_id = projects.id WHERE project_students.student_id = $1 AND projects.id = $2 AND project_students.deleted_at IS NULL`
	checkIfJoinedUniq = `SELECT project_students.id FROM "project_students" LEFT JOIN projects ON project_students.project_id = projects.id WHERE project_students.student_id = $1 AND projects.id = $2 AND project_students.deleted_at IS NULL`
)
