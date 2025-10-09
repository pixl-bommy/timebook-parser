import { TimebookSummary } from "../../../bindings/timebook";

import "./CakeView.css";

export function CakeView({ timebookSummary }: { timebookSummary: TimebookSummary }) {
    let lastStartAngle = 0;

    const slices = timebookSummary.Entries.sort(
        (a, b) => b.ReceivedMinutes - a.ReceivedMinutes,
    ).map((entry, index, entries) => {
        const hue = Math.round((index / entries.length) * 120); // 0 (red) to 120 (green)

        const startAngle = lastStartAngle;
        const angle = entry.FactorOfTotal * 360;

        lastStartAngle += angle;

        return (
            <div
                key={entry.TaskName}
                className="cake-slice"
                style={{
                    // @ts-expect-error CSS variable
                    "--start": `${startAngle}deg`,
                    "--angle": `${angle.toFixed(0)}deg`,
                    "--sliceColor": `hsl(${hue}, 70%, 30%)`,
                }}
            >
                <div>{entry.TaskShort}</div>
            </div>
        );
    });

    return <div className="cake">{slices}</div>;
}
