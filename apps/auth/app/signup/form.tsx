"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { Button } from "@workspace/ui/components/button";
import {
	Field,
	FieldDescription,
	FieldGroup,
	FieldLabel,
	FieldSeparator,
} from "@workspace/ui/components/field";
import { Input } from "@workspace/ui/components/input";
import { Spinner } from "@workspace/ui/components/spinner";
import { cn } from "cn";
import { GalleryVerticalEndIcon } from "lucide-react";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { useForm } from "react-hook-form";
import * as z from "zod/v3";
import { signUp } from "@/lib/auth-client";

export interface SignUpFormProps extends React.ComponentProps<"div"> {
	searchParams?: string;
}

export const signUpSchema = z
	.object({
		name: z.string().max(255, "Name is too long"),
		email: z.string().min(1, "Emaill is required").email(),
		password: z
			.string()
			.min(8, "Password is too short")
			.max(128, "Password is too long"),
		confirmPassword: z.string(),
		isLoading: z.boolean(),
	})
	.refine((values) => values.password === values.confirmPassword, {
		message: "Passwords do not match",
		path: ["confirmPassword"],
	});

export type SignUpSchema = z.infer<typeof signUpSchema>;

export function Form({ className, searchParams, ...props }: SignUpFormProps) {
	const router = useRouter();

	const {
		register,
		reset,
		handleSubmit,
		setValue,
		getValues,
		formState: { errors },
	} = useForm<SignUpSchema>({
		resolver: zodResolver(signUpSchema),
		defaultValues: {
			name: "",
			email: "",
			password: "",
			confirmPassword: "",
			isLoading: false,
		},
	});

	async function onSubmit(formData: SignUpSchema) {
		setValue("isLoading", true);

		try {
			const { error } = await signUp.email({
				name: formData.name,
				email: formData.email,
				password: formData.password,
			});

			if (error) {
				alert(error.message);
				return;
			}

			router.push("/");
			router.refresh();
		} catch (err) {
			console.error(err);

			return;
		} finally {
			setValue("isLoading", false);
			reset();
		}
	}

	return (
		<div className={cn("flex flex-col gap-6", className)} {...props}>
			<form
				onSubmit={handleSubmit(onSubmit)}
				aria-disabled={getValues("isLoading")}
			>
				<FieldGroup>
					<div className="flex flex-col items-center gap-2 text-center">
						<Link
							href="/"
							className="flex flex-col items-center gap-2 font-medium"
						>
							<div className="flex size-8 items-center justify-center rounded-md">
								<GalleryVerticalEndIcon className="size-6" />
							</div>
							<span className="sr-only">Gonvy.</span>
						</Link>
						<h1 className="text-xl font-bold">Welcome to Gonvy.</h1>
						<FieldDescription>
							Already have an account? <a href="/signin">Sign in</a>
						</FieldDescription>
					</div>

					<Field>
						<FieldLabel htmlFor="name">Name</FieldLabel>
						<Input {...register("name")} type="text" placeholder="John Doe" />
						{errors.name && (
							<small className="text-red-600">{errors.name.message}</small>
						)}
					</Field>

					<Field>
						<FieldLabel htmlFor="email">Email</FieldLabel>
						<Input
							{...register("email", { required: true })}
							type="email"
							placeholder="m@example.com"
						/>
						{errors.email && (
							<small className="text-red-600">{errors.email.message}</small>
						)}
					</Field>

					<Field>
						<FieldLabel htmlFor="password">Password</FieldLabel>
						<Input
							{...register("password", { required: true })}
							type="password"
							placeholder="********"
						/>
						{errors.password && (
							<small className="text-red-600">{errors.password.message}</small>
						)}
					</Field>

					<Field>
						<FieldLabel htmlFor="confirmPassword">Confirm Password</FieldLabel>
						<Input
							{...register("confirmPassword", { required: true })}
							type="password"
							placeholder="********"
						/>
						{errors.confirmPassword && (
							<small className="text-red-600">
								{errors.confirmPassword.message}
							</small>
						)}
					</Field>

					<Field>
						<Button type="submit" disabled={getValues("isLoading")}>
							{getValues("isLoading") ? <Spinner /> : "Create Account"}
						</Button>
					</Field>

					<FieldSeparator>Or</FieldSeparator>

					<Field className="grid gap-4 sm:grid-cols-2">
						<Button variant="outline" type="button">
							<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
								<title>Apple</title>
								<path
									d="M12.152 6.896c-.948 0-2.415-1.078-3.96-1.04-2.04.027-3.91 1.183-4.961 3.014-2.117 3.675-.546 9.103 1.519 12.09 1.013 1.454 2.208 3.09 3.792 3.039 1.52-.065 2.09-.987 3.935-.987 1.831 0 2.35.987 3.96.948 1.637-.026 2.676-1.48 3.676-2.948 1.156-1.688 1.636-3.325 1.662-3.415-.039-.013-3.182-1.221-3.22-4.857-.026-3.04 2.48-4.494 2.597-4.559-1.429-2.09-3.623-2.324-4.39-2.376-2-.156-3.675 1.09-4.61 1.09zM15.53 3.83c.843-1.012 1.4-2.427 1.245-3.83-1.207.052-2.662.805-3.532 1.818-.78.896-1.454 2.338-1.273 3.714 1.338.104 2.715-.688 3.559-1.701"
									fill="currentColor"
								/>
							</svg>
							Continue with Apple
						</Button>

						<Button variant="outline" type="button">
							<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
								<title>Google</title>
								<path
									d="M12.48 10.92v3.28h7.84c-.24 1.84-.853 3.187-1.787 4.133-1.147 1.147-2.933 2.4-6.053 2.4-4.827 0-8.6-3.893-8.6-8.72s3.773-8.72 8.6-8.72c2.6 0 4.507 1.027 5.907 2.347l2.307-2.307C18.747 1.44 16.133 0 12.48 0 5.867 0 .307 5.387.307 12s5.56 12 12.173 12c3.573 0 6.267-1.173 8.373-3.36 2.16-2.16 2.84-5.213 2.84-7.667 0-.76-.053-1.467-.173-2.053H12.48z"
									fill="currentColor"
								/>
							</svg>
							Continue with Google
						</Button>
					</Field>
				</FieldGroup>
			</form>
		</div>
	);
}
