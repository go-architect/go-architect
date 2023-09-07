<script lang="ts" setup>

const props = defineProps(['data', 'previous'])

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
  console.log("current: " + current, "previous: "+previous, `-${previous-current} since last analysis`)
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
      <div class="text-right">
        <div class="btn btn-sm btn-default action" title="View Metrics Chart"
             data-toggle="modal" data-target="#modal-types-metrics-chart">
          <i class="fa-solid fa-chart-line"></i>
        </div>
      </div>
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
            <i v-if="props.previous != undefined"
              :class="resolveIcon(props.data.packages, props.previous.num_packages)"
              :style="resolveStyle(props.data.packages, props.previous.num_packages)"
              :title="resolveHint(props.data.packages, props.previous.num_packages)"
            ></i>
          </div>
        </li>
        <li class="list-group-item">
          <b>Source Files</b>
          <div class="float-right">
            {{ props.data.source_files.toLocaleString('en-us', {}) }}
            <i v-if="props.previous != undefined"
              :class="resolveIcon(props.data.source_files, props.previous.source_files)"
              :style="resolveStyle(props.data.source_files, props.previous.source_files)"
              :title="resolveHint(props.data.source_files, props.previous.source_files)"
            ></i>
          </div>
        </li>
        <li class="list-group-item">
          <b>Structs</b>
          <div class="float-right">
            {{ props.data.structs.toLocaleString('en-us', {}) }}
            <i v-if="props.previous != undefined"
              :class="resolveIcon(props.data.structs, props.previous.structs)"
              :style="resolveStyle(props.data.structs, props.previous.structs)"
              :title="resolveHint(props.data.structs, props.previous.structs)"
            ></i>
          </div>
        </li>
        <li class="list-group-item">
          <b>Interfaces</b>
          <div class="float-right">
            {{ props.data.interfaces.toLocaleString('en-us', {}) }}
            <i v-if="props.previous != undefined"
              :class="resolveIcon(props.data.interfaces, props.previous.interfaces)"
              :style="resolveStyle(props.data.interfaces, props.previous.interfaces)"
              :title="resolveHint(props.data.interfaces, props.previous.interfaces)"
            ></i>
          </div>
        </li>
        <li class="list-group-item">
          <b>Functions</b>
          <div class="float-right">
            {{ props.data.functions.toLocaleString('en-us', {}) }}
            <i v-if="props.previous != undefined"
              :class="resolveIcon(props.data.functions, props.previous.functions)"
              :style="resolveStyle(props.data.functions, props.previous.functions)"
              :title="resolveHint(props.data.functions, props.previous.functions)"
            ></i>
          </div>
        </li>
        <li class="list-group-item">
          <b>Methods</b>
          <div class="float-right">
            {{ props.data.methods.toLocaleString('en-us', {}) }}
            <i v-if="props.previous != undefined"
              :class="resolveIcon(props.data.methods, props.previous.methods)"
              :style="resolveStyle(props.data.methods, props.previous.methods)"
              :title="resolveHint(props.data.methods, props.previous.methods)"
            ></i>
          </div>
        </li>
        <li class="list-group-item">
          <b>Variables</b>
          <div class="float-right">
            {{ props.data.variables.toLocaleString('en-us', {}) }}
            <i v-if="props.previous != undefined"
              :class="resolveIcon(props.data.variables, props.previous.variables)"
              :style="resolveStyle(props.data.variables, props.previous.variables)"
              :title="resolveHint(props.data.variables, props.previous.variables)"
            ></i>
          </div>
        </li>
        <li class="list-group-item">
          <b>Constants</b>
          <div class="float-right">
            {{ props.data.constants.toLocaleString('en-us', {}) }}
            <i v-if="props.previous != undefined"
              :class="resolveIcon(props.data.constants, props.previous.constants)"
              :style="resolveStyle(props.data.constants, props.previous.constants)"
              :title="resolveHint(props.data.constants, props.previous.constants)"
            ></i>
          </div>
        </li>
      </ul>
    </div>
  </div>

  <div class="modal fade" id="modal-types-metrics-chart" style="display: none;" aria-hidden="true">
    <div class="modal-dialog modal-lg">
      <div class="modal-content">
        <div class="modal-header">
          <h4 class="modal-title">Type Metrics Chart</h4>
          <button type="button" class="close" data-dismiss="modal" aria-label="Close">
            <span aria-hidden="true">×</span>
          </button>
        </div>
        <div class="modal-body">
          Hola
        </div>
        <div class="modal-footer justify-content-between">
          <button type="button" class="btn btn-default" data-dismiss="modal">Cancel</button>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.list-group-item{
  padding: .25rem .25rem
}
</style>
