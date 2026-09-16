"use client";

import {
	Avatar,
	AvatarFallback,
	AvatarImage,
} from "@workspace/ui/components/avatar";
import {
	DropdownMenu,
	DropdownMenuContent,
	DropdownMenuGroup,
	DropdownMenuItem,
	DropdownMenuLabel,
	DropdownMenuSeparator,
	DropdownMenuTrigger,
} from "@workspace/ui/components/dropdown-menu";
import {
	SidebarMenu,
	SidebarMenuButton,
	SidebarMenuItem,
	useSidebar,
} from "@workspace/ui/components/sidebar";
import {
	ChevronsUpDownIcon,
	LogOutIcon,
	Settings,
	Toolbox,
	User,
} from "lucide-react";
import Link from "next/link";
import { useSession } from "@/lib/auth-client";

export function NavUser() {
	const { isMobile } = useSidebar();

	const { data: session } = useSession.get();
	const user = session?.user;

	const fallbackInitials = user?.name
		? user.name
				.split(" ")
				.map((name) => name[0])
				.join("")
				.toUpperCase()
				.slice(0, 2)
		: "GU";

	return (
		<SidebarMenu>
			<SidebarMenuItem>
				<DropdownMenu>
					<DropdownMenuTrigger
						render={
							<SidebarMenuButton size="lg" className="aria-expanded:bg-muted" />
						}
					>
						<Avatar>
							<AvatarImage src={user?.image as string} alt={"user"} />
							<AvatarFallback>{fallbackInitials}</AvatarFallback>
						</Avatar>
						<div className="grid flex-1 text-left text-sm leading-tight">
							{user?.name && (
								<span className="truncate font-medium">{user?.name}</span>
							)}
							<span className="truncate text-xs">{user?.email}</span>
						</div>
						<ChevronsUpDownIcon className="ml-auto size-4" />
					</DropdownMenuTrigger>
					<DropdownMenuContent
						className="w-fit"
						side={isMobile ? "bottom" : "right"}
						align="end"
						sideOffset={4}
					>
						<DropdownMenuGroup>
							<DropdownMenuLabel className="p-0 font-normal">
								<div className="flex items-center gap-2 px-1 py-1.5 text-left text-sm">
									<Avatar>
										<AvatarImage src={"user.avatar"} alt={"user.name"} />
										<AvatarFallback>CN</AvatarFallback>
									</Avatar>
									<div className="grid flex-1 text-left text-sm leading-tight">
										<span className="truncate font-medium">{user?.name}</span>
										<span className="truncate text-xs">{user?.email}</span>
									</div>
								</div>
							</DropdownMenuLabel>
						</DropdownMenuGroup>
						<DropdownMenuSeparator />
						<DropdownMenuGroup>
							<DropdownMenuItem render={<Link href="/profile" />}>
								<User />
								Profile
							</DropdownMenuItem>
							<DropdownMenuItem>
								<Toolbox />
								Preferences
							</DropdownMenuItem>
							<DropdownMenuItem>
								<Settings />
								Settings
							</DropdownMenuItem>
						</DropdownMenuGroup>
						<DropdownMenuSeparator />
						<DropdownMenuItem onClick={() => alert("Sign Out triggered")}>
							<LogOutIcon />
							Log out
						</DropdownMenuItem>
					</DropdownMenuContent>
				</DropdownMenu>
			</SidebarMenuItem>
		</SidebarMenu>
	);
}
