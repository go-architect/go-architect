<template>
  <div class="card">
    <div class="card-header">
      <h3 class="card-title">Abstractness & Instability Chart</h3>
      <div class="card-tools">
        <router-link to="/analysis" class="btn btn-sm btn-primary" title="Back to Project Analysis">
          <i class="fa fa-circle-left" /> Back to Project Analysis
        </router-link>
      </div>
    </div>
    <div class="card-body">
      <div class="row">
        <div class="col-12 col-md-5">
          <instability-chart :packages=collection />
        </div>
        <div class="col-12 col-md-7">
          <instability-table :packages=collection />
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import InstabilityChart from "./InstabilityChart.vue"
import InstabilityTable from "./InstabilityTable.vue"
import {GetInstability} from "../../../../../wailsjs/go/main/Api";
import {getSelectedProject} from "../../../../utils/storage";
import {defineComponent} from "vue";
import {instability} from "../../../../../wailsjs/go/models";

export default defineComponent({
  name: "InstabilityTool",
  components: {InstabilityChart, InstabilityTable},
  data() {
    return {
      collection: [] as instability.PackageInstability[]
    }
  },
  async mounted() {
    const project = getSelectedProject()
    this.collection = await GetInstability(project)
  }
})
</script>

<style scoped>
</style>
