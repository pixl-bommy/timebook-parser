import type { TaskSummary } from "./parseFileContent.ts";

export function visualizeSummary(summary: TaskSummary): string {
    const output: string[] = [
        "# Summary of tasks",
        "",
        "## Totals",
        "",
        `  - tasks  : ${summary.tasks.length}`,
        `  - hours  : ${summary.totalHours.toFixed(2)}`,
        `  - minutes: ${summary.totalMinutes}`,
    ];

    // sort tasks by duration in descending order
    summary.tasks.sort((a, b) => b.minutes - a.minutes);

    // create a table of distribution of tasks
    output.push("", "## Distribution of tasks", "");
    const distribution = summary.tasks.map((task) => {
        const percentage = Math.round(task.percentage);

        const taskName = task.task;
        const percentageString = `${percentage}`.padStart(2, " ");
        const bar = "█".repeat(percentage);

        return `  ${taskName} ${percentageString}% | ${bar}`;
    });
    output.push("```", ...distribution, "```");

    // create list of tasks including their duration
    output.push("", "## List of tasks", "");
    const taskList = summary.tasks.map((task) => {
        const hours = task.hours.toFixed(2).padStart(5, " ");
        const minutes = task.minutes.toFixed(0).padStart(4, " ");
        return `  ${task.task} ${hours}h (${minutes} minutes)`;
    });
    output.push("```", ...taskList, "```");

    return output.join("\n") + "\n";
}
