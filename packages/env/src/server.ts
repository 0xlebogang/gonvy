import { createEnv } from "@t3-oss/env-core";
import * as z from "zod";

export const globalEnv = createEnv({
	server: {
		PORT: z.string().optional(),
	},

	runtimeEnvStrict: {
		PORT: process.env.PORT,
	},
	emptyStringAsUndefined: true,
});
