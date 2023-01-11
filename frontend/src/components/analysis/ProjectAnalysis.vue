<script lang="ts">
import {defineComponent, reactive} from 'vue'
import ProjectSummary from "./ProjectSummary.vue";
import Diagrams from "./AnalysisTools.vue";
import ProjectMetrics from "./ProjectMetrics.vue";
import {getSelectedProject, removeSelectedProject} from "../../utils/storage";
import {useRouter} from "vue-router";
import {router} from "../../router";

const data = reactive({
})

export default defineComponent({
  emits: ["close"],
  components: {ProjectSummary, Diagrams, ProjectMetrics },
  data() {
    return {
      project: null,
    }
  },
  methods: {
    closeProject: (vm: any) => {
      removeSelectedProject()
      vm.$router.push("/projects")
    },
  },
  mounted(){
    this.project = getSelectedProject()
    console.log("mounted", this.project)
  }
})
</script>

<template>
  <div class="row">
    <div class="col-md-12">
      <project-summary v-if="project!=null" :project=project @close="() => closeProject(this)" />
    </div>
  </div>
  <div class="row">
    <div class="col-md-12">
      <diagrams />
    </div>
  </div>
  <div class="row">
    <div class="col-md-12">
      <project-metrics v-if="project!=null" :project=project />
    </div>
  </div>
</template>

<style scoped>
</style>
