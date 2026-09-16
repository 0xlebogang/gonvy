import { auth } from "@workspace/better-auth/server";
import { headers } from "next/headers";
import { type NextRequest, NextResponse } from "next/server";
import { env } from "./env.server";

export async function proxy(req: NextRequest) {
	if (req.nextUrl.pathname.includes("/api/auth")) {
		return NextResponse.next();
	}

	const session = await auth.api.getSession({
		headers: await headers(),
	});

	if (session) {
		return NextResponse.redirect(env.PLATFORM_URL);
	}

	if (req.nextUrl.pathname === "/") {
		return NextResponse.redirect(new URL("/sign-in", req.url));
	}

	return NextResponse.next();
}

export const config = {
	matcher: ["/:path*"],
};
