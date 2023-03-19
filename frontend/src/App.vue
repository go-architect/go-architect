<script lang="ts">
import {reactive, defineComponent} from "vue";
import { useRouter } from 'vue-router'
import ProjectSelection from './components/ProjectSelection.vue'
import ProjectAnalysis from './components/analysis/ProjectAnalysis.vue'
import { getSelectedProject } from './utils/storage'

const data = reactive({
  selectedProject: "",
})

export default defineComponent({
  components: {ProjectSelection, ProjectAnalysis},
  data() {
    return {
      selectedProject: null,
    }
  },
  created() {
    const router = useRouter()
    const currentProject = getSelectedProject()
    if (currentProject!=null){
      this.selectedProject = currentProject
      router.push("/analysis")
    } else {
      router.push("/projects")
    }
  }
})
</script>

<template>
  <router-view />
</template>

<style>
</style>
