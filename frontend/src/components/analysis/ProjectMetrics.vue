<script lang="ts" setup>
import {
  GetMetricsLOC,
  GetMetricsInterfaces,
  GetMetricsTypes,
  GetMetricsComments,
  GetMetricsComplexity
} from "../../../wailsjs/go/main/Api";
import Comments from "./metrics/Comments.vue"
import Complexity from "./metrics/Complexity.vue"
import Interfaces from "./metrics/Interfaces.vue"
import LinesOfCode from "./metrics/LinesOfCode.vue"
import Types from "./metrics/Types.vue"
import {comments, interfaces, loc, project, types} from "../../../wailsjs/go/models";
import {onMounted, reactive} from "vue";

const props = defineProps(['project'])
const data = reactive({
  loc: undefined as loc.ProjectLOC | undefined,
  comments: undefined as comments.CommentsMetrics | undefined,
  types: undefined as types.ProjectTypes | undefined,
  interfaces: undefined as interfaces.InterfaceMetrics | undefined,
  cyclo: undefined,
})

onMounted(async () => {
  if(props.project != null){
    [data.loc, data.comments, data.types, data.interfaces, data.cyclo] = await Promise.all([
      GetMetricsLOC(props.project),
      GetMetricsComments(props.project),
      GetMetricsTypes(props.project),
      GetMetricsInterfaces(props.project),
      GetMetricsComplexity(props.project)
    ])
    console.log(data.cyclo)
  }
})
</script>

<template>
  <div class="card">
    <div class="card-header">
      <h3 class="card-title">Project Metrics</h3>
    </div>
    <div class="card-body">
      <div class="row">
        <div class="col-md-12 col-lg-4">
          <lines-of-code :data=data.loc />
          <comments :data=data.comments />
        </div>
        <div class="col-md-12 col-lg-4">
          <types :data=data.types />
          <interfaces :data="data.interfaces" />
        </div>
        <div class="col-md-12 col-lg-4">
          <complexity :data=data.cyclo />
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
</style>
