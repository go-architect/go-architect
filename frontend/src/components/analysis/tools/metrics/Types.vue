<script lang="ts" setup>
const props = defineProps(['data'])

const previousData = {
  packages: 10,
  source_files: 12,
  structs: 5,
  interfaces: 0,
  functions: 7,
  methods: 3,
  variables: 7,
  constants: 38
}
const resolveIcon = (current: number, previous:number): string => {
  if (current == previous) {
    return "fa-solid fa-circle"
  }
  if (current > previous) {
    return "fa-solid fa-circle-up"
  }
  return "fa-solid fa-circle-down"
}
const resolveStyle = (current: number, previous:number): string => {
  if (current == previous) {
    return "color: #f1d104;"
  }
  if (current > previous) {
    return "color: #06a741;"
  }
  return "color: #df0707;"
}
const resolveHint = (current: number, previous:number): string => {
  if (current == previous) {
    return "No changes since last analysis"
  }
  if (current > previous) {
    return `+${current-previous} since last analysis`
  }
  return `-${previous-current} since last analysis`
}
</script>

<template>
  <div class="card card-primary card-outline">
    <div class="card-header">
      <h3 class="card-title">Types</h3>
    </div>
    <div class="card-body">
      <div v-if="props.data == undefined" class="overlay-wrapper">
        <div class="overlay">
          <i class="fas fa-3x fa-sync-alt fa-spin"></i>
          <div class="text-bold pt-2">Loading...</div>
        </div>
        .......
      </div>
      <ul v-else class="list-group list-group-unbordered mb-3">
        <li class="list-group-item">
          <b>Packages</b>
          <div class="float-right">
            {{ props.data.packages.toLocaleString('en-us', {}) }}
            <i v-if="previousData != undefined"
              :class="resolveIcon(props.data.packages, previousData.packages)"
              :style="resolveStyle(props.data.packages, previousData.packages)"
              :title="resolveHint(props.data.packages, previousData.packages)"
            ></i>
          </div>
        </li>
        <li class="list-group-item">
          <b>Source Files</b>
          <div class="float-right">
            {{ props.data.source_files.toLocaleString('en-us', {}) }}
            <i v-if="previousData != undefined"
              :class="resolveIcon(props.data.source_files, previousData.source_files)"
              :style="resolveStyle(props.data.source_files, previousData.source_files)"
              :title="resolveHint(props.data.source_files, previousData.source_files)"
            ></i>
          </div>
        </li>
        <li class="list-group-item">
          <b>Structs</b>
          <div class="float-right">
            {{ props.data.structs.toLocaleString('en-us', {}) }}
            <i v-if="previousData != undefined"
              :class="resolveIcon(props.data.structs, previousData.structs)"
              :style="resolveStyle(props.data.structs, previousData.structs)"
              :title="resolveHint(props.data.structs, previousData.structs)"
            ></i>
          </div>
        </li>
        <li class="list-group-item">
          <b>Interfaces</b>
          <div class="float-right">
            {{ props.data.interfaces.toLocaleString('en-us', {}) }}
            <i v-if="previousData != undefined"
              :class="resolveIcon(props.data.interfaces, previousData.interfaces)"
              :style="resolveStyle(props.data.interfaces, previousData.interfaces)"
              :title="resolveHint(props.data.interfaces, previousData.interfaces)"
            ></i>
          </div>
        </li>
        <li class="list-group-item">
          <b>Functions</b>
          <div class="float-right">
            {{ props.data.functions.toLocaleString('en-us', {}) }}
            <i v-if="previousData != undefined"
              :class="resolveIcon(props.data.functions, previousData.functions)"
              :style="resolveStyle(props.data.functions, previousData.functions)"
              :title="resolveHint(props.data.functions, previousData.functions)"
            ></i>
          </div>
        </li>
        <li class="list-group-item">
          <b>Methods</b>
          <div class="float-right">
            {{ props.data.methods.toLocaleString('en-us', {}) }}
            <i v-if="previousData != undefined"
              :class="resolveIcon(props.data.methods, previousData.methods)"
              :style="resolveStyle(props.data.methods, previousData.methods)"
              :title="resolveHint(props.data.methods, previousData.methods)"
            ></i>
          </div>
        </li>
        <li class="list-group-item">
          <b>Variables</b>
          <div class="float-right">
            {{ props.data.variables.toLocaleString('en-us', {}) }}
            <i v-if="previousData != undefined"
              :class="resolveIcon(props.data.variables, previousData.variables)"
              :style="resolveStyle(props.data.variables, previousData.variables)"
              :title="resolveHint(props.data.variables, previousData.variables)"
            ></i>
          </div>
        </li>
        <li class="list-group-item">
          <b>Constants</b>
          <div class="float-right">
            {{ props.data.constants.toLocaleString('en-us', {}) }}
            <i v-if="previousData != undefined"
              :class="resolveIcon(props.data.constants, previousData.constants)"
              :style="resolveStyle(props.data.constants, previousData.constants)"
              :title="resolveHint(props.data.constants, previousData.constants)"
            ></i>
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>

<style scoped>
.list-group-item{
  padding: .25rem .25rem
}
</style>
