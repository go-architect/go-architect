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
          <div class="col-12">
            <date-range-selection @run="updateDateRange" />
          </div>
        </div>
        <div v-if="vcs.AuthorsInfo.TotalCommits === 0" class="row">
          <div class="col-12">
            No changes since {{ months }} month{{ months>1 ? 's':'' }} ago
          </div>
        </div>
        <div v-else class="row">
          <div class="col-9">
            <changes-summary :data="getPathChangesSummaryData" :contributors="vcs.AuthorsInfo.AuthorDetails.length" :colors="getAssignedColorsMap" />
            <changes-summary :data="getFileChangesSummaryData" :contributors="vcs.AuthorsInfo.AuthorDetails.length" :colors="getAssignedColorsMap" />
          </div>
          <div class="col-3">
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
import DateRangeSelection from "./DateRangeSelection.vue";

export default defineComponent({
  name: "VCSAnalysis",
  components: {Contributors,ChangesSummary,DateRangeSelection},
  data() {
    return {
      months: 6,
      vcs: undefined as any,
    }
  },
  computed: {
    getPathChangesSummaryData(){
      return {
        title: "Changes Summary by Path",
        kind: "Path",
        modifications: this.vcs.ModificationsInfo.PathModifications
      }
    },
    getFileChangesSummaryData(){
      return {
        title: "Changes Summary by File",
        kind: "File",
        modifications: this.vcs.ModificationsInfo.FileModifications
      }
    },
    getAssignedColorsMap(){
      return assignColors(this.vcs.AuthorsInfo.AuthorDetails)

    }
  },
  methods: {
    async updateDateRange({range}: any) {
      const project = getSelectedProject()
      this.months = range
      this.vcs = await GetVCSAnalysisInfo(project, this.months)
    }
  },
  async mounted() {
    const project = getSelectedProject()
    this.vcs = await GetVCSAnalysisInfo(project, this.months)
  }
})
</script>

<style scoped>
</style>
