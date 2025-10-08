import { SummaryEntry, TimebookSummary } from "../../bindings/timebook";

import "./HorizontalBarView.css";

export function HorizontalCategoryBarView({
    timebookSummary,
}: {
    timebookSummary: TimebookSummary;
}) {
    const bars = timebookSummary.Entries.reduce((reduced, entry) => {
        const existingEntry = reduced.find((e) => e.CategoryShort === entry.CategoryShort);
        if (existingEntry) {
            existingEntry.ReceivedMinutes += entry.ReceivedMinutes;
            existingEntry.FactorOfTotal += entry.FactorOfTotal;
            existingEntry.ExpectedMinutes += entry.ExpectedMinutes;
            existingEntry.FactorOfExpected += entry.FactorOfExpected;
            existingEntry.CountTasks += entry.CountTasks;
        } else {
            reduced.push({
                CategoryShort: entry.CategoryShort,
                CategoryName: entry.CategoryName,
                ReceivedMinutes: entry.ReceivedMinutes,
                FactorOfTotal: entry.FactorOfTotal,
                ExpectedMinutes: entry.ExpectedMinutes,
                FactorOfExpected: entry.FactorOfExpected,
                CountTasks: entry.CountTasks,
            });
        }
        return reduced;
    }, [] as Omit<SummaryEntry, "TaskShort" | "TaskName">[])
        .sort((a, b) => b.ReceivedMinutes - a.ReceivedMinutes)
        .map((entry, _, entries) => {
            const width = Math.round((entry.FactorOfTotal / entries[0].FactorOfTotal) * 100);

            const hue = 120 - Math.round((width / 100) * 120); // 0 (red) to 120 (green)
            const durationHours = (entry.ReceivedMinutes / 60).toFixed(1);
            const percentage = (entry.FactorOfTotal * 100).toFixed(0);

            return (
                <div
                    key={entry.CategoryName}
                    className="bar"
                    style={{
                        width: width + "%",
                        backgroundColor: `hsl(${hue}, 70%, 30%)`,
                    }}
                >
                    {entry.CategoryName}: {percentage}% ({durationHours} hours)
                </div>
            );
        });

    return <div className="bars">{bars}</div>;
}
