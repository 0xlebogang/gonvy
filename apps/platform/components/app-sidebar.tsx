import {
	Sidebar,
	SidebarContent,
	SidebarFooter,
	SidebarHeader,
	SidebarRail,
} from "@workspace/ui/components/sidebar";
import {
	CreditCardIcon,
	FolderGit2Icon,
	GalleryVerticalEndIcon,
	KeyRoundIcon,
	LayersIcon,
	LayoutDashboardIcon,
	Settings2Icon,
	ShieldCheckIcon,
	TerminalIcon,
} from "lucide-react";
import type * as React from "react";
import { NavMain } from "@/components/nav-main";
import { NavProjects } from "@/components/nav-projects";
import { NavUser } from "@/components/nav-user";
import { TeamSwitcher } from "@/components/team-switcher";
import { fetchActiveSession } from "@/data/auth.data";

// This is sample data.
const data = {
	teams: [
		{
			name: "Acme Inc",
			logo: <GalleryVerticalEndIcon />,
			plan: "Enterprise",
		},
	],
	navMain: [
		{
			title: "Projects",
			url: "/projects",
			icon: <FolderGit2Icon />,
			isActive: true,
			items: [
				{ title: "All Projects", url: "/projects" },
				{ title: "Environments", url: "/environments" },
				{ title: "Global Variables", url: "/variables/global" },
			],
		},
		{
			title: "Integrations",
			url: "/integrations",
			icon: <LayersIcon />,
			items: [
				{ title: "Connected Apps", url: "/integrations/active" },
				{ title: "Webhooks", url: "/integrations/webhooks" },
				{ title: "Pipes & Sync", url: "/integrations/sync" },
			],
		},
		{
			title: "Access & Security",
			url: "/security",
			icon: <KeyRoundIcon />,
			items: [
				{ title: "Service Accounts", url: "/security/service-accounts" },
				{ title: "API Tokens", url: "/security/tokens" },
				{ title: "Audit Logs", url: "/security/audit-logs" },
			],
		},
		{
			title: "Settings",
			url: "/settings",
			icon: <Settings2Icon />,
			items: [
				{ title: "General", url: "/settings/general" },
				{ title: "Members & RBAC", url: "/settings/members" },
				{ title: "Billing", url: "/settings/billing" },
			],
		},
	],
	projects: [
		{
			name: "Core API",
			url: "/projects/core-api",
			icon: <TerminalIcon />,
		},
		{
			name: "Web Dashboard",
			url: "/projects/web-dashboard",
			icon: <LayoutDashboardIcon />,
		},
		{
			name: "Auth Service",
			url: "/projects/auth-service",
			icon: <ShieldCheckIcon />,
		},
		{
			name: "Payment Gateway",
			url: "/projects/payment-gateway",
			icon: <CreditCardIcon />,
		},
	],
};

export async function AppSidebar({
	...props
}: React.ComponentProps<typeof Sidebar>) {
	const activeSession = await fetchActiveSession();

	return (
		<Sidebar collapsible="icon" {...props}>
			<SidebarHeader>
				<TeamSwitcher teams={data.teams} />
			</SidebarHeader>
			<SidebarContent>
				<NavMain items={data.navMain} />
				<NavProjects projects={data.projects} />
			</SidebarContent>
			<SidebarFooter>
				<NavUser user={activeSession?.user || null} />
			</SidebarFooter>
			<SidebarRail />
		</Sidebar>
	);
}
