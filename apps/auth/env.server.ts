import { createEnv } from "@t3-oss/env-nextjs";
import { globalEnv } from "@workspace/env/server";
import * as z from "zod";

export const env = createEnv({
	extends: [globalEnv],
	server: {
		DATABASE_URL: z.url(),
	},
	experimental__runtimeEnv: process.env,
});
