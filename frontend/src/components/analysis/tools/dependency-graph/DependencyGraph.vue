<template>
  <div class="card">
    <div class="card-header">
      <h3 class="card-title">Dependency Graph</h3>
      <div class="card-tools">
        <router-link to="/analysis" class="btn btn-sm btn-primary" title="Back to Project Analysis">
          <i class="fa fa-circle-left" /> Back to Project Analysis
        </router-link>
      </div>
    </div>
    <div class="card-body">
      <div class="row">
        <div class="col-9">
          <v-network-graph
              class="graph"
              :nodes="nodes"
              :edges="edges"
              :layouts="layouts"
              :configs="configs"
          />
        </div>
        <div class="col-3">
          <legend-box />
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import {defineComponent, ref} from "vue";
import { Node, Edge, Layouts, VNetworkGraphInstance } from "v-network-graph";
//@ts-ignore
import dagre from "dagre/dist/dagre.min.js"

import {getSelectedProject} from "../../../../utils/storage";
import {dependency, project} from "../../../../../wailsjs/go/models";
import {GetDependencyGraph} from "../../../../../wailsjs/go/main/Api";
import LegendBox from "./LegendBox.vue";

const nodes: { [key: string]: Node } = {}
const edges: { [key: string]: Edge } = {}
const layouts: Layouts = { nodes: {}}

const graph = ref<VNetworkGraphInstance>()

export default defineComponent( {
  name: "DependencyGraph",
  components: { LegendBox },
  data() {
    return {
      nodes: nodes,
      edges: edges,
      layouts: layouts,
      configs: {
        node: {
          normal: {
            type: "circle",
            radius: (node: any) => node.size, // Use the value of each node object
            color: (node: any) => node.color,
          },
          hover: {
            radius: (node: any) => node.size + 2,
            color: (node: any) => node.color,
          },
          selectable: true,
          label: {
            visible: true,
            fontFamily: undefined,
            fontSize: 11,
            lineHeight: 1.1,
            color: "#000000",
            margin: 4,
            direction: "south",
            text: "name",
          },
          focusring: {
            color: "darkgray",
          },
        },
        edge: {
          marker: {
            source: {
              type: "none",
            },
            target: {
              type: "arrow"
            }
          }
        }
      },
      project: undefined as project.ProjectInfo | undefined,
      result: undefined as dependency.ModuleDependencyGraph | undefined,
      nodesSize: 0
    }
  },
  methods: {
    mapToNetworkGraph(dg: dependency.ModuleDependencyGraph) {
      this.mapNodes(dg)
      this.mapEdges(dg)
      this.layout("TB")
    },
    mapEdges(dg: dependency.ModuleDependencyGraph) {
      for(let origin in dg.relations){
        for(let idx in dg.relations[origin]){
          this.edges[origin+"_"+dg.relations[origin][idx]] = { source: origin, target: dg.relations[origin][idx] }
        }
      }
    },
    mapNodes(dg: dependency.ModuleDependencyGraph){
      for(let n in dg.internal){
        this.nodes[dg.internal[n]] = { name: dg.internal[n], size: 8, color: "#80d6ff" }
      }
      for(let n in dg.external){
        this.nodes[dg.external[n]] = { name: dg.external[n], size: 8, color: "#f06868" }
      }
      for(let n in dg.standard){
        this.nodes[dg.standard[n]] = { name: dg.standard[n], size: 8, color: "#edf798" }
      }
      for(let n in dg.organization){
        this.nodes[dg.organization[n]] = { name: dg.organization[n], size: 8, color: "#fab57a" }
      }
      this.nodesSize = dg.internal.length + dg.external.length + dg.standard.length
      if(dg.organization != null) this.nodesSize += dg.organization.length
    },
    layout(direction: "TB" | "LR") {
      if (Object.keys(this.nodes).length <= 1 || Object.keys(this.edges).length == 0) {
        return
      }

      // convert graph
      // ref: https://github.com/dagrejs/dagre/wiki
      const g = new dagre.graphlib.Graph()
      // Set an object for the graph label
      g.setGraph({
        rankdir: direction,
        nodesep: this.nodesSize * 3,
        edgesep: this.nodesSize,
        ranksep: this.nodesSize * 2,
      })
      // Default to assigning a new object as a label for each new edge.
      g.setDefaultEdgeLabel(() => ({}))

      // Add nodes to the graph. The first argument is the node id. The second is
      // metadata about the node. In this case we're going to add labels to each of
      // our nodes.
      Object.entries(this.nodes).forEach(([nodeId, node]) => {
        g.setNode(nodeId, { label: node.name, width: this.nodesSize, height: this.nodesSize })
      })

      // Add edges to the graph.
      Object.values(this.edges).forEach(edge => {
        g.setEdge(edge.source, edge.target)
      })

      dagre.layout(g)

      const box: Record<string, number | undefined> = {}
      g.nodes().forEach((nodeId: string) => {
        // update node position
        const x = g.node(nodeId).x
        const y = g.node(nodeId).y
        this.layouts.nodes[nodeId] = { x, y }

        // calculate bounding box size
        box.top = box.top ? Math.min(box.top, y) : y
        box.bottom = box.bottom ? Math.max(box.bottom, y) : y
        box.left = box.left ? Math.min(box.left, x) : x
        box.right = box.right ? Math.max(box.right, x) : x
      })

      const graphMargin = this.nodesSize * 2
      const viewBox = {
        top: (box.top ?? 0) - graphMargin,
        bottom: (box.bottom ?? 0) + graphMargin,
        left: (box.left ?? 0) - graphMargin,
        right: (box.right ?? 0) + graphMargin,
      }
      graph.value?.setViewBox(viewBox)
    }
  },
  async mounted() {
    this.nodes = nodes
    this.edges = edges
    this.layouts = layouts
    this.nodesSize = 0

    this.project = getSelectedProject()
    if(this.project !== undefined){
      const dg = await GetDependencyGraph(this.project)
      this.mapToNetworkGraph(dg)
    }
  }
})
</script>

<style scoped>
.graph {
  width: 100%;
  height: 600px;
  border: 1px solid black;
  border-radius: 5px;
}
</style>
