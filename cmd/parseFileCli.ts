import { loadFile } from "../internal/loadFile.ts";

// read filename from command line arguments
const filename = process.argv[2];

// load the file content
console.log("Loading file content");
console.log(`  - file  : ${filename}`);
let fileContent: string[];
try {
    fileContent = loadFile(filename);
    console.log("  - status: loaded");
    console.log(`  - lines : ${fileContent.length}`);
} catch (error) {
    console.log(`  - status: error - ${error instanceof Error ? error.message : error}`);
    process.exit(1);
}
