<script lang="ts" setup>
const props = defineProps(['data'])
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
          <div class="float-right">{{ props.data.packages.toLocaleString('en-us', {}) }}</div>
        </li>
        <li class="list-group-item">
          <b>Source Files</b>
          <div class="float-right">{{ props.data.source_files.toLocaleString('en-us', {}) }}</div>
        </li>
        <li class="list-group-item">
          <b>Structs</b>
          <div class="float-right">{{ props.data.structs.toLocaleString('en-us', {}) }}</div>
        </li>
        <li class="list-group-item">
          <b>Interfaces</b>
          <div class="float-right">{{ props.data.interfaces.toLocaleString('en-us', {}) }}</div>
        </li>
        <li class="list-group-item">
          <b>Functions</b>
          <div class="float-right">{{ props.data.functions.toLocaleString('en-us', {}) }}</div>
        </li>
        <li class="list-group-item">
          <b>Methods</b>
          <div class="float-right">{{ props.data.methods.toLocaleString('en-us', {}) }}</div>
        </li>
        <li class="list-group-item">
          <b>Variables</b>
          <div class="float-right">{{ props.data.variables.toLocaleString('en-us', {}) }}</div>
        </li>
        <li class="list-group-item">
          <b>Constants</b>
          <div class="float-right">{{ props.data.constants.toLocaleString('en-us', {}) }}</div>
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
