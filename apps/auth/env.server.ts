import { createEnv } from "@t3-oss/env-nextjs";
import { globalEnv } from "@workspace/env/server";
import * as z from "zod";

export const env = createEnv({
	extends: [globalEnv],
	server: {
		DATABASE_URL: z.url(),
		PLATFORM_URL: z.url(),
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
	experimental__runtimeEnv: process.env,
});
