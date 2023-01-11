<script lang="ts">

import {reactive} from 'vue'

const data = reactive({
  title: "",
})
export default {
  emits: ["delete","selected"],
  props: {
    project: {
      type: Object
    },
    callback: {
      type: Function,
      default: () => {}
    }
  },
  methods: {
    deleteProject: (vm: any, id: string) => {
      vm.$emit("delete", { id: id })
    },
    selectProject(vm: any, project: string) {
      console.log("ProjectBox: Select Project")
      vm.$emit("selected", project)
    }
  }
}
</script>

<template>
  <div class="col-lg-3 col-6">
    <div class="small-box bg-gray">
      <div class="inner" v-on:click="selectProject(this, project)">
        <div class="project-name">{{ project.name }}</div>
        <div class="project-package"><i class="fa fa-boxes-stacked" title="Main Package"></i> {{ project.package }}</div>
        <div class="project-path"><i class="fa fa-folder" title="Project Path"></i> {{ project.path }}</div>
      </div>
      <div class="small-box-footer" v-on:click="deleteProject(this, project.id)">Delete Project <i class="fas fa-trash"></i></div>
    </div>
  </div>
</template>

<style scoped>
.inner {
  text-align: left;
  padding: 10px;
}

.inner:hover {
  color: white;
  background-color: darkgray;
}

.project-name {
  font-weight: bolder;
  font-size: 1.2em;
  padding: 0 0 5px 0
}

.project-package {
  font-size: 14px;
}

.project-path {
  font-size: 14px;
}
</style>
