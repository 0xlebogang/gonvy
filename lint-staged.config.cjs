const path = require("node:path");

function buildGoLintCommand(files) {
	const packages = new Set(files.map((f) => path.dirname(f)));
	const commands = Array.from(packages).map(
		(p) => `golangci-lint run --fix ${p}`,
	);
	return commands;
}

module.exports = {
	"*.{js,cjs,mjs,ts,mts,tsx,json,css,scss}": "biome check --write",
	"*.go": [buildGoLintCommand, "gofmt -w"],
};
