<script lang="ts" setup>
const props = defineProps(['data'])
</script>

<template>
  <div class="card card-primary card-outline">
    <div class="card-header">
      <h3 class="card-title">Interfaces</h3>
    </div>
    <div class="card-body">
      <div v-if="data == undefined" class="overlay-wrapper">
        <div class="overlay">
          <i class="fas fa-3x fa-sync-alt fa-spin"></i>
          <div class="text-bold pt-2">Loading...</div>
        </div>
        .......
      </div>
      <ul v-else class="list-group list-group-unbordered mb-3">
        <li class="list-group-item">
          <b>Average Methods</b>
          <div class="float-right">{{ data?.average_methods.toLocaleString('en-us', {}) }}</div>
        </li>
        <li v-if="(data.min_methods.length + data.max_methods.length) == 0" class="list-group-item">
          <b>The Go Project does not contain interfaces</b>
        </li>
        <li v-if="data.min_methods.length > 0" class="list-group-item">
          <b>Interfaces with Min number of methods</b>
          <div class="float-right">{{ data.min_methods[0].methods.toLocaleString('en-us', {}) }}</div>
          <div>
            <ul class="list-unstyled">
              <li class="details" v-for="i in data.min_methods">
                <b>Package: </b>{{i.package }}
                <ul>
                  <li>{{ i.name }} in {{ i.file }}</li>
                </ul>
              </li>
            </ul>
          </div>
        </li>
        <li v-if="data.max_methods.length > 0" class="list-group-item">
          <b>Interfaces with Max number of methods</b>
          <div class="float-right">{{ data.max_methods[0].methods.toLocaleString('en-us', {}) }}</div>
          <div>
            <ul class="list-unstyled">
              <li class="details" v-for="i in data.max_methods">
                <b>Package: </b>{{i.package }}
                <ul>
                  <li>{{ i.name }} in {{ i.file }}</li>
                </ul>
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
