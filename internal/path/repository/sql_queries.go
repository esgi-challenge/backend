package repository

// This file has for only purpose to serve the repository_test.go file

const (
	createQuery = `INSERT INTO "paths" (.+) RETURNING`
	getQuery    = `SELECT (.+) FROM "paths" WHERE "paths"."id" = \$1 (.+)`
	getAllQuery = `SELECT (.+) FROM "paths" WHERE (.+)`
	deleteQuery = `UPDATE "paths" SET "deleted_at"=\$1 WHERE "paths"."id" = \$2 (.+)`
	// Update use INSERT mainly because of gorm way to work that INSERT instead of UPDATE if record does not exist
	updateQuery = `INSERT INTO "paths" (.+) RETURNING`
)
