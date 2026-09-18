import { createNextjsEnv } from "@workspace/env";
import { globalEnv } from "@workspace/env/server";
import * as z from "zod/v3";

export const env = createNextjsEnv({
	extends: [globalEnv],
	server: {
		PLATFORM_URL: z.string().url(),
		BETTER_AUTH_URL: z.string().url(),
		BETTER_AUTH_SECRET: z
			.string()
			.min(32, "Must be at least 32 characters long"),
	},
	experimental__runtimeEnv: process.env,
});
