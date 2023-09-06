<script lang="ts">
import {reactive, defineComponent} from "vue";
import { useRouter } from 'vue-router'
import ProjectSelection from './components/ProjectSelection.vue'
import ProjectAnalysis from './components/analysis/ProjectAnalysis.vue'
import { getSelectedProject } from './utils/storage'
import {project} from "../wailsjs/go/models";
import ProjectInfo = project.ProjectInfo;

const data = reactive({
  selectedProject: "",
})

export default defineComponent({
  components: {ProjectSelection, ProjectAnalysis},
  data() {
    return {
      selectedProject: null as ProjectInfo | null,
    }
  },
  async created() {
    const router = useRouter()
    const currentProject = await getSelectedProject()
    if (currentProject != null) {
      this.selectedProject = currentProject
      await router.push("/analysis")
    } else {
      await router.push("/projects")
    }
  }
})
</script>

<template>
  <router-view />
</template>

<style>
</style>
