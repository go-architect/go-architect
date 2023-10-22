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
      <div class="small-box-footer" data-toggle="modal" data-target="#modal-delete">Delete Project <i class="fas fa-trash"></i></div>
    </div>
  </div>

  <div class="modal fade" id="modal-delete" style="display: none;" aria-hidden="true">
    <div class="modal-dialog">
      <div class="modal-content bg-danger">
        <div class="modal-header">
          <h4 class="modal-title">Remove Project Confirmation</h4>
          <button type="button" class="close" data-dismiss="modal" aria-label="Close">
            <span aria-hidden="true">×</span>
          </button>
        </div>
        <div class="modal-body">
          <p>Please confirm if you really want to remove this project from Go-Architect. This action cannot be reversed.</p>
        </div>
        <div class="modal-footer justify-content-between">
          <button type="button" class="btn btn-outline-light" data-dismiss="modal">Cancel</button>
          <button type="button" class="btn btn-outline-light" data-dismiss="modal" v-on:click="deleteProject(this, project.id)">Confirm</button>
        </div>
      </div>

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
