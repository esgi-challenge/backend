package repository

// This file has for only purpose to serve the repository_test.go file

const (
	createQuery = `INSERT INTO "documents" (.+) RETURNING`
	getQuery    = `SELECT (.+) FROM "documents" WHERE "documents"."id" = \$1 (.+)`
	getAllQuery = `SELECT (.+) FROM "documents" WHERE (.+)`
	deleteQuery = `UPDATE "documents" SET "deleted_at"=\$1 WHERE "documents"."id" = \$2 (.+)`
	// Update use INSERT mainly because of gorm way to work that INSERT instead of UPDATE if record does not exist
	updateQuery = `INSERT INTO "documents" (.+) RETURNING`
)
