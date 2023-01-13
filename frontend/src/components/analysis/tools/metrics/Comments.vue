<script lang="ts" setup>
const props = defineProps(['data'])
</script>

<template>
  <div class="card card-primary card-outline">
    <div class="card-header">
      <h3 class="card-title">Comments</h3>
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
          <b>Commented Lines</b>
          <div class="float-right">{{ props.data.total_lines.toLocaleString('en-us', {}) }}</div>
          <div class="ratio">{{ props.data.ratio.toLocaleString('en-us', {}) }}% of Total Lines of Code</div>
        </li>
        <li class="list-group-item">
          <b>Files with comments</b>
          <div class="float-right">{{ props.data.files_with_comments.toLocaleString('en-us', {}) }}</div>
          <div class="ratio">{{ props.data.files_with_comments_ratio.toLocaleString('en-us', {}) }}% of Total Files</div>
        </li>
      </ul>
    </div>
  </div>
</template>

<style scoped>
.list-group-item{
  padding: .25rem .25rem
}
.ratio {
  font-size: 12px;
  margin-left: 20px;
  margin-top: 10px;
}
</style>
