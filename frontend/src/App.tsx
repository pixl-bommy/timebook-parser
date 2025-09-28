import { useState, useEffect, useCallback } from "react";
import { GoSync } from "react-icons/go";
import { TimebookService } from "../bindings/timebook";

import "./App.css";
import { HorizontalBarView } from "./views/HorizontalBarView";

export function App() {
    const [content, setContent] = useState<React.ReactNode>("");
    const [filename, setFilename] = useState<string>("");

    useEffect(() => {
        if (filename === "") {
            setContent("No file selected");
            return;
        }

        loadFile();
    }, [filename]);

    const loadFile = useCallback(async () => {
        try {
            const timebookSummary = await TimebookService.ParseFile(filename);
            if (!timebookSummary) {
                setContent(<div>Something went wrong.</div>);
                return;
            }

            console.log("Parsed result:", timebookSummary);

            const view = <HorizontalBarView timebookSummary={timebookSummary} />;
            setContent(view);
        } catch (error) {
            setContent(<div>Error loading file: {(error as Error)?.message}</div>);
        }
    }, [filename]);

    function handleFileSelect() {
        TimebookService.SelectFile().then((filePath) => {
            if (filePath) {
                setFilename(filePath);
            }
        });
    }

    return (
        <>
            <div className="container">{content}</div>
            <div>
                <button onClick={handleFileSelect}>Open Timebook File</button>
                {filename && (
                    <button onClick={() => loadFile()}>
                        <GoSync />
                    </button>
                )}
            </div>
            <div>{filename && <p>Selected file: {filename}</p>}</div>
        </>
    );
}
