export type Project = {
    id: string
    name: string
    package: string
    path: string
}

export type InstabilityInfo = {
    package: string
    afferentCoupling: number
    efferentCoupling: number
    abstractness: number
    instability: number
    distanceFromDiagonal: number
}
