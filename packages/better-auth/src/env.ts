import { createEnv } from "@t3-oss/env-core";
import { globalEnv } from "@workspace/env/server";
import * as z from "zod";

export const env = createEnv({
	server: {
		DATABASE_URL: z.url(),
	},
	runtimeEnvStrict: {
		DATABASE_URL: process.env.DATABASE_URL,
	},
	extends: [globalEnv],
});
