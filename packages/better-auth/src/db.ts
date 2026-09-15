import {PostgresDialect} from 'kysely'
import {Pool} from 'pg'
import { env } from './env'

export const db = new PostgresDialect({
	pool: new Pool({
		connectionString: env.DATABASE_URL
	}),
})
