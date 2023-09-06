<script lang="ts" setup>
import {
  GetMetricsLOC,
  GetMetricsInterfaces,
  GetMetricsTypes,
  GetMetricsComments,
  GetMetricsComplexity,
  GetMetricsCognitiveComplexity
} from "../../../../../wailsjs/go/main/Api";
import Comments from "./Comments.vue"
import Complexity from "./Complexity.vue"
import CognitiveComplexity from "./CognitiveComplexity.vue"
import Interfaces from "./Interfaces.vue"
import LinesOfCode from "./LinesOfCode.vue"
import Types from "./Types.vue"
import {comments, interfaces, loc, project, types} from "../../../../../wailsjs/go/models";
import {onMounted, reactive} from "vue";
import {getSelectedProject} from "../../../../utils/storage";

const data = reactive({
  loc: undefined as loc.ProjectLOC | undefined,
  comments: undefined as comments.CommentsMetrics | undefined,
  types: undefined as types.ProjectTypes | undefined,
  interfaces: undefined as interfaces.InterfaceMetrics | undefined,
  cyclomaticComplexity: undefined,
  cognitiveComplexity: undefined,
  cyclomaticComplexityError: undefined,
  cognitiveComplexityError: undefined,
})

onMounted(async () => {
  const project = await getSelectedProject()
  if(project != null){
    [data.loc, data.comments, data.types, data.interfaces] = await Promise.all([
      GetMetricsLOC(project),
      GetMetricsComments(project),
      GetMetricsTypes(project),
      GetMetricsInterfaces(project),
    ])
    try {
      data.cognitiveComplexity = await GetMetricsCognitiveComplexity(project)
    } catch(error: any) {
      data.cognitiveComplexityError = error
    }
    try {
      data.cyclomaticComplexity = await GetMetricsComplexity(project)
    } catch(error: any) {
      data.cyclomaticComplexity = error
    }
  }
})
</script>

<template>
  <div class="card">
    <div class="card-header">
      <h3 class="card-title">Project Metrics</h3>
      <div class="card-tools">
        <router-link to="/analysis" class="btn btn-sm btn-primary" title="Back to Project Analysis">
          <i class="fa fa-circle-left" /> Back to Project Analysis
        </router-link>
      </div>
    </div>
    <div class="card-body">
      <div class="row">
        <div class="col-md-12 col-lg-4">
          <lines-of-code :data=data.loc />
          <interfaces :data="data.interfaces" />
        </div>
        <div class="col-md-12 col-lg-3">
          <types :data=data.types />
          <comments :data=data.comments />
        </div>
        <div class="col-md-12 col-lg-5">
          <complexity :data=data.cyclomaticComplexity :error=data.cyclomaticComplexityError />
          <cognitive-complexity :data=data.cognitiveComplexity :error=data.cognitiveComplexityError />
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
</style>
