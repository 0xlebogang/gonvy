import { createAuthClient } from "better-auth/react";

export function authClient(baseUrl?: string) {
	return createAuthClient({
		baseURL: baseUrl,
		fetchOptions: {
			credentials: "include",
		},
		plugins: [],
	});
}
