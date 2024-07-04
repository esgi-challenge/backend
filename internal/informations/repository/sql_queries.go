package repository

// This file has for only purpose to serve the repository_test.go file

const (
	createQuery = `INSERT INTO "informationss" (.+) RETURNING`
	getQuery    = `SELECT (.+) FROM "informationss" WHERE "informationss"."id" = \$1 (.+)`
	getAllQuery = `SELECT (.+) FROM "informationss" WHERE (.+)`
	deleteQuery = `UPDATE "informationss" SET "deleted_at"=\$1 WHERE "informationss"."id" = \$2 (.+)`
	// Update use INSERT mainly because of gorm way to work that INSERT instead of UPDATE if record does not exist
	updateQuery = `INSERT INTO "informationss" (.+) RETURNING`
)
