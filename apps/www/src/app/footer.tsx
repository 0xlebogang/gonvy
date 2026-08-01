"use client";

import { Column, FlexGrid } from "@repo/ui/index";
import { usePathname } from "next/navigation";
import React from "react";

export default function Footer() {
	const [showFooter, setShowFooter] = React.useState<boolean>(true);
	const pathname = usePathname();

	React.useEffect(() => {
		if (pathname !== "/") {
			setShowFooter(false);
		}
	}, [pathname]);

	if (!showFooter) {
		return null;
	}

	return (
		<Column span="100%">
			<FlexGrid as="div" align="center">
				<p>Footer | Placeholder</p>
			</FlexGrid>
		</Column>
	);
}
