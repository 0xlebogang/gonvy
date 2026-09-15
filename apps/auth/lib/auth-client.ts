import { CreateAuthClient } from "@workspace/better-auth/client";

export const { useSession, signIn, signOut, signUp } = CreateAuthClient();
