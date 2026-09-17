import { createNextjsEnv } from "@workspace/env";
import { globalEnv } from "@workspace/env/server";
import * as z from "zod/v3";

// Regex checking for relative path prefixes (./, ../), absolute paths (/), or standard file names ending in an extension
const sqliteDBFilepathRegex =
	/^(\.?\.?\/|[a-zA-Z]:\\|\/)?[\w\-. /]+\.[a-zA-Z0-0]+$/;

export const env = createNextjsEnv({
	extends: [globalEnv],
	server: {
		DATABASE_URL: z.union([
			z.string().url(),
			z
				.string()
				.regex(
					sqliteDBFilepathRegex,
					"Must be a valid relative or absolute file path",
				),
		]),
		BETTER_AUTH_URL: z.string().url("Must be a URL"),
		BETTER_AUTH_SECRET: z
			.string()
			.min(32, "Must be at least 32 characters long"),
	},
	experimental__runtimeEnv: process.env,
});
