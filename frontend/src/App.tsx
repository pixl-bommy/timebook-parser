import { Link, Route, Routes } from "react-router-dom";
import { GoCode } from "react-icons/go";

import { TimebookPage } from "./pages/TimebookPage";
import { DummyTestPage } from "./pages/DummyTestPage";

import "./App.css";
import { useState } from "react";

export function App() {
    const [isNavCollapsed, setIsNavCollapsed] = useState(false);

    return (
        <div id="app-root">
            <nav className={isNavCollapsed ? "collapsed" : ""}>
                <div className="nav-toggle" onClick={() => setIsNavCollapsed((prev) => !prev)}>
                    <GoCode />
                </div>

                <div className="nav-links">
                    <Link to="/">Timebook</Link>
                    <Link to="/dummy">Dummy Test</Link>
                </div>
            </nav>
            <main>
                <Routes>
                    <Route path="/" element={<TimebookPage />} />
                    <Route path="/dummy" element={<DummyTestPage />} />
                </Routes>
            </main>
        </div>
    );
}
