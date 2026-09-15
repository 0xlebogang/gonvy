import { constants, copyFile } from "node:fs";
import { glob } from "glob";

async function main() {
	const exampleFiles = await glob("**/*.example", {
		cwd: process.cwd(),
		dot: true,
		dotRelative: true,
	});
	for (const f of exampleFiles) {
		copyFile(
			`${f}`,
			`${f.replace(".example", "")}`,
			constants.COPYFILE_EXCL,
			(error) => error && console.error("failed to copy file:", f),
		);
	}
}

main().catch((error) => console.error("operation failed:", error.message));
