"use client";

import {
	DropdownMenu,
	DropdownMenuTrigger,
} from "@workspace/ui/components/dropdown-menu";
import {
	SidebarMenu,
	SidebarMenuButton,
	SidebarMenuItem,
} from "@workspace/ui/components/sidebar";
import { GalleryVerticalEndIcon } from "lucide-react";

export function Brand() {
	return (
		<SidebarMenu>
			<SidebarMenuItem>
				<DropdownMenu>
					<DropdownMenuTrigger render={<SidebarMenuButton size="lg" />}>
						<div className="flex aspect-square size-8 items-center justify-center rounded-lg bg-sidebar-primary text-sidebar-primary-foreground">
							<GalleryVerticalEndIcon />
						</div>
						<div className="grid flex-1 text-left text-sm leading-tight">
							<span className="truncate font-medium">Gonvy.</span>
						</div>
					</DropdownMenuTrigger>
				</DropdownMenu>
			</SidebarMenuItem>
		</SidebarMenu>
	);
}
