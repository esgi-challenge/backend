package repository

// This file has for only purpose to serve the repository_test.go file

const (
	createQuery = `INSERT INTO "notes" (.+) RETURNING`
	getQuery    = `SELECT (.+) FROM "notes" WHERE "notes"."id" = \$1 (.+)`
	getAllQuery = `SELECT (.+) FROM "notes" WHERE (.+)`
	deleteQuery = `UPDATE "notes" SET "deleted_at"=\$1 WHERE "notes"."id" = \$2 (.+)`
	// Update use INSERT mainly because of gorm way to work that INSERT instead of UPDATE if record does not exist
	updateQuery = `INSERT INTO "notes" (.+) RETURNING`
)
