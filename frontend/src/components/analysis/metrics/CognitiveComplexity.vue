<script lang="ts" setup>
const props = defineProps(['data', 'error'])
</script>

<template>
  <div class="card card-primary card-outline">
    <div class="card-header">
      <h3 class="card-title">Functions Cognitive Complexity</h3>
    </div>
    <div class="card-body">
      <div v-if="props.data == undefined && props.error == undefined" class="overlay-wrapper">
        <div class="overlay">
          <i class="fas fa-3x fa-sync-alt fa-spin"></i>
          <div class="text-bold pt-2">Loading...</div>
        </div>
        .......
      </div>
      <ul v-else-if="props.error != null" class="list-group list-group-unbordered mb-3">
        <li class="list-group-item">
          <div><b>Can't execute tool</b></div>
          <div>{{ props.error }}</div>
        </li>
      </ul>
      <ul v-else class="list-group list-group-unbordered mb-3">
        <li class="list-group-item">
          <b>Average Complexity</b>
          <div class="float-right">{{ props.data.AverageComplexity }}</div>
        </li>
        <li class="list-group-item">
          <b>Top 5 Complex Functions</b>
          <div>
            <ul v-for="f in props.data.ComplexityDetails" class="list-unstyled">
              <li>
                <b>Package: </b>{{ f.Package }}
                <ul>
                  <li class="function-details">({{ f.Complexity }}) {{ f.Function }} in <i>{{ f.File }}</i></li>
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
.function-details {
  font-size: 12px;
}
</style>
