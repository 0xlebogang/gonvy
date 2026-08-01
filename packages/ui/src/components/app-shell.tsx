"use client";

import { Column, FlexGrid, Grid, Theme } from "@repo/ui";
import type * as React from "react";

export interface AppShellProps {
	children: React.ReactNode;
}

export default function AppShell(props: AppShellProps) {
	return (
		<Theme theme="g90">
			<Grid withRowGap>
				<Column span="100%">
					<FlexGrid as="div" align="center">
						<p>Global Navbar | Placeholder</p>
					</FlexGrid>
				</Column>
				<Column span="100%">{props.children}</Column>
			</Grid>
		</Theme>
	);
}
