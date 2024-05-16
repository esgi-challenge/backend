package repository

const(
  createQuery = `INSERT INTO "examples" (.+) RETURNING`
  getQuery = `SELECT (.+) FROM "examples" WHERE "examples"."id" = \$1 (.+)`
  getAllQuery = `SELECT (.+) FROM "examples" WHERE (.+)`
  deleteQuery = `UPDATE "examples" SET "deleted_at"=\$1 WHERE "examples"."id" = \$2 (.+)`
  // updateQuery = `UPDATE "examples" SET (.+)`
  updateQuery = `INSERT INTO "examples" (.+) RETURNING`
)
