<script lang="ts">
import {reactive} from 'vue'
import Welcome from "./Welcome.vue";
import ProjectList from "./ProjectList.vue";
import {saveSelectedProject} from "../utils/storage";

const data = reactive({
})

export default {
  components: {ProjectList, Welcome},
  methods: {
    handleSelectedProject(vm: any, project:any) {
      saveSelectedProject(project)
      vm.$router.push("/analysis")
    }
  }
}
</script>

<template>
  <welcome />
  <project-list @selected="(params: any) => handleSelectedProject(this, params)"/>
</template>

<style scoped>
</style>
