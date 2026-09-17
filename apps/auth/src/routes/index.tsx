import { createFileRoute } from "@tanstack/react-router";
import { Button } from "@workspace/ui/components/button";

export const Route = createFileRoute("/")({ component: Home });

function Home() {
	return (
		<div className="p-8 w-full min-h-screen flex flex-col items-center justify-center gap-4">
			<h1 className="text-4xl font-bold">Welcome to TanStack Start</h1>
			<p className="mt-4 text-lg">
				Edit <code>src/routes/index.tsx</code> to get started.
			</p>
			<Button onClick={() => alert("hello from Auth")}>Click Me!</Button>
		</div>
	);
}
