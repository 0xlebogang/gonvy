import "./env.server";
import type { NextConfig } from "next";

const nextConfig: NextConfig = {
	transpilePackages: [
		"@workspace/ui",
		"@workspace/better-auth",
		"@workspace/env",
	],
	devIndicators: {
		position: "bottom-right",
	},
};

export default nextConfig;
