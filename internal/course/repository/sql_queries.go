package repository

// This file has for only purpose to serve the repository_test.go file

const (
	createQuery = `INSERT INTO "courses" (.+) RETURNING`
	getQuery    = `SELECT (.+) FROM "courses" WHERE "courses"."id" = \$1 (.+)`
	getAllQuery = `SELECT (.+) FROM "courses" WHERE (.+)`
	deleteQuery = `UPDATE "courses" SET "deleted_at"=\$1 WHERE "courses"."id" = \$2 (.+)`
	// Update use INSERT mainly because of gorm way to work that INSERT instead of UPDATE if record does not exist
	updateQuery = `INSERT INTO "courses" (.+) RETURNING`
)
