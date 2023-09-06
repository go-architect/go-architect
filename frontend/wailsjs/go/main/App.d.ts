// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {storage} from '../models';
import {context} from '../models';

export function CheckEnvironmentPath():Promise<boolean>;

export function GetGoVersion():Promise<string>;

export function GetProjectList():Promise<storage.ProjectList>;

export function GetSelectedProject():Promise<storage.SelectedProject>;

export function OpenDirectoryDialog():Promise<string>;

export function RemoveSelectedProject():Promise<void>;

export function SaveProjectList(arg1:storage.ProjectList):Promise<void>;

export function SaveSelectedProject(arg1:storage.Project):Promise<void>;

export function SetContext(arg1:context.Context):Promise<void>;
