<script lang="ts">
import {defineComponent, reactive} from 'vue'
import {saveSelectedProject} from "../utils/storage";
import Welcome from "./Welcome.vue";
import ProjectList from "./ProjectList.vue";
import CheckEnvironment from "./CheckEnvironment.vue";
import {CheckEnvironmentPath} from "../../wailsjs/go/main/App";

const data = reactive({
})

export default defineComponent({
  components: {ProjectList, Welcome, CheckEnvironment},
  data() {
    return {
      check: false,
    }
  },
  methods: {
    handleSelectedProject(vm: any, project:any) {
      saveSelectedProject(project)
      vm.$router.push("/analysis")
    }
  },
  async mounted(){
    this.check = await CheckEnvironmentPath()
    console.log("check", this.check)
  }
})
</script>

<template>
  <welcome />
  <check-environment v-if="!check" />
  <project-list v-else @selected="(params: any) => handleSelectedProject(this, params)"/>
</template>

<style scoped>
</style>
