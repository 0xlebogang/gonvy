import { Geist_Mono, Inter } from "next/font/google";

import "@workspace/ui/globals.css";
import { Toaster } from "@workspace/ui/components/toast";
import { cn } from "@workspace/ui/lib/utils";
import type { Metadata } from "next";
import { ThemeProvider } from "@/components/theme-provider";
import { env } from "@/env.server";

const inter = Inter({ subsets: ["latin"], variable: "--font-sans" });

const fontMono = Geist_Mono({
	subsets: ["latin"],
	variable: "--font-mono",
});

export const metadata: Metadata = {
	title: "Auth | Gonvy",
	description: "Authenticated to your Gonvy instance",
	other: {
		preconnect: env.PLATFORM_URL,
	},
};

export default function RootLayout({
	children,
}: Readonly<{
	children: React.ReactNode;
}>) {
	return (
		<html
			lang="en"
			suppressHydrationWarning
			className={cn(
				"antialiased",
				fontMono.variable,
				"font-sans",
				inter.variable,
			)}
		>
			<body>
				<ThemeProvider>
					<div className="flex min-h-svh flex-col items-center justify-center gap-6 bg-background p-6 md:p-10">
						<div className="w-full max-w-sm">{children}</div>
					</div>
					<Toaster />
				</ThemeProvider>
			</body>
		</html>
	);
}
