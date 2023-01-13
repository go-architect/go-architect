// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {project} from '../models';
import {dsm} from '../models';
import {coupling} from '../models';
import {dependency} from '../models';
import {instability} from '../models';
import {gocyclo} from '../models';
import {comments} from '../models';
import {interfaces} from '../models';
import {loc} from '../models';
import {types} from '../models';
import {repository} from '../models';
import {git} from '../models';
import {context} from '../models';

export function GetDSM(arg1:project.ProjectInfo):Promise<dsm.DependencyStructureMatrix>;

export function GetDependencyCoupling(arg1:project.ProjectInfo,arg2:string):Promise<coupling.DependencyCoupling>;

export function GetDependencyGraph(arg1:project.ProjectInfo):Promise<dependency.ModuleDependencyGraph>;

export function GetInstability(arg1:project.ProjectInfo):Promise<Array<instability.PackageInstability>>;

export function GetMetricsCognitiveComplexity(arg1:project.ProjectInfo):Promise<gocyclo.GoCycloOutput>;

export function GetMetricsComments(arg1:project.ProjectInfo):Promise<comments.CommentsMetrics>;

export function GetMetricsComplexity(arg1:project.ProjectInfo):Promise<gocyclo.GoCycloOutput>;

export function GetMetricsInterfaces(arg1:project.ProjectInfo):Promise<interfaces.InterfaceMetrics>;

export function GetMetricsLOC(arg1:project.ProjectInfo):Promise<loc.ProjectLOC>;

export function GetMetricsTypes(arg1:project.ProjectInfo):Promise<types.ProjectTypes>;

export function GetProjectInfo(arg1:string):Promise<project.ProjectInfo>;

export function GetRepositoryInfo(arg1:string):Promise<repository.RepositoryInfo>;

export function GetVCSAnalysisInfo(arg1:project.ProjectInfo,arg2:number):Promise<git.VCSAnalysisInfo>;

export function SetContext(arg1:context.Context):Promise<void>;
