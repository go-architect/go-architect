<template>
  <div class="card">
    <div class="card-header">
      <h3 class="card-title">VCS Analysis Tool</h3>
      <div class="card-tools">
        <router-link to="/analysis" class="btn btn-sm btn-primary" title="Back to Project Analysis">
          <i class="fa fa-circle-left" /> Back to Project Analysis
        </router-link>
      </div>
    </div>
    <div class="card-body">
      <div v-if="vcs === undefined" class="overlay-wrapper" style="height: 300px;">
        <div class="overlay">
          <i class="fas fa-3x fa-sync-alt fa-spin"></i>
          <div class="text-bold pt-2">Loading Information from VCS...</div>
        </div>
      </div>
      <div v-else>
        <div class="row">
          <div class="col-8">
            <changes-summary :data="getPathChangesSummaryData" :contributors="vcs.AuthorsInfo.AuthorDetails.length" :colors="getAssignedColorsMap" />
            <changes-summary :data="getFileChangesSummaryData" :contributors="vcs.AuthorsInfo.AuthorDetails.length" :colors="getAssignedColorsMap" />
          </div>
          <div class="col-4">
            <contributors :data="vcs.AuthorsInfo" :colors="getAssignedColorsMap" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import {defineComponent} from "vue";
import {getSelectedProject} from "../../../../utils/storage";
import {GetVCSAnalysisInfo} from "../../../../../wailsjs/go/main/Api";

import Contributors from "./Contributors.vue"
import ChangesSummary from "./ChangesSummary.vue"
import {assignColors} from "./helpers";

export default defineComponent({
  name: "VCSAnalysis",
  components: {Contributors,ChangesSummary},
  data() {
    return {
      vcs: undefined as any,
    }
  },
  computed: {
    getPathChangesSummaryData(){
      return {
        title: "Changes Summary by Path",
        modifications: this.vcs.ModificationsInfo.PathModifications
      }
    },
    getFileChangesSummaryData(){
      return {
        title: "Changes Summary by File",
        modifications: this.vcs.ModificationsInfo.FileModifications
      }
    },
    getAssignedColorsMap(){
      return assignColors(this.vcs.AuthorsInfo.AuthorDetails)

    }
  },
  async mounted() {
    const project = getSelectedProject()
    const result = await GetVCSAnalysisInfo(project, 6)
    console.log(result)
    this.vcs = result
  }
})
</script>

<style scoped>
.vcs {
  width: unset;
  font-size: 8px;
}
</style>
