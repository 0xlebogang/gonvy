import { createEnv } from "@t3-oss/env-nextjs";
import { globalEnv } from "@workspace/env/server";

export const env = createEnv({
	server: {},
	runtimeEnv: process.env,
	extends: [globalEnv],
});
