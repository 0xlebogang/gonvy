'use client'

import { Theme, Grid, Column, FlexGrid } from "@repo/ui";
import * as React from "react";

export interface AppShellProps {
	path: string;
	children: React.ReactNode
}

export default function AppShell(props: AppShellProps) {
	const [showFooter, setShowFooter] = React.useState<boolean>()

	React.useEffect(() => {
		if (props.path !== "/") {
			setShowFooter(false)
			return
		}
		setShowFooter(true)
	}, [props.path])

	return (
		<Theme theme="g90">
			<Grid withRowGap>
				<Column span="100%">
					<FlexGrid as="div" align="center">
						<p>Global Navbar | Placeholder</p>
					</FlexGrid>
				</Column>
				<Column span="100%">
					{props.children}
				</Column>
			</Grid>
		</Theme>
	)
}
