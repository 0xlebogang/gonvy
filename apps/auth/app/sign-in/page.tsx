import { Button } from "@workspace/ui/components/button";

export default function SignIn() {
	return (
		<div className="flex min-h-svh p-6">
			<div className="flex max-w-md min-w-0 flex-col gap-4 text-sm leading-loose items-center justify-center">
				<div>
					<h1 className="text-5xl mb-6">Sign In Page</h1>
					<Button className="mt-2">Button</Button>
				</div>
			</div>
		</div>
	);
}
