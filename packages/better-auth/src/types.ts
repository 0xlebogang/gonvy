import { authClient } from "./client";

// Initialized solely for the purpose of accessing the required session data type
const client = authClient();

export type SessionData = typeof client.$Infer.Session;
