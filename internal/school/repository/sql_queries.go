package repository

// This file has for only purpose to serve the repository_test.go file

const (
	createQuery = `INSERT INTO "schools" (.+) RETURNING`
	getQuery    = `SELECT (.+) FROM "schools" WHERE "schools"."id" = \$1 (.+)`
	getAllQuery = `SELECT (.+) FROM "schools" WHERE (.+)`
	deleteQuery = `UPDATE "schools" SET "deleted_at"=\$1 WHERE "schools"."id" = \$2 (.+)`
	// Update use INSERT mainly because of gorm way to work that INSERT instead of UPDATE if record does not exist
)
