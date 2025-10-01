import { useState, useEffect } from "react";

import { BiSolidBarChartAlt2, BiSolidPieChartAlt2 } from "react-icons/bi";
import { GoSync } from "react-icons/go";

import { TimebookService, TimebookSummary } from "../bindings/timebook";
import { CakeView } from "./views/CakeView";
import { HorizontalBarView } from "./views/HorizontalBarView";
import { HorizontalCategoryBarView } from "./views/HorizontalCategoryBarView";

import "./App.css";

export function App() {
    const [currentView, setCurrentView] = useState<"cake" | "bar" | "barCategory" | null>("barCategory");
    const [timebookSummary, setTimebookSummary] = useState<TimebookSummary | null>(null);
    const [filename, setFilename] = useState<string>("");

    useEffect(() => {
        if (filename === "") {
            setTimebookSummary(null);
            return;
        }

        handleLoadFile();
    }, [filename]);

    async function handleLoadFile() {
        try {
            const timebookSummary = await TimebookService.LoadFile(filename);
            if (!timebookSummary) {
                setTimebookSummary(null);
                return;
            }

            console.log("Parsed result:", timebookSummary);
            setTimebookSummary(timebookSummary);
        } catch (error) {
            setTimebookSummary(null);
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
            <div className="container">
                {timebookSummary
                    ? (currentView === "cake" && <CakeView timebookSummary={timebookSummary} />) ||
                      (currentView === "bar" && (
                          <HorizontalBarView timebookSummary={timebookSummary} />
                      )) ||
                      (currentView === "barCategory" && (
                          <HorizontalCategoryBarView timebookSummary={timebookSummary} />
                      )) ||
                      "No view selected."
                    : "No data to display."}
            </div>
            <div className="toolbar">
                <button onClick={handleFileSelect}>Open Timebook File</button>
                {filename && (
                    <button onClick={handleLoadFile}>
                        <GoSync />
                    </button>
                )}
                <div>
                    <button onClick={() => setCurrentView("bar")}>
                        <BiSolidBarChartAlt2 />
                    </button>
                    <button onClick={() => setCurrentView("cake")}>
                        <BiSolidPieChartAlt2 />
                    </button>
                </div>
            </div>
            <div className="toolbar">{filename && `Selected file: ${filename}`}</div>
        </>
    );
}
