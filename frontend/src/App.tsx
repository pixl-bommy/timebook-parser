import { useState, useEffect, useCallback } from "react";
import { GoSync } from "react-icons/go";
import { TimebookService } from "../bindings/timebook";

import "./App.css";

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
            const result = await TimebookService.ParseFile(filename);
            if (!result) {
                setContent(<div>Something went wrong.</div>);
                return;
            }

            console.log("Parsed result:", result);

            const bars = result.Entries.sort((a, b) => b.ReceivedMinutes - a.ReceivedMinutes).map(
                (entry, _, entries) => {
                    const width = Math.round((entry.FactorOfTotal / entries[0].FactorOfTotal) * 100);

                    const hue = 120 - Math.round((width / 100) * 120); // 0 (red) to 120 (green)
                    const durationHours = (entry.ReceivedMinutes / 60).toFixed(1);
                    const percentage = (entry.FactorOfTotal*100).toFixed(0);

                    return (
                        <div
                            key={entry.TaskName}
                            className="bar"
                            style={{
                                width: width + "%",
                                backgroundColor: `hsl(${hue}, 70%, 30%)`,
                            }}
                        >
                            {entry.TaskName}: {percentage}% ({durationHours} hours)
                        </div>
                    );
                },
            );

            setContent(bars);
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
