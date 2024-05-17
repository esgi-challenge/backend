package repository

// This file has for only purpose to serve the repository_test.go file

const(
  createQuery = `INSERT INTO "examples" (.+) RETURNING`
  getQuery = `SELECT (.+) FROM "examples" WHERE "examples"."id" = \$1 (.+)`
  getAllQuery = `SELECT (.+) FROM "examples" WHERE (.+)`
  deleteQuery = `UPDATE "examples" SET "deleted_at"=\$1 WHERE "examples"."id" = \$2 (.+)`
  // Update use INSERT mainly because of gorm way to work that INSERT instead of UPDATE if record does not exist
  updateQuery = `INSERT INTO "examples" (.+) RETURNING`
)
