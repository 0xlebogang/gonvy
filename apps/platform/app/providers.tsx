import { SidebarProvider } from "@workspace/ui/components/sidebar";
import type React from "react";
import { ThemeProvider } from "@/components/theme-provider";

export default function Providers({ children }: { children: React.ReactNode }) {
	return (
		<ThemeProvider>
			<SidebarProvider>{children}</SidebarProvider>
		</ThemeProvider>
	);
}
