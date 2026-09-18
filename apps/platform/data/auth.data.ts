import { auth } from "@workspace/better-auth/server";
import { headers } from "next/headers";
import { cache } from "react";

export const fetchActiveSession = cache(
	async () => await auth.api.getSession({ headers: await headers() }),
);
