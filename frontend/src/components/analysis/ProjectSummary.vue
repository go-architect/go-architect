<script lang="ts" setup>
import {onMounted, reactive, watch} from "vue";
import {GetRepositoryInfo} from "../../../wailsjs/go/main/Api";

const emits = defineEmits(["close"])
const props = defineProps(['project'])
const data = reactive({
  repository: undefined
})

function closeProject() {
  console.log("Summary - Close")
  emits("close")
}

onMounted(() => {
  GetRepositoryInfo(props.project.path).then((result: any) => {
    data.repository = result
  })
})
</script>

<template>
  <div class="card">
    <div class="card-header">
      <h3 class="card-title">Project Summary</h3>
      <div class="text-right">
        <div class="btn btn-sm btn-primary" title="Close Project" v-on:click="closeProject">
          <i class="fa fa-close" />
        </div>
      </div>

    </div>
    <div class="card-body">
      <div class="row">
        <div class="col-12 col-md-4 col-lg-4">
          <div class="text-muted">
            <p class="text-sm">
              <b class="d-block">Project Name</b>
              {{ props.project.id }}
            </p>
          </div>
        </div>
        <div class="col-12 col-md-4 col-lg-4">
          <div class="text-muted">
            <p class="text-sm">
              <b class="d-block">Main Package</b>
              {{ props.project.name }}
            </p>
          </div>
        </div>
        <div class="col-12 col-md-4 col-lg-4">
          <div class="text-muted">
            <p class="text-sm">
              <b class="d-block">Project Path</b>
              {{ props.project.path }}
            </p>
          </div>
        </div>
        <div class="col-12 col-md-4 col-lg-4">
          <div class="text-muted">
            <p class="text-sm">
              <b class="d-block">Repository URL</b>
              {{ (data.repository==undefined) ? "-":data.repository["url"] }}
            </p>
          </div>
        </div>
        <div class="col-12 col-md-4 col-lg-4">
          <div class="text-muted">
            <p class="text-sm">
              <b class="d-block">Current Branch</b>
              {{ (data.repository==undefined) ? "-":data.repository["branch"] }}
            </p>
          </div>
        </div>
        <div class="col-12 col-md-4 col-lg-4">
          <div class="text-muted">
            <p class="text-sm">
              <b class="d-block">Current Revision</b>
              {{ (data.repository==undefined) ? "-":data.repository["revision"] }}
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
</style>
