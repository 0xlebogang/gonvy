import { createAuthClient } from "better-auth/client";

export function CreateAuthClient(baseUrl?: string) {
	return createAuthClient({
		baseURL: baseUrl,
		fetchOptions: {
			credentials: "include",
		},
	});
}
