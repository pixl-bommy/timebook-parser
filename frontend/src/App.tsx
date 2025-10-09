import { Link, Route, Routes } from "react-router-dom";
import { TimebookPage } from "./pages/TimebookPage";
import { DummyTestPage } from "./pages/DummyTestPage";

import "./App.css";

export function App() {
    return (
        <>
            <nav>
                <Link to="/">Timebook</Link>
                <Link to="/dummy">Dummy Test</Link>
            </nav>
            <main>
                <Routes>
                    <Route path="/" element={<TimebookPage />} />
                    <Route path="/dummy" element={<DummyTestPage />} />
                </Routes>
            </main>
        </>
    );
}
