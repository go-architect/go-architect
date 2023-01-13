export const assignColors = (authors: any[]) => {
    const colors = ["#00FF00",
        "#0000FF",
        "#808000",
        "#00FFFF",
        "#FF00FF",
        "#C0C0C0",
        "#808080",
        "#800000",
        "#008000",
        "#FFFF00",
        "#800080",
        "#008080",
        "#FF0000",
        "#FFA500",
        "#000080",
        "#C0C0C0",
        "#FFEBCD",
        "#8FBC8F",
        "#7B68EE",
        "#BDB76B",
        "#93FBD7",
        "#043826",
        "#E3B0A5",
        "#715976",
        "#3F5C9A",
        "#B3BE86",
        "#684D03",
        "#4D78C3"
    ]
    const colorsMap = new Map();
    authors.sort((a, b) => {
        return b.ModifiedLines - a.ModifiedLines
    })
    authors.forEach((v, i) => {
        colorsMap.set(v.Name, colors[i])
    })
    return colorsMap
}


/*
Silver	C0C0C0	192,192,192
White	FFFFFF	255,255,255
 */
