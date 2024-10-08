export function unixToDateString(unixTimestamp: number): string {
    const date = new Date(unixTimestamp * 1000);
    const day = date.getDate().toString().padStart(2, '0');
    const month = date.toLocaleString('en-US', { month: 'short' });
    const year = date.getFullYear().toString().slice(-2);

    return `${day}-${month}-${year}`;
}