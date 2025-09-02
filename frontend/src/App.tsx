import { useState, useEffect } from "react";
import { TimebookService } from "../bindings/timebook";

import "./App.css";

export function App() {
    const [content, setContent] = useState<React.ReactNode>("");

    useEffect(() => {
        TimebookService.ParseFile("./.exampleFiles/2025-Q3.md").then((result) => {
            if (result === null) {
                setContent("null");
                return;
            }

            const bars = Object.entries(result.Entries)
                .map(([key, value]) => ({
                    key,
                    minutes: value,
                    percentage: Math.round((value / result.TotalMins) * 100),
                }))
                .sort((a, b) => b.minutes - a.minutes)
                .map((entry, _, entries) => {
                    const width = Math.round((entry.minutes / entries[0].minutes) * 100);

                    const hue = 120 - Math.round((width / 100) * 120); // 0 (red) to 120 (green)

                    return (
                        <div
                            key={entry.key}
                            className="bar"
                            style={{ width: width + "%", backgroundColor: `hsl(${hue}, 70%, 50%)` }}
                        >
                            {entry.key}: {entry.percentage}% ({entry.minutes} mins)
                        </div>
                    );
                });

            setContent(bars);
        });
    }, []);

    return <div className="container">{content}</div>;
}
