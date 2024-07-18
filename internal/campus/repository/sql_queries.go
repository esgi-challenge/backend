package repository

// This file has for only purpose to serve the repository_test.go file

const (
	createQuery = `INSERT INTO "campus" (.+) RETURNING`
	getQuery    = `SELECT (.+) FROM "campus" (.+)`
	getAllQuery = `SELECT (.+) FROM "campus" WHERE (.+)`
	deleteQuery = `UPDATE "campus" SET "deleted_at"=\$1 WHERE "campus"."id" = \$2 (.+)`
	// Update use INSERT mainly because of gorm way to work that INSERT instead of UPDATE if record does not exist
	updateQuery = `INSERT INTO "campus" (.+) RETURNING`
)
