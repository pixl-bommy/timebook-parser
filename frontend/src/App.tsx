import { Route, Routes } from "react-router-dom";
import { TimebookPage } from "./pages/TimebookPage";

export function App() {
    return (
        <Routes>
            <Route path="/" element={<TimebookPage />} />
        </Routes>
    );
}
