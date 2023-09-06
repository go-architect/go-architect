<script lang="ts">
import {defineComponent, reactive} from 'vue'
import ProjectSummary from "./ProjectSummary.vue";
import Tools from "./AnalysisTools.vue";
import {getSelectedProject, removeSelectedProject} from "../../utils/storage";
import {project} from "../../../wailsjs/go/models";
import ProjectInfo = project.ProjectInfo;

const data = reactive({
})

export default defineComponent({
  emits: ["close"],
  components: {ProjectSummary, Tools },
  data() {
    return {
      project: null as ProjectInfo | null,
    }
  },
  methods: {
    closeProject: (vm: any) => {
      removeSelectedProject()
      vm.$router.push("/projects")
    },
  },
  async mounted() {
    this.project = await getSelectedProject()
    console.log("mounted", this.project)
  }
})
</script>

<template>
  <div class="row">
    <div class="col-md-12">
      <project-summary v-if="project!=undefined" :project=project @close="() => closeProject(this)" />
    </div>
  </div>
  <div class="row">
    <div class="col-md-12">
      <tools />
    </div>
  </div>
</template>

<style scoped>
</style>
