import Database from "better-sqlite3";
import { Kysely, PostgresDialect, SqliteDialect } from "kysely";
import { Pool } from "pg";

let db: Kysely<any>;

if (process.env.NODE_ENV !== "production") {
	db = new Kysely({
		dialect: new SqliteDialect({
			database: new Database(
				process.env.DATABASE_URL || `${process.cwd()}/sqlite.db`,
			),
		}),
	});
} else {
	db = new Kysely({
		dialect: new PostgresDialect({
			pool: new Pool({
				connectionString:
					process.env.DATABASE_URL ||
					"postgresql://postgres:postgres@localhost:5432/postgres",
			}),
		}),
	});
}

export { db };
