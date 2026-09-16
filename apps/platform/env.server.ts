import { createEnv } from "@t3-oss/env-nextjs";
import { globalEnv } from "@workspace/env/server";

export const env = createEnv({
	extends: [globalEnv],
	server: {},
	experimental__runtimeEnv: process.env,
});
