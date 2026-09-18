import "./env.server";
import type { NextConfig } from "next";
import { env } from "./env.server";

const nextConfig: NextConfig = {
	transpilePackages: ["@workspace/ui"],
	devIndicators: {
		position: "bottom-right",
	},
	rewrites: () => [
		{
			source: "/api/auth/:path*",
			destination: `${env.BETTER_AUTH_URL}/api/auth/:path*`,
		},
	],
};

export default nextConfig;
