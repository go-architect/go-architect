import { createRouter, createWebHashHistory } from 'vue-router'
import ProjectSelection from './components/ProjectSelection.vue'
import ProjectAnalysis from './components/analysis/ProjectAnalysis.vue'
import DependencyStructureMatrix from './components/analysis/tools/dsm/DependencyStructureMatrix.vue'
import DependencyGraph from './components/analysis/tools/dependency-graph/DependencyGraph.vue'
import DependencyCoupling from './components/analysis/tools/dependency-coupling/DependencyCoupling.vue'
import InstabilityTool from './components/analysis/tools/instability/InstabilityTool.vue'

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
            path: '/dc'
        },
        {
            name: 'InstabilityTool',
            component: InstabilityTool,
            path: '/it'
        }
    ]
})
