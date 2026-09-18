import { createNextjsEnv } from "@workspace/env";
import { globalEnv } from "@workspace/env/server";
import * as z from "zod/v3";

// Regex checking for relative path prefixes (./, ../), absolute paths (/), or standard file names ending in an extension
const sqliteDBFilepathRegex =
	/^(\.?\.?\/|[a-zA-Z]:\\|\/)?[\w\-. /]+\.[a-zA-Z0-0]+$/;

// Regex supporting standard domains, subdomains, and optionally localhost
const domainRegex =
	/^(?:[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,}$|^localhost$/;

export const env = createNextjsEnv({
	extends: [globalEnv],
	server: {
		BETTER_AUTH_DOMAIN: z
			.string()
			.trim()
			.min(1, "BETTER_AUTH_DOMAIN is required")
			.regex(
				domainRegex,
				"Must be a valid domain name (e.g., example.com, auth.dev.com or localhost",
			),
		CORS_ALLOWED_ORIGINS: z
			.string()
			.transform((value) =>
				value
					.split(",")
					.map((host) => host.trim())
					.filter(Boolean),
			)
			.pipe(z.array(z.string().url())),
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
