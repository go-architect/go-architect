
export const resolvePackageLayer = (entryName: string, mainPackage: string, orgPackages: string[]): string => {
    if (entryName === "Internal Packages") return "internal"
    if (entryName === "Organization Packages") return "organization"
    if (entryName === "External Packages") return "external"
    if (entryName === "StandardLib Packages") return "standard"

    if (entryName.startsWith(mainPackage)) return "internal"
    if (isOrganizationPackage(entryName, orgPackages)) return "organization"
    if (entryName.startsWith("golang.org/x")) return "standard"
    if (!entryName.includes(".")) return "standard"

    return "external"
}

export const resolvePackageLayerIndex = (layer: string): number => {
    if (layer === "internal") return 0
    if (layer === "organization") return 1
    if (layer === "external") return 2
    if (layer === "standard") return 3

    return 0
}

const isOrganizationPackage = (packageName: string, orgPackages: string[]): boolean => {
    return orgPackages.some(op => packageName.startsWith(op))
}