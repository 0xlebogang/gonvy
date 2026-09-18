import { type NextRequest, NextResponse } from "next/server";
import { env } from "./env.server";
import { fetchActiveSession } from "./service/auth.service";

export async function proxy(_request: NextRequest) {
	const session = await fetchActiveSession();
	if (!session?.session) {
		return NextResponse.redirect(new URL(env.AUTH_APP_URL));
	}

	return NextResponse.next();
}

export const config = {
	matcher: "/:path*",
};
