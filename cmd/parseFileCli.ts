import { writeFileSync } from "node:fs";
import { loadFile } from "../internal/loadFile.ts";
import {
    filterFileContent,
    mapTaskLinesToSummary,
    parseTaskLine,
} from "../internal/parseFileContent.ts";
import path from "node:path";
import { visualizeSummary } from "../internal/visualizeSummary.ts";

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
console.log(); // just add a new line for better readability

// pre-filter the file content
console.log("Pre-filtering file content");
const filteredContent = filterFileContent(fileContent);
console.log(`  - lines : ${filteredContent.length}`);
console.log(); // just add a new line for better readability

// parse the file content
console.log("Parsing file content");
const parsed = filteredContent
    .map((line) => {
        try {
            return parseTaskLine(line);
        } catch (_) {}
    })
    .filter((line) => line !== undefined);
console.log(`  - tasks : ${parsed.length}`);
console.log(); // just add a new line for better readability

// calculate the summary
console.log("Calculating summary");
const summary = mapTaskLinesToSummary(parsed);
console.log(`  - total hours   : ${summary.totalHours.toFixed(2)}`);
console.log(`  - total minutes : ${summary.totalMinutes}`);
console.log(); // just add a new line for better readability

// write a visualization to the console
console.log("Visualizing summary");
const visualization = visualizeSummary(summary);
console.log(visualization);
console.log(); // just add a new line for better readability

writeFileSync(path.resolve(".exampleFiles/out.md"), visualization, "utf-8");
