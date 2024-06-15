package repository

// This file has for only purpose to serve the repository_test.go file

const (
	createQuery = `INSERT INTO "schedules" (.+) RETURNING`
	getQuery    = `SELECT (.+) FROM "schedules" WHERE "schedules"."id" = \$1 (.+)`
	getAllQuery = `SELECT (.+) FROM "schedules" WHERE (.+)`
	deleteQuery = `UPDATE "schedules" SET "deleted_at"=\$1 WHERE "schedules"."id" = \$2 (.+)`
	// Update use INSERT mainly because of gorm way to work that INSERT instead of UPDATE if record does not exist
	updateQuery = `INSERT INTO "schedules" (.+) RETURNING`
)