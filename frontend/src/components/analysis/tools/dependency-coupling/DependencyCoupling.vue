<script lang="ts">
import PackageSelection from "./PackageSelection.vue"
import CouplingDetails from "./CouplingDetails.vue"
import {getSelectedProject} from "../../../../utils/storage";
import {project, coupling} from "../../../../../wailsjs/go/models";
import {defineComponent} from "vue";
import {GetDependencyCoupling} from "../../../../../wailsjs/go/main/Api";

export default defineComponent({
  name: "DependencyCouplingTool",
  components: {PackageSelection, CouplingDetails},
  data() {
    return {
      project: undefined as project.ProjectInfo | undefined,
      result: undefined as coupling.DependencyCoupling | undefined
    }
  },
  methods: {
    async runDependencyCoupling({dependency}: any) {
      if(this.project!=undefined){
        console.log("Project", this.project)
        this.result = await GetDependencyCoupling(this.project, dependency)
        console.log(this.result)
      }
    }
  },
  async mounted() {
    this.project = getSelectedProject()
  }
})
</script>

<template>
  <div class="card">
    <div class="card-header">
      <h3 class="card-title">Dependency Coupling Analysis</h3>
      <div class="card-tools">
        <router-link to="/analysis" class="btn btn-sm btn-primary" title="Back to Project Analysis">
          <i class="fa fa-circle-left" /> Back to Project Analysis
        </router-link>
      </div>
    </div>
    <div class="card-body">
      <div>
        <package-selection @run="runDependencyCoupling" />
      </div>
      <div>
        <coupling-details v-if="result !== undefined" :result=result />
      </div>
    </div>
  </div>
</template>

<style scoped>
</style>
