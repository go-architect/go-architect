import { createRouter, createWebHashHistory } from 'vue-router'
import ProjectSelection from './components/ProjectSelection.vue'
import ProjectAnalysis from './components/analysis/ProjectAnalysis.vue'
import DependencyStructureMatrix from './components/analysis/tools/dsm/DependencyStructureMatrix.vue'
import DependencyGraph from './components/analysis/tools/dependency-graph/DependencyGraph.vue'
import DependencyCoupling from './components/analysis/tools/dependency-coupling/DependencyCoupling.vue'
import InstabilityTool from './components/analysis/tools/instability/InstabilityTool.vue'
import ProjectMetrics from './components/analysis/tools/metrics/ProjectMetrics.vue'
import VCSAnalysisTool from './components/analysis/tools/vcs/VersionControlSystemTool.vue'

export const router = createRouter({
    history: createWebHashHistory(),
    routes: [
        {
            name: 'ProjectSelection',
            component: ProjectSelection,
            path: '/'
        },
        {
            name: 'ProjectSelection',
            component: ProjectSelection,
            path: '/projects'
        },
        {
            name: 'ProjectAnalysis',
            component: ProjectAnalysis,
            path: '/analysis'
        },
        {
            name: 'ProjectMetrics',
            component: ProjectMetrics,
            path: '/metrics'
        },
        {
            name: 'DependencyStructureMatrix',
            component: DependencyStructureMatrix,
            path: '/dsm'
        },
        {
            name: 'DependencyGraph',
            component: DependencyGraph,
            path: '/dg'
        },
        {
            name: 'DependencyCoupling',
            component: DependencyCoupling,
            path: '/coupling'
        },
        {
            name: 'InstabilityTool',
            component: InstabilityTool,
            path: '/instability'
        },
        {
            name: 'VCSAnalysisTool',
            component: VCSAnalysisTool,
            path: '/vcs'
        }
    ]
})
