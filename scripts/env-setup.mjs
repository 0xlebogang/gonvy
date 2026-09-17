import { copyFile } from "node:fs";
import { globSync } from "glob";

function main() {
	const exampleFiles = globSync("**/*.example", {
		ignore: "node_modules",
		dotRelative: true,
		dot: true,
	});
	for (const f of exampleFiles) {
		copyFile(f, f.replace(".example", ""), (err) => err && console.error(err));
	}
}

main();
