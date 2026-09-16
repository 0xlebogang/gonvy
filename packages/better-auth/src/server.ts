import { betterAuth } from "better-auth";
import { nextCookies } from "better-auth/next-js";
import { db } from "./db";

export const auth = betterAuth({
	database: {
		dialect: db,
		type: "postgres",
		schemaName: "auth",
	},

	emailAndPassword: {
		enabled: true,
	},

	plugins: [nextCookies()],
});
