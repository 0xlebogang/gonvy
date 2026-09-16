import { createEnv } from "@t3-oss/env-core";
import { globalEnv } from "@workspace/env/server";
import * as z from "zod";

export const env = createEnv({
	server: {
		DATABASE_URL: z.url(),
		BETTER_AUTH_URL: z.url(),
		TRUSTED_DOMAINS: z
			.string()
			.min(1, "TRUSTED_DOMAINS is required")
			.transform((value) =>
				value
					.split(",")
					.map((host) => host.trim())
					.filter(Boolean),
			)
			.pipe(z.array(z.url())),
	},
	runtimeEnvStrict: {
		DATABASE_URL: process.env.DATABASE_URL,
		TRUSTED_DOMAINS: process.env.TRUSTED_DOMAINS,
		BETTER_AUTH_URL: process.env.BETTER_AUTH_URL,
	},
	extends: [globalEnv],
});
