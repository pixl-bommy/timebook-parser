import { useState, useEffect } from "react";
import { TimebookService } from "../bindings/timebook";

function App() {
    const [json, setJson] = useState<React.ReactNode>("");

    useEffect(() => {
        TimebookService.ParseFile("./.exampleFiles/2025-Q3.md").then((result) => {
            const sum = Object.entries(result).reduce((acc, [key, duration]) => acc + duration, 0);
            setJson(JSON.stringify({ ...result, "∑": sum }, null, 4));
        });
    }, []);

    return (
        <div className="container">
            <code>{json}</code>
        </div>
    );
}

export default App;
