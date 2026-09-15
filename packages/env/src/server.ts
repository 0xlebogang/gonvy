import { createEnv } from "@t3-oss/env-core";
import * as z from "zod";

export const globalEnv = createEnv({
	server: {
		PORT: z.string(),
		DATABASE_URL: z.url(),
	},

	runtimeEnv: process.env,
	emptyStringAsUndefined: true,
});
