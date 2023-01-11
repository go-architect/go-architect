<script lang="ts">
import AnalysisTool from "./AnalysisTool.vue"

export default {
  props: {},
  components: {AnalysisTool},
  data() {
    return {
      tools: [
        {
          title: "Dependencies Graph",
          description: "",
          image: "/images/dependency-graph.png",
//          image: "/src/components/analysis/assets/dependency-graph.png",
          href: "/dg",
          beta: false,
          soon: false
        },
        {
          title: "Dependency Structure Matrix",
          description: "",
          image: "/images/dsm.png",
          href: "/dsm",
          beta: false,
          soon: false
        },
        {
          title: "Instability & Abtractness",
          description: "",
          image: "/images/abstractness.png",
          href:"/it",
          beta: false,
          soon: false,
          onclick: () => {
            console.log("open Instability & Abstractness Tool")
          }
        },
        {
          title: "Dependency Coupling",
          description: "This tool is useful when you want to know the coupling level of your whole project related to a specific dependency",
          image: "/images/abstractness.png",
          href: "/dc",
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
