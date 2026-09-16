import { auth } from "@workspace/better-auth/server";
import { headers } from "next/headers";
import { type NextRequest, NextResponse } from "next/server";
import { env } from "./env.server";

export async function proxy(_req: NextRequest) {
	const session = await auth.api.getSession({
		headers: await headers(),
	});

	if (!session) return NextResponse.redirect(env.BETTER_AUTH_URL);

	return NextResponse.next();
}

export const config = {
	matcher: ["/:path*"],
};
