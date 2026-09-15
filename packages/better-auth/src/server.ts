import { betterAuth } from "better-auth";
import { db } from "./db";

export const auth = betterAuth({
  database: {
    dialect: db,
    type: "postgres",
    schemaName: "auth",
  },

	emailAndPassword: {
		enabled: true
	}
})
