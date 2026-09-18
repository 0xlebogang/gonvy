import { SidebarInset } from "@workspace/ui/components/sidebar";
import { AppSidebar } from "@/components/app-sidebar";
import UIShell from "@/components/ui-shell";

export default function Dashboard() {
	return (
		<>
			<AppSidebar />
			<SidebarInset>
				<UIShell>
					<div className="grid auto-rows-min gap-4 md:grid-cols-3">
						<div className="aspect-video rounded-xl bg-muted/50" />
						<div className="aspect-video rounded-xl bg-muted/50" />
						<div className="aspect-video rounded-xl bg-muted/50" />
					</div>
					<div className="min-h-screen flex-1 rounded-xl bg-muted/50 md:min-h-min" />
				</UIShell>
			</SidebarInset>
		</>
	);
}
