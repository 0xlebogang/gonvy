import {PostgresDialect} from 'kysely'
import {Pool} from 'pg'

export const db = new PostgresDialect({
	pool: new Pool({
		connectionString: process.env.DATABASE_URL as string
	}),
})
