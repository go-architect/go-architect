<script lang="ts">
import {defineComponent, reactive} from 'vue'
import ProjectSummary from "./ProjectSummary.vue";
import Tools from "./AnalysisTools.vue";
import {getSelectedProject, removeSelectedProject} from "../../utils/storage";

const data = reactive({
})

export default defineComponent({
  emits: ["close"],
  components: {ProjectSummary, Tools },
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
      <tools />
    </div>
  </div>
</template>

<style scoped>
</style>
