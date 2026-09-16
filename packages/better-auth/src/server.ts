import { betterAuth } from "better-auth";
import { nextCookies } from "better-auth/next-js";
import { db } from "./db";
import { env } from "./env";

export const auth = betterAuth({
	baseURL: env.BETTER_AUTH_URL,

	database: {
		dialect: db,
		type: "postgres",
		schemaName: "auth",
	},

	emailAndPassword: {
		enabled: true,
	},

	trustedOrigins: env.TRUSTED_DOMAINS,

	plugins: [nextCookies()],
});
