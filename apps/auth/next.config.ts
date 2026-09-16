import "./env.server";
import type { NextConfig } from "next";

const nextConfig: NextConfig = {
	transpilePackages: [
		"@workspace/ui",
		"@workspace/better-auth",
		"@workspace/env",
	],
};

export default nextConfig;
