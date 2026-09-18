import { betterAuth } from "better-auth";
import { openAPI } from "better-auth/plugins";
import { db } from "./db";

const isDev = process.env.NODE_ENV !== "production";

export const auth = betterAuth({
	baseURL: process.env.BETTER_AUTH_URL,
	secret: process.env.BETTER_AUTH_SECRET,
	database: {
		db: db,
		type: isDev ? "sqlite" : "postgres",
		...(isDev ? {} : { schemaName: "auth" }),
	},
	emailAndPassword: {
		enabled: true,
	},
	plugins: [openAPI()],
});
