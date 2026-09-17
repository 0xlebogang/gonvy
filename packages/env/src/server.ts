import { createEnv } from "@t3-oss/env-core";
import * as z from "zod";

export const globalEnv = createEnv({
	server: {
		PORT: z.string().optional(),
		NODE_ENV: z.enum(
			["development", "testing", "production"],
			"NODE_ENV should be set to 'development', 'testing' or 'production'",
		),
	},
	runtimeEnv: {
		NODE_ENV: process.env.NODE_ENV,
		PORT: process.env.PORT,
	},
	emptyStringAsUndefined: true,
});
