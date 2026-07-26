export function padTime(time: number): string {
    if (time >= 10) return time.toString();
    return "0" + time.toString()
}
