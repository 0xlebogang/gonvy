import { auth } from "@workspace/better-auth/server";
import { headers } from "next/headers";
import { type NextRequest, NextResponse } from "next/server";
import { env } from "./env.server";

export async function proxy(request: NextRequest) {
	if (request.nextUrl.pathname === "/") {
		return NextResponse.redirect(new URL("/login", request.url));
	}

	const session = await auth.api.getSession({
		headers: await headers(),
	});

	if (session?.session) {
		return NextResponse.redirect(new URL(env.PLATFORM_URL));
	}

	return NextResponse.next();
}

export const config = {
	matcher: "/:path*",
};
