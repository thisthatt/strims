"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const child_process_1 = require("child_process");
const crypto = require("crypto");
const fs_1 = require("fs");
const path_1 = require("path");
const versionPkg = "github.com/MemeLabs/go-ppspp/pkg/version";
function loader(contents) {
    const cb = this.async();
    const opts = {
        env: {
            GOPATH: process.env.GOPATH,
            GOROOT: process.env.GOROOT,
            GOCACHE: path_1.join(__dirname, "./.gocache"),
            GOOS: "js",
            GOARCH: "wasm",
        },
    };
    const goBin = `${opts.env.GOROOT}/bin/go`;
    const outFile = `${this.resourcePath}.wasm`;
    const args = ["build", "-mod", "readonly"];
    if (this.mode === "production") {
        const rev = process.env.VERSION || child_process_1.execFileSync("git", ["rev-parse", "HEAD"]).toString().substr(0, 8);
        args.push("-trimpath", "-ldflags", `-s -w -X '${versionPkg}.Platform=web' -X '${versionPkg}.Version=${rev}'`);
    }
    else {
        args.push("-ldflags", `-X '${versionPkg}.Platform=web'`);
    }
    args.push("-o", outFile, this.resourcePath);
    child_process_1.execFile(goBin, args, opts, (err) => {
        if (err) {
            cb(err);
            return;
        }
        const out = fs_1.readFileSync(outFile);
        const hash = crypto.createHash("sha256");
        hash.write(out);
        const digest = hash.digest().toString("hex").substring(0, 20);
        fs_1.unlinkSync(outFile);
        const emittedFilename = path_1.basename(this.resourcePath, ".go") + `.${digest}.wasm`;
        this.emitFile(emittedFilename, out, null);
        cb(null, [
            `require("${path_1.join(__dirname, "..", "lib", "wasm_exec.js")}");`,
            `import gobridge from "${path_1.join(__dirname, "..", "dist", "gobridge.js")}";`,
            `export default gobridge("${emittedFilename}");`,
        ].join("\n"));
    });
}
exports.default = loader;
//# sourceMappingURL=index.js.map