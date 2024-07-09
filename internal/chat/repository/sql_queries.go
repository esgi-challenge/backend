package repository

// This file has for only purpose to serve the repository_test.go file

const (
	createQuery = `INSERT INTO "chats" (.+) RETURNING`
	getQuery    = `SELECT (.+) FROM "chats" WHERE "chats"."id" = \$1 (.+)`
	getAllQuery = `SELECT (.+) FROM "chats" WHERE (.+)`
	deleteQuery = `UPDATE "chats" SET "deleted_at"=\$1 WHERE "chats"."id" = \$2 (.+)`
	// Update use INSERT mainly because of gorm way to work that INSERT instead of UPDATE if record does not exist
	updateQuery = `INSERT INTO "chats" (.+) RETURNING`
)