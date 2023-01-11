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
    console.log("App.vue created", data.selectedProject)
    const currentProject = getSelectedProject()
    if (currentProject!=null){
      console.log("selected project in localstorage: ", currentProject)
      this.selectedProject = currentProject
      router.push("/analysis")
    } else {
      console.log("There is no selected project in localstorage")
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
