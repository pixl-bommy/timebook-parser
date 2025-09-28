import { useState, useEffect } from "react";
import { GoSync } from "react-icons/go";

import { TimebookService } from "../bindings/timebook";
import { CakeView } from "./views/CakeView";

import "./App.css";

export function App() {
    const [content, setContent] = useState<React.ReactNode>("");
    const [filename, setFilename] = useState<string>("");

    useEffect(() => {
        if (filename === "") {
            setContent("No file selected");
            return;
        }

        handleLoadFile();
    }, [filename]);

    async function handleLoadFile() {
        try {
            const timebookSummary = await TimebookService.ParseFile(filename);
            if (!timebookSummary) {
                setContent(<div>Something went wrong.</div>);
                return;
            }

            console.log("Parsed result:", timebookSummary);

            const view = <CakeView timebookSummary={timebookSummary} />;
            setContent(view);
        } catch (error) {
            setContent(<div>Error loading file: {(error as Error)?.message}</div>);
        }
    }

    async function handleFileSelect() {
        try {
            const filePath = await TimebookService.SelectFile();
            if (!filePath) throw new Error("Selection cancelled, as no file path was returned.");

            setFilename(filePath);
        } catch (error) {
            console.log("File selection was cancelled due to error.", error);
            setFilename("");
        }
    }

    return (
        <>
            <div className="container">{content}</div>
            <div>
                <button onClick={handleFileSelect}>Open Timebook File</button>
                {filename && (
                    <button onClick={handleLoadFile}>
                        <GoSync />
                    </button>
                )}
            </div>
            <div>{filename && <p>Selected file: {filename}</p>}</div>
        </>
    );
}
