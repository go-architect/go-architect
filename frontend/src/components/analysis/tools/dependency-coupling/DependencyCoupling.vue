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
      loading: false,
      project: undefined as project.ProjectInfo | undefined,
      result: undefined as coupling.DependencyCoupling | undefined
    }
  },
  methods: {
    async runDependencyCoupling({dependency}: any) {
      this.result = undefined
      this.loading = true
      if(this.project!=undefined){
        this.result = await GetDependencyCoupling(this.project, dependency)
        this.loading = false
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
        <div v-if="loading" class="overlay">
          <i class="fas fa-3x fa-sync-alt fa-spin"></i>
          <div class="text-bold pt-2">Analyzing coupling...</div>
        </div>

        <coupling-details v-if="result !== undefined" :result=result />
      </div>
    </div>
  </div>
</template>

<style scoped>
</style>
