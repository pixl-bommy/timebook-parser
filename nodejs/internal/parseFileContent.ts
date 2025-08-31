export function filterFileContent(fileContent: string[]): string[] {
    // Filter out empty lines and lines that are just whitespace
    return fileContent.filter((line) => {
        const trimmedLine = line.trim();
        return (
            trimmedLine.length > 0 &&
            // split by task
            trimmedLine.startsWith("- (")
        );
    });
}

export function parseTaskLine(line: string) {
    const taskMatch = line.match(/- \(([^)]+)\) (.+)/);
    if (!taskMatch || taskMatch?.length < 3) {
        throw new Error(`Invalid task line format: ${line}`);
    }

    // const description = taskMatch[2]!.trim();
    const timeMatch = taskMatch[1]!.match(/([a-zA-Z]{1}) (\d{1,2}:\d{2}) - (\d{1,2}:\d{2})/);
    if (!timeMatch || timeMatch.length < 4) {
        throw new Error(`Invalid time format in task line: ${line}`);
    }

    const task = timeMatch[1]!;
    const start = timeMatch[2]!.split(":").map(Number) as [hours: number, minutes: number];
    const end = timeMatch[3]!.split(":").map(Number) as [hours: number, minutes: number];

    const durationInMinutes = end[0] * 60 + end[1] - (start[0] * 60 + start[1]);

    return {
        task,
        start,
        end,
        durationInMinutes,
    };
}
export type TaskLine = ReturnType<typeof parseTaskLine>;

export function mapTaskLinesToSummary(taskLines: TaskLine[]): TaskSummary {
    const summary: TaskSummary = {
        totalHours: 0,
        totalMinutes: 0,
        tasks: [],
    };

    // summarize minutes grouped by task
    const minutesByTask: Record<string, number> = {};
    for (const entry of taskLines) {
        const task = entry.task?.toLocaleLowerCase();
        if (!task) continue; // skip if task is not defined

        minutesByTask[task] = minutesByTask[task]
            ? minutesByTask[task] + entry.durationInMinutes
            : entry.durationInMinutes;
    }

    // summarize total minutes and calculate total hours
    summary.totalMinutes = Object.values(minutesByTask).reduce((acc, minutes) => acc + minutes, 0);
    summary.totalHours = summary.totalMinutes / 60;

    // map task minutes to full summary entries
    summary.tasks = Object.entries(minutesByTask).map(([task, minutes]) => {
        return {
            task,
            hours: minutes / 60,
            minutes,
            percentage: (minutes / summary.totalMinutes) * 100,
        };
    });

    return summary;
}

export interface TaskSummary {
    totalHours: number;
    totalMinutes: number;
    tasks: TaskSummaryEntry[];
}

export interface TaskSummaryEntry {
    task: string;
    hours: number;
    minutes: number;
    percentage: number;
}
