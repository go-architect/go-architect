<template>
  <div class="card card-primary card-outline">
    <div class="card-header">
      <h3 class="card-title">Contributors</h3>
    </div>
    <div class="card-body">
      <ul class="list-group list-group-unbordered mb-3">
        <li v-for="author in sortedContributors()" class="list-group-item contributor">
          <i class="fa-solid fa-circle author-color" :style="resolveColor(author.Name)"></i>
          <b>{{ author.Name }}</b>
          <div class="float-right">{{ (100 * author.ModifiedLines / props.data.TotalModifications).toLocaleString('en-us', {}) }}%</div>
          <div class="ratio">{{ author.Commits }} Commits, {{ author.ModifiedLines }} Modified lines</div>
        </li>
      </ul>
    </div>
  </div>
</template>

<script lang="ts" setup>
const props = defineProps(['data','colors'])

function sortedContributors() {
  props.data.AuthorDetails.sort((a: any, b: any) => b.ModifiedLines - a.ModifiedLines)
  return props.data.AuthorDetails
}

function resolveColor(author: string) {
  return "color: "+ props.colors.get(author) +";"
}
</script>

<style scoped>
.contributor {
  font-size: 14px;
}
.author-color {
  margin-right: 10px;
}
.ratio {
  font-size: 12px;
}
</style>
