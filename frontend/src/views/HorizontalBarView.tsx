import { TimebookSummary } from "../../bindings/timebook";

import "./HorizontalBarView.css";

export function HorizontalBarView({ timebookSummary }: { timebookSummary: TimebookSummary }) {
    const bars = timebookSummary.Entries.sort((a, b) => b.ReceivedMinutes - a.ReceivedMinutes).map(
        (entry, _, entries) => {
            const width = Math.round((entry.FactorOfTotal / entries[0].FactorOfTotal) * 100);

            const hue = 120 - Math.round((width / 100) * 120); // 0 (red) to 120 (green)
            const durationHours = (entry.ReceivedMinutes / 60).toFixed(1);
            const percentage = (entry.FactorOfTotal * 100).toFixed(0);

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

    return <>{bars}</>;
}
