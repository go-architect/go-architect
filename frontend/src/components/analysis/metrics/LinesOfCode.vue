<script lang="ts" setup>
import {reactive, watch} from "vue";
import {loc} from "../../../../wailsjs/go/models";

const props = defineProps(['data'])
const data = reactive({
  longestPackage: undefined as (loc.PackageLOC | undefined),
  longestFile: undefined as (loc.FileLOC | undefined),
})

watch(() => props.data, async (newVal: loc.ProjectLOC) => {
  data.longestPackage = newVal.packages.reduce(function(prev, current) {
    return (prev.loc > current.loc) ? prev : current
  })
  data.longestFile = newVal.files.reduce(function(prev, current) {
    return (prev.loc > current.loc) ? prev : current
  })
})
</script>

<template>
  <div class="card card-primary card-outline">
    <div class="card-header">
      <h3 class="card-title">Lines of Code</h3>
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
          <b>Total Lines of Code</b>
          <div class="float-right">{{ props.data.total.toLocaleString('en-us', {}) }}</div>
        </li>
        <li class="list-group-item">
          <b>Longest Package</b>
          <div class="float-right">{{ data.longestPackage?.loc }}</div>
          <div>
            <ul class="list-unstyled">
              <li class="details">
                <b>Package: </b>{{ data.longestPackage?.package }}
              </li>
            </ul>
          </div>
        </li>
        <li class="list-group-item">
          <b>Longest File</b>
          <div class="float-right">{{ data.longestFile?.loc }}</div>
          <div>
            <ul class="list-unstyled">
              <li class="details">
                <b>Package: </b>{{ data.longestFile?.package }}<br/>
                <b>File: </b>{{ data.longestFile?.file }}
              </li>
            </ul>
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
.details {
  font-size: 12px;
}
</style>
