package repository

// This file has for only purpose to serve the repository_test.go file

const (
	createQuery = `INSERT INTO "campuss" (.+) RETURNING`
	getQuery    = `SELECT (.+) FROM "campuss" WHERE "campuss"."id" = \$1 (.+)`
	getAllQuery = `SELECT (.+) FROM "campuss" WHERE (.+)`
	deleteQuery = `UPDATE "campuss" SET "deleted_at"=\$1 WHERE "campuss"."id" = \$2 (.+)`
	// Update use INSERT mainly because of gorm way to work that INSERT instead of UPDATE if record does not exist
	updateQuery = `INSERT INTO "campuss" (.+) RETURNING`
)
