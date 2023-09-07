import colorLib from "@kurkle/color";

export interface DataSetEntry {
    label: string;
    borderColor: string;
    backgroundColor?: string;
    data: (number|null)[];
}

export const transparentize = (value: string, opacity?: number) => {
    const alpha = opacity === undefined ? 0.5 : 1 - opacity;
    return colorLib(value).alpha(alpha).rgbString();
}

export const COLORS = {
    red: 'rgb(255, 99, 132)',
    orange: 'rgb(255, 159, 64)',
    yellow: 'rgb(255, 205, 86)',
    green: 'rgb(75, 192, 192)',
    blue: 'rgb(54, 162, 235)',
    purple: 'rgb(153, 102, 255)',
    grey: 'rgb(201, 203, 207)',
    pink: 'rgb(252, 71, 189)',
    waterblue: 'rgb(5, 231, 247)'
};

export const resolveIcon = (current: number, previous:number): string => {
    if (current == previous) {
        return "fa-solid fa-circle"
    }
    if (current > previous) {
        return "fa-solid fa-circle-up"
    }
    return "fa-solid fa-circle-down"
}
export const resolveStyle = (current: number, previous:number): string => {
    if (current == previous) {
        return "color: #cccccc;"
    }
    if (current > previous) {
        return "color: #06a741;"
    }
    return "color: #df0707;"
}
export const resolveHint = (current: number, previous:number): string => {
    if (current == previous) {
        return "No changes since last analysis"
    }
    if (current > previous) {
        return `+${current-previous} since last analysis`
    }
    return `-${previous-current} since last analysis`
}
