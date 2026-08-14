import * as React from 'react'
import '@repo/ui/main.scss'

export interface RootLayoutProps {
	children: React.ReactNode
}

export default function RootLayout({ children }: RootLayoutProps) {
	return (
		<html lang='en' suppressHydrationWarning>
			<body>
				{children}
			</body>
		</html>
	)
}
