import * as React from 'react'
import { Theme } from '@repo/ui/index'

import '@repo/ui/styles/main.scss'

export interface RootLayoutProps {
	children: React.ReactNode
}

export default function RootLayout(props: RootLayoutProps) {
	return (
		<html lang='en' suppressHydrationWarning>
			<body>
				<Theme theme='g90'>
					{props.children}
				</Theme>
			</body>
		</html>
	)
}
