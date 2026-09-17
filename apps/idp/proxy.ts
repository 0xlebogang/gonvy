import { type NextRequest, NextResponse } from "next/server";
import { env } from "./env";

const allowedOrigins = env.CORS_ALLOWED_ORIGINS;

export function proxy(request: NextRequest) {
	const origin = request.headers.get("origin");
	const isAllowed = origin && allowedOrigins.includes(origin);

	if (request.method === "OPTIONS") {
		const response = new NextResponse(null, { status: 204 });

		if (isAllowed) {
			response.headers.set("Access-Control-Allow-Origin", origin);
			response.headers.set(
				"Access-Control-Allow-Methods",
				"GET, POST, PUT, DELETE, OPTIONS",
			);
			response.headers.set("Access-Control-Allow-Credentials", "true");
		}

		return response;
	}

	const response = NextResponse.next();
	if (isAllowed) {
		response.headers.set("Access-Control-Allow-Origin", origin);
		response.headers.set("Access-Control-Allow-Credentials", "true");
	}

	return response;
}

export const config = {
	matcher: "/:path*",
};
