import * as React from 'react'
import AppShell from '@repo/ui/components/app-shell'

import '@repo/ui/styles/main.scss'
import Footer from './footer'

export interface RootLayoutProps {
	children: React.ReactNode
}

export default function RootLayout(props: RootLayoutProps) {
	return (
		<html lang='en' suppressHydrationWarning>
			<body>
				<AppShell path={""}>
					{props.children}
					<Footer />
				</AppShell>
			</body>
		</html>
	)
}
