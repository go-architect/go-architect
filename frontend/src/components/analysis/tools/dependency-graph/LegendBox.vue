<script lang="ts">
import {defineComponent} from "vue";

export default defineComponent({
  name: "Legend"
})
</script>

<template>
  <div class="card card-primary card-outline">
    <div class="card-header">
      <h3 class="card-title">Legend</h3>
    </div>
    <div class="card-body">
      <ul class="list-group list-group-unbordered mb-3">
        <li class="list-group-item">
          <b>Internal Packages</b>
          <div class="float-right">
            <span class="dot" style="background-color: #80d6ff"></span>
          </div>
        </li>
        <li class="list-group-item">
          <b>External Packages</b>
          <div class="float-right">
            <span class="dot" style="background-color: #f06868"></span>
          </div>
        </li>
        <li class="list-group-item">
          <b>StandardLib Packages</b>
          <div class="float-right">
            <span class="dot" style="background-color: #edf798"></span>
          </div>
        </li>
        <li class="list-group-item">
          <b>Organization Packages</b>
          <div class="float-right">
            <span class="dot" style="background-color: #fab57a"></span>
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
.dot {
  height: 25px;
  width: 25px;
  border-radius: 50%;
  display: inline-block;
}
</style>
