import fs from "node:fs";
import path from "node:path";

/**
 * Load a file from the filesystem.
 * @param filePath the path to the file to load
 * @returns The content of the file as a string.
 * @throws {Error} If the file path is not provided, if the file does not exist
 *      or if there is an error reading the file.
 */
export function loadFile(filePath: string | undefined): string[] {
    if (!filePath) {
        throw new Error("File path is required");
    }

    // Resolve the file path to an absolute path
    const absolutePath = path.resolve(filePath);

    // Check if the file exists
    if (!fs.existsSync(absolutePath)) {
        throw new Error(`File not found: ${absolutePath}`);
    }

    // Read the file content
    const content = fs.readFileSync(absolutePath, "utf-8");
    return content.split(/\r?\n/);
}
