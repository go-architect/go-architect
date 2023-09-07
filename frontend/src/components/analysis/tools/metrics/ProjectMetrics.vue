<script lang="ts" setup>
import {onMounted, reactive} from "vue";
import moment from "moment";
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
import {comments, gocyclo, interfaces, loc, storage, types} from "../../../../../wailsjs/go/models";
import {getSelectedProject, saveSelectedProject} from "../../../../utils/storage";
import HistoricalMetrics = storage.HistoricalMetrics;
import Metrics = storage.Metrics;

const data = reactive({
  loc: undefined as loc.ProjectLOC | undefined,
  comments: undefined as comments.CommentsMetrics | undefined,
  types: undefined as types.ProjectTypes | undefined,
  interfaces: undefined as interfaces.InterfaceMetrics | undefined,
  cyclomaticComplexity: undefined as gocyclo.GoCycloOutput | undefined,
  cognitiveComplexity: undefined as gocyclo.GoCycloOutput | undefined,
  cyclomaticComplexityError: undefined,
  cognitiveComplexityError: undefined,
  historicalMetrics: [] as HistoricalMetrics[],
  previousMetricsElement: undefined as HistoricalMetrics | undefined,
  previousMetrics: undefined as Metrics | undefined,
  todayFormatted: "",
  lastSavedOn: ""
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
    if(project.historical_metrics != null && project.historical_metrics.length>0) {
      data.historicalMetrics = [...project.historical_metrics.sort((a, b) => a.date.localeCompare(b.date))]
      data.previousMetricsElement = project.historical_metrics.sort((a, b) => b.date.localeCompare(a.date)).at(0)
//      console.log("LastData ", data.previousMetricsElement)
      data.previousMetrics = data.previousMetricsElement?.metrics
//      console.log("LastMetrics ", data.previousMetrics)
//      console.log("historicalMetrics ", data.historicalMetrics)
    }
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

    const today = new Date();
    let mm: (number|string) = today.getMonth() + 1; // Months start at 0!
    let dd: (number|string)  = today.getDate();
    if (mm < 10) mm = '0' + mm;
    if (dd < 10) dd = '0' + dd;
    data.todayFormatted = today.getFullYear()+"-"+mm+"-"+dd;
    if (data.previousMetricsElement?.date != data.todayFormatted) {
      const dt = moment(data.previousMetricsElement?.date)
      data.lastSavedOn = "on " + dt.format("LL")
    } else {
      const dt = moment(data.previousMetricsElement?.date)
      data.lastSavedOn = "Today ("+dt.format("LL")+")"
    }
  }
})

const saveMetrics = async () => {
  const project = await getSelectedProject()
  data.historicalMetrics = data.historicalMetrics.filter(hv => hv.date !== data.todayFormatted)
  data.historicalMetrics.push(HistoricalMetrics.createFrom({
    date: data.todayFormatted,
    metrics: Metrics.createFrom({
      num_packages: data.types?.packages || 0,
      source_files: data.types?.source_files || 0,
      structs: data.types?.structs || 0,
      interfaces: data.types?.interfaces || 0,
      functions: data.types?.functions || 0,
      methods: data.types?.methods || 0,
      variables: data.types?.variables || 0,
      constants: data.types?.constants || 0,
      total_loc: data.loc?.total || 0,
      interface_avg_size: data.interfaces?.average_methods || 0,
      interface_max_size: data.interfaces?.max_methods[0]?.methods || 0,
      interface_min_size: data.interfaces?.min_methods[0]?.methods || 0,
      comments_total_lines: data.comments?.total_lines || 0,
      comments_lines_ratio: data.comments?.ratio || 0,
      files_with_comments: data.comments?.files_with_comments || 0,
      files_with_comments_ratio: data.comments?.files_with_comments_ratio || 0,
      cyclomatic_complexity_avg: data.cyclomaticComplexity?.average_complexity || 0,
      cognitive_complexity_avg: data.cognitiveComplexity?.average_complexity || 0
    })
  }))
  project!.historical_metrics = data.historicalMetrics
  await saveSelectedProject(project!)
  const dt = moment()
  data.lastSavedOn = "Today ("+dt.format("LL")+")"
}
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
      <div class="row" style="margin-bottom: 10px;">
        <div class="col-md-12">
          <div class="save-div">
            <div style="display: inline; margin-right: 5px;">The metrics were last saved {{ data.lastSavedOn }}.</div>
            <div class="btn btn-sm btn-default action" title="Save Metrics" v-on:click="saveMetrics">
              <i class="fa-regular fa-floppy-disk"></i>
            </div>
          </div>
        </div>
      </div>
      <div class="row">
        <div class="col-md-12 col-lg-4">
          <lines-of-code :data=data.loc :previous=data.previousMetrics :historical=data.historicalMetrics />
          <interfaces :data="data.interfaces" :previous=data.previousMetrics :historical=data.historicalMetrics />
        </div>
        <div class="col-md-12 col-lg-3">
          <types :data=data.types :previous=data.previousMetrics :historical=data.historicalMetrics />
          <comments :data=data.comments :previous=data.previousMetrics :historical=data.historicalMetrics />
        </div>
        <div class="col-md-12 col-lg-5">
          <complexity
              :data=data.cyclomaticComplexity
              :error=data.cyclomaticComplexityError
              :previous=data.previousMetrics
              :historical=data.historicalMetrics />
          <cognitive-complexity
              :data=data.cognitiveComplexity
              :error=data.cognitiveComplexityError
              :previous=data.previousMetrics
              :historical=data.historicalMetrics />
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.save-div {
  float: right;
  border: 2px dotted #007bff;
  border-radius: 5px;
  padding: 10px;
}
</style>
