import { toNextJsHandler } from "@workspace/better-auth";
import { auth } from "@workspace/better-auth/server";

export const { GET, POST } = toNextJsHandler(auth);
