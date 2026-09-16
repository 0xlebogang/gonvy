import { CreateAuthClient } from "@workspace/better-auth/client";

export const { getSession, useSession, signOut } = CreateAuthClient();
