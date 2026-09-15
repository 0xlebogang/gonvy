import { createEnv } from "@t3-oss/env-core";
import * as z from "zod";

export const globalEnv = createEnv({
	server: {
		PORT: z.string(),
		DATABASE_URL: z.url(),
	},

	runtimeEnvStrict: {
		DATABASE_URL: process.env.DATABASE_URL,
		PORT: process.env.PORT,
	},
	emptyStringAsUndefined: true,
});
