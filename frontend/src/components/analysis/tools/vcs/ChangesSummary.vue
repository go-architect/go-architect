<template>
  <div class="card card-primary card-outline">
    <div class="card-header">
      <h3 class="card-title">{{ props.data.title }}</h3>
    </div>
    <div class="card-body">
      <ul class="list-group list-group-unbordered mb-3">
        <table class="table table-bordered">
          <thead>
            <tr>
              <th style="width: 30px">#</th>
              <th>{{ props.data.kind }}</th>
              <th style="width: 90px">Status</th>
              <th style="width: 90px">Changes</th>
              <th style="width: 550px">Contributors</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(m,idx) in props.data.modifications">
              <td>{{idx+1}}</td>
              <td>{{ m.Source }}</td>
              <td class="text-center">
                <span :class="resolveStatus(m.AuthorsSummary)" :title="resolveStatusComments(m.AuthorsSummary)">
                  {{resolveStatusText(m.AuthorsSummary)}}
                </span>
              </td>
              <td class="text-center">{{ m.TotalModifications }}</td>
              <td>
                <div v-for="summary in sorted(m.AuthorsSummary)" :style="resolveStyle(summary, m.TotalModifications)"></div>
              </td>
            </tr>
          </tbody>
        </table>
      </ul>
    </div>
  </div>
</template>

<script lang="ts" setup>
const props = defineProps(['data','colors','contributors'])

function resolveStyle(summary: any, totalModifications: number) {
  const maxWidth = 500;
  const ratio = (100 * summary.Modifications / totalModifications)
  const width = Math.round(ratio * maxWidth / 100)
  return "float:left;height:20px;background-color:"+ props.colors.get(summary.Name) +"; width:"+width+"px"
}

function resolveStatus(authors: any) {
  if(authors.length === 1)
    return "badge bg-danger"

  if(props.contributors > 4 && authors.length <3)
    return "badge bg-warning"

  return "badge bg-success"
}

function resolveStatusText(authors: any) {
  if(authors.length === 1)
    return "Alert"

  if(props.contributors > 4 && authors.length <3)
    return "Warning"

  return "OK"
}

function resolveStatusComments(authors: any) {
  if(authors.length === 1)
    return "Just one contributor could pose a lack of maintainability risk"

  if(props.contributors > 4 && authors.length <3)
    return "Not enough contributors for this source item"

  return "Enough contributors for this source item"
}

function sorted(list: any[]){
  list.sort((a, b) => b.Modifications - a.Modifications)
  return list
}
</script>

<style scoped>
.table td {
  font-size: 10px;
  word-break:break-all;
  padding: 10px;
}

.badge {
  width: 70px;
  height: 15px;
  padding-top: 5px;
}
</style>
