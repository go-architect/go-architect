export namespace comments {

	export class CommentsMetrics {
	    total_lines: number;
	    ratio: number;
	    files_with_comments: number;
	    files_with_comments_ratio: number;

	    static createFrom(source: any = {}) {
	        return new CommentsMetrics(source);
	    }

	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.total_lines = source["total_lines"];
	        this.ratio = source["ratio"];
	        this.files_with_comments = source["files_with_comments"];
	        this.files_with_comments_ratio = source["files_with_comments_ratio"];
	    }
	}

}

export namespace coupling {

	export class Detail {
	    line: number;
	    details: string;

	    static createFrom(source: any = {}) {
	        return new Detail(source);
	    }

	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.line = source["line"];
	        this.details = source["details"];
	    }
	}
	export class FileCoupling {
	    package: string;
	    file: string;
	    coupling_level: number;
	    details: Detail[];

	    static createFrom(source: any = {}) {
	        return new FileCoupling(source);
	    }

	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.package = source["package"];
	        this.file = source["file"];
	        this.coupling_level = source["coupling_level"];
	        this.details = this.convertValues(source["details"], Detail);
	    }

		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class PackageCoupling {
	    package: string;
	    coupling_level: number;
	    details: FileCoupling[];

	    static createFrom(source: any = {}) {
	        return new PackageCoupling(source);
	    }

	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.package = source["package"];
	        this.coupling_level = source["coupling_level"];
	        this.details = this.convertValues(source["details"], FileCoupling);
	    }

		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class DependencyCoupling {
	    dependency: string;
	    coupling_level: number;
	    details: PackageCoupling[];

	    static createFrom(source: any = {}) {
	        return new DependencyCoupling(source);
	    }

	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.dependency = source["dependency"];
	        this.coupling_level = source["coupling_level"];
	        this.details = this.convertValues(source["details"], PackageCoupling);
	    }

		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace dependency {

	export class ModuleDependencyGraph {
	    module: string;
	    internal: string[];
	    external: string[];
	    organization: string[];
	    standard: string[];
	    relations: {[key: string]: string[]};

	    static createFrom(source: any = {}) {
	        return new ModuleDependencyGraph(source);
	    }

	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.module = source["module"];
	        this.internal = source["internal"];
	        this.external = source["external"];
	        this.organization = source["organization"];
	        this.standard = source["standard"];
	        this.relations = source["relations"];
	    }
	}

}

export namespace dsm {

	export class DependencyStructureMatrix {
	    module: string;
	    packages: string[];
	    dependencies: number[][];

	    static createFrom(source: any = {}) {
	        return new DependencyStructureMatrix(source);
	    }

	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.module = source["module"];
	        this.packages = source["packages"];
	        this.dependencies = source["dependencies"];
	    }
	}

}

export namespace gocyclo {

	export class GoCycloOutput {
	    average_complexity: number;

	    static createFrom(source: any = {}) {
	        return new GoCycloOutput(source);
	    }

	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.average_complexity = source["average_complexity"];
	    }
	}

}

export namespace instability {

	export class PackageInstability {
	    package_name: string;
	    abstractions_count: number;
	    implementations_count: number;
	    afferent_coupling: number;
	    efferent_coupling: number;
	    instability: number;
	    abstractness: number;
	    distance: number;

	    static createFrom(source: any = {}) {
	        return new PackageInstability(source);
	    }

	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.package_name = source["package_name"];
	        this.abstractions_count = source["abstractions_count"];
	        this.implementations_count = source["implementations_count"];
	        this.afferent_coupling = source["afferent_coupling"];
	        this.efferent_coupling = source["efferent_coupling"];
	        this.instability = source["instability"];
	        this.abstractness = source["abstractness"];
	        this.distance = source["distance"];
	    }
	}

}

export namespace interfaces {

	export class InterfaceInfo {
	    package: string;
	    file: string;
	    name: string;
	    methods: number;

	    static createFrom(source: any = {}) {
	        return new InterfaceInfo(source);
	    }

	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.package = source["package"];
	        this.file = source["file"];
	        this.name = source["name"];
	        this.methods = source["methods"];
	    }
	}
	export class InterfaceMetrics {
	    average_methods: number;
	    max_methods: InterfaceInfo[];
	    min_methods: InterfaceInfo[];

	    static createFrom(source: any = {}) {
	        return new InterfaceMetrics(source);
	    }

	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.average_methods = source["average_methods"];
	        this.max_methods = this.convertValues(source["max_methods"], InterfaceInfo);
	        this.min_methods = this.convertValues(source["min_methods"], InterfaceInfo);
	    }

		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace loc {

	export class FileLOC {
	    package: string;
	    file: string;
	    loc: number;

	    static createFrom(source: any = {}) {
	        return new FileLOC(source);
	    }

	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.package = source["package"];
	        this.file = source["file"];
	        this.loc = source["loc"];
	    }
	}
	export class PackageLOC {
	    package: string;
	    loc: number;

	    static createFrom(source: any = {}) {
	        return new PackageLOC(source);
	    }

	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.package = source["package"];
	        this.loc = source["loc"];
	    }
	}
	export class ProjectLOC {
	    total: number;
	    packages: PackageLOC[];
	    files: FileLOC[];

	    static createFrom(source: any = {}) {
	        return new ProjectLOC(source);
	    }

	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.total = source["total"];
	        this.packages = this.convertValues(source["packages"], PackageLOC);
	        this.files = this.convertValues(source["files"], FileLOC);
	    }

		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace project {

	export class ProjectInfo {
	    name: string;
	    path: string;
	    package: string;
	    organization_packages: string[];

	    static createFrom(source: any = {}) {
	        return new ProjectInfo(source);
	    }

	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.path = source["path"];
	        this.package = source["package"];
	        this.organization_packages = source["organization_packages"];
	    }
	}

}

export namespace repository {

	export class RepositoryInfo {
	    path: string;
	    url: string;
	    branch: string;
	    revision: string;
	    is_up_to_date: boolean;

	    static createFrom(source: any = {}) {
	        return new RepositoryInfo(source);
	    }

	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.path = source["path"];
	        this.url = source["url"];
	        this.branch = source["branch"];
	        this.revision = source["revision"];
	        this.is_up_to_date = source["is_up_to_date"];
	    }
	}

}

export namespace storage {

	export class Metrics {
	    source_files: number;
	    structs: number;
	    interfaces: number;
	    functions: number;
	    methods: number;
	    variables: number;
	    constants: number;
	    public_structs: number;
	    public_interfaces: number;
	    public_functions: number;
	    public_methods: number;
	    public_variables: number;
	    public_constants: number;
	    num_packages: number;
	    total_loc: number;
	    interface_avg_size: number;
	    interface_max_size: number;
	    interface_min_size: number;
	    comments_total_lines: number;
	    comments_lines_ratio: number;
	    files_with_comments: number;
	    files_with_comments_ratio: number;
	    cyclomatic_complexity_avg: number;
	    cognitive_complexity_avg: number;

	    static createFrom(source: any = {}) {
	        return new Metrics(source);
	    }

	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.source_files = source["source_files"];
	        this.structs = source["structs"];
	        this.interfaces = source["interfaces"];
	        this.functions = source["functions"];
	        this.methods = source["methods"];
	        this.variables = source["variables"];
	        this.constants = source["constants"];
	        this.public_structs = source["public_structs"];
	        this.public_interfaces = source["public_interfaces"];
	        this.public_functions = source["public_functions"];
	        this.public_methods = source["public_methods"];
	        this.public_variables = source["public_variables"];
	        this.public_constants = source["public_constants"];
	        this.num_packages = source["num_packages"];
	        this.total_loc = source["total_loc"];
	        this.interface_avg_size = source["interface_avg_size"];
	        this.interface_max_size = source["interface_max_size"];
	        this.interface_min_size = source["interface_min_size"];
	        this.comments_total_lines = source["comments_total_lines"];
	        this.comments_lines_ratio = source["comments_lines_ratio"];
	        this.files_with_comments = source["files_with_comments"];
	        this.files_with_comments_ratio = source["files_with_comments_ratio"];
	        this.cyclomatic_complexity_avg = source["cyclomatic_complexity_avg"];
	        this.cognitive_complexity_avg = source["cognitive_complexity_avg"];
	    }
	}
	export class HistoricalMetrics {
	    date: string;
	    commit: string;
	    metrics: Metrics;

	    static createFrom(source: any = {}) {
	        return new HistoricalMetrics(source);
	    }

	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.date = source["date"];
	        this.commit = source["commit"];
	        this.metrics = this.convertValues(source["metrics"], Metrics);
	    }

		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

	export class Project {
	    id: string;
	    name: string;
	    path: string;
	    package: string;
	    organization_packages: string[];
	    historical_metrics: HistoricalMetrics[];

	    static createFrom(source: any = {}) {
	        return new Project(source);
	    }

	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.path = source["path"];
	        this.package = source["package"];
	        this.organization_packages = source["organization_packages"];
	        this.historical_metrics = this.convertValues(source["historical_metrics"], HistoricalMetrics);
	    }

		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ProjectList {
	    projects: Project[];

	    static createFrom(source: any = {}) {
	        return new ProjectList(source);
	    }

	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.projects = this.convertValues(source["projects"], Project);
	    }

		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace types {

	export class PackageTypes {
	    source_files: number;
	    structs: number;
	    interfaces: number;
	    functions: number;
	    methods: number;
	    variables: number;
	    constants: number;
	    public_structs: number;
	    public_interfaces: number;
	    public_functions: number;
	    public_methods: number;
	    public_variables: number;
	    public_constants: number;
	    package: string;

	    static createFrom(source: any = {}) {
	        return new PackageTypes(source);
	    }

	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.source_files = source["source_files"];
	        this.structs = source["structs"];
	        this.interfaces = source["interfaces"];
	        this.functions = source["functions"];
	        this.methods = source["methods"];
	        this.variables = source["variables"];
	        this.constants = source["constants"];
	        this.public_structs = source["public_structs"];
	        this.public_interfaces = source["public_interfaces"];
	        this.public_functions = source["public_functions"];
	        this.public_methods = source["public_methods"];
	        this.public_variables = source["public_variables"];
	        this.public_constants = source["public_constants"];
	        this.package = source["package"];
	    }
	}
	export class ProjectTypes {
	    source_files: number;
	    structs: number;
	    interfaces: number;
	    functions: number;
	    methods: number;
	    variables: number;
	    constants: number;
	    public_structs: number;
	    public_interfaces: number;
	    public_functions: number;
	    public_methods: number;
	    public_variables: number;
	    public_constants: number;
	    project_package: string;
	    packages: number;
	    details: PackageTypes[];

	    static createFrom(source: any = {}) {
	        return new ProjectTypes(source);
	    }

	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.source_files = source["source_files"];
	        this.structs = source["structs"];
	        this.interfaces = source["interfaces"];
	        this.functions = source["functions"];
	        this.methods = source["methods"];
	        this.variables = source["variables"];
	        this.constants = source["constants"];
	        this.public_structs = source["public_structs"];
	        this.public_interfaces = source["public_interfaces"];
	        this.public_functions = source["public_functions"];
	        this.public_methods = source["public_methods"];
	        this.public_variables = source["public_variables"];
	        this.public_constants = source["public_constants"];
	        this.project_package = source["project_package"];
	        this.packages = source["packages"];
	        this.details = this.convertValues(source["details"], PackageTypes);
	    }

		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}
