package repository

// This file has for only purpose to serve the repository_test.go file

const (
	createQuery = `INSERT INTO "classs" (.+) RETURNING`
	getQuery    = `SELECT (.+) FROM "classs" WHERE "classs"."id" = \$1 (.+)`
	getAllQuery = `SELECT (.+) FROM "classs" WHERE (.+)`
	deleteQuery = `UPDATE "classs" SET "deleted_at"=\$1 WHERE "classs"."id" = \$2 (.+)`
	// Update use INSERT mainly because of gorm way to work that INSERT instead of UPDATE if record does not exist
	updateQuery = `INSERT INTO "classs" (.+) RETURNING`
)