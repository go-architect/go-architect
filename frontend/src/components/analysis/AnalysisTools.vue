<script lang="ts">
import AnalysisTool from "./AnalysisTool.vue"

export default {
  props: {},
  components: {AnalysisTool},
  data() {
    return {
      tools: [
        {
          title: "Project Metrics",
          description: "This tool displays different metrics as Lines of code, project complexity and interfaces analysis.",
          image: "/images/metrics.png",
          href: "/metrics",
          beta: false,
          soon: false
        },
        {
          title: "Dependencies Graph",
          description: "A dependency graph is a directed graph representing dependencies between packages in the Go Project.",
          image: "/images/dependency-graph.png",
          href: "/dg",
          beta: false,
          soon: false
        },
        {
          title: "Dependency Structure Matrix",
          description: "A DSM is a simple and visual representation of a system in the form of a square matrix, displaying dependency relations between packages.",
          image: "/images/dsm.png",
          href: "/dsm",
          beta: false,
          soon: false
        },
        {
          title: "Instability & Abtractness",
          description: "This tool helps understanding the balance between packages that are open to extension and packages that are easy to modify.",
          image: "/images/abstractness.png",
          href:"/instability",
          beta: false,
          soon: false,
          onclick: () => {
            console.log("open Instability & Abstractness Tool")
          }
        },
        {
          title: "Dependency Coupling",
          description: "This tool is useful when you want to know the coupling level of your whole project related to a specific dependency",
          image: "/images/coupling.png",
          href: "/coupling",
          beta: false,
          soon: false
        },
        {
          title: "Version Control Analysis",
          description: "This tool helps you to gather information about how your team uses your version control system.",
          image: "/images/vcs.png",
          href: "/vcs",
          beta: false,
          soon: false
        },
      ]
    }
  },
  methods: {
  }
}
</script>

<template>
  <div class="card">
    <div class="card-header">
      <h3 class="card-title">Analysis Tools</h3>
      <div class="card-tools">
        <button type="button" class="btn btn-tool" data-card-widget="collapse"><i class="fas fa-minus"></i>
        </button>
      </div>
    </div>
    <div class="card-body">
      <div class="row mt-4">
        <div v-for="tool in tools" class="col-lg-3 col-md-12">
          <analysis-tool :tool=tool />
        </div>
      </div>
    </div>
  </div>

</template>

<style scoped>
</style>
