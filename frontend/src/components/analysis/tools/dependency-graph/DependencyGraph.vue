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
          <div id="dependency-graph"></div>
        </div>
        <div class="col-3">
          <legend-box />
          <div class="card card-primary card-outline">
            <div class="card-header">
              <h3 class="card-title">Package Details</h3>
            </div>
            <div v-if="selectedNode != undefined" class="card-body package-details">
              <ul class="list-group list-group-unbordered mb-3">
                <li class="list-group-item">
                  <b>Selected Package</b>
                  <div>{{ selectedNode.package }}</div>
                </li>
                <li class="list-group-item">
                  <b>Depends on Packages</b>
                  <ul>
                    <li v-for="d in selectedNode.dependencies">{{ d }}</li>
                  </ul>
                </li>
                <li class="list-group-item">
                  <b>Dependant Packages</b>
                  <ul>
                    <li v-for="d in selectedNode.dependants">{{ d }}</li>
                  </ul>
                </li>
              </ul>
            </div>
            <div v-else class="card-body package-details">
              No node selected
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import {defineComponent} from "vue";
import * as vis from "vis-network/standalone";

import {getSelectedProject} from "../../../../utils/storage";
import {dependency, project} from "../../../../../wailsjs/go/models";
import {GetDependencyGraph} from "../../../../../wailsjs/go/main/Api";
import LegendBox from "./LegendBox.vue";

const nodeEdges: any[] = []
const nodesArray: vis.Node[] = []
const edgesArray: vis.Edge[] = []
type SelectedNodeType = {
  package: string,
  dependencies: string[],
  dependants: string[]
}

export default defineComponent( {
  name: "DependencyGraph",
  components: { LegendBox },
  data() {
    return {
      nodes: nodesArray,
      edges: edgesArray,
      selectedNode: undefined as SelectedNodeType | undefined,
      project: null as project.ProjectInfo | null,
      result: undefined as dependency.ModuleDependencyGraph | undefined,
      nodesSize: 0
    }
  },
  methods: {
    mapToNetworkGraph(dg: dependency.ModuleDependencyGraph) {
      this.mapNodes(dg)
      this.mapEdges(dg)
    },
    mapEdges(dg: dependency.ModuleDependencyGraph) {
      for(let origin in dg.relations){
        for(let idx in dg.relations[origin]){
          edgesArray.push({
            from: origin,
            to: dg.relations[origin][idx],
            color: {
              color: "#80d6ff",
              highlight: "#ff0000"
            },
            arrows: {
              to: true
            },
            chosen: {
              label: false,
              edge: true
            }
          })
          const idx2 = nodeEdges.findIndex((ne) => ne.name === origin)
          if (idx2 >= 0) {
            nodeEdges[idx2].count++
          } else {
            nodeEdges.push({ name: origin, count: 1 })
          }
        }
      }
    },
    mapNodes(dg: dependency.ModuleDependencyGraph){
      for(let n in dg.internal){
        nodesArray.push({
          id: dg.internal[n],
          label: dg.internal[n],
          color: "#80d6ff",
          shape: 'box'
        })
      }
      for(let n in dg.organization){
        nodesArray.push({
          id: dg.organization[n],
          label: dg.organization[n],
          color: "#fab57a",
          shape: 'box'
        })
      }
      for(let n in dg.external){
        nodesArray.push({
          id: dg.external[n],
          label: dg.external[n],
          color: "#f06868",
          shape: 'box'
        })
      }
      for(let n in dg.standard){
        nodesArray.push({
          id: dg.standard[n],
          label: dg.standard[n],
          color: "#edf798",
          shape: 'box'
        })
      }
      this.nodesSize = dg.internal.length + dg.external.length + dg.standard.length
      if(dg.organization != null) this.nodesSize += dg.organization.length
    }
  },
  async mounted() {
    this.project = await getSelectedProject()
    if(this.project !== null){
      const dg = await GetDependencyGraph(this.project!)
      this.mapToNetworkGraph(dg)
    }

    const graphData: vis.Data = {
      nodes: nodesArray,
      edges: edgesArray
    };

    const options: vis.Options = {
      manipulation: false,
      height: '100%',
      width: '100%',
      layout: {
        hierarchical: {
          enabled: true,
          levelSeparation: 80
        }
      },
      physics: {
        hierarchicalRepulsion: {
          nodeDistance: 300,
          avoidOverlap: 1
        }
      },
      interaction: {
        navigationButtons: true,
        zoomView: false
      },
    }
    const container = document.getElementById('dependency-graph');
    const network = new vis.Network(container!, graphData, options);

    network.on("select", (params) => {
      function getDependencies(edgesArray: vis.Edge[], node: string) {
        const dependencies: string[] = []
        edgesArray.forEach(edge => {
          if (edge.from == node && edge.to != undefined) {
            dependencies.push(edge.to+"")
          }
        })
        return dependencies
      }
      function getDependants(edgesArray: vis.Edge[], node: string) {
        const dependants: string[] = []
        edgesArray.forEach(edge => {
          if (edge.to == node && edge.from != undefined) {
            dependants.push(edge.from+"")
          }
        })
        return dependants
      }

      if (params.nodes.length === 0) {
        this.selectedNode = undefined
      } else {
        this.selectedNode = {
          package: params.nodes[0],
          dependencies: getDependencies(edgesArray, params.nodes[0]),
          dependants: getDependants(edgesArray, params.nodes[0])
        }

      }
    });
  }
})
</script>

<style scoped>
#dependency-graph {
  width: 100%;
  height: 600px;
  border: 1px solid black;
  border-radius: 5px;
}

.package-details {
  font-size: 12px;
}
</style>
