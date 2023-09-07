<script lang="ts" setup>
import {reactive} from "vue";
import {storage} from "../../../../../wailsjs/go/models";
import {COLORS, DataSetEntry, transparentize} from "./HistoryChartTypes";
import {resolveIcon, resolveStyle, resolveHint} from "./HistoryChartTypes";
import HistoricChart from "./HistoryChart.vue";
import HistoricalMetrics = storage.HistoricalMetrics;
import Metrics = storage.Metrics;

const props = defineProps(['data', 'previous', 'historical'])


const localData = reactive({
  historical: [] as HistoricalMetrics[],
  chartData: {
    labels: [] as string[],
    datasets: [] as DataSetEntry[]
  }
})

const showChart = () => {
  if(props.historical !== undefined) {
    localData.historical = [...props.historical]
    if(localData.historical.length > 10){
      localData.historical = localData.historical.slice(props.historical.length - 10)
    }
  }
  localData.historical.push(createTodayData())
  localData.chartData = resolveChartData();
}

const createTodayData = () => {
  return HistoricalMetrics.createFrom({
    date: 'Today',
    commit: '',
    metrics: Metrics.createFrom({
      comments_total_lines: props.data.total_lines,
      comments_lines_ratio: props.data.ratio,
      files_with_comments: props.data.files_with_comments,
      files_with_comments_ratio: props.data.files_with_comments_ratio
    })
  });
}

const resolveChartData = () => {
  const labels = ['']
  localData.historical.forEach(value => labels.push(value.date))
  labels.push('')

  const commentsData: (number | null)[] = [null]
  const commentsRatioData: (number | null)[] = [null]
  const filesData: (number | null)[] = [null]
  const filesRatioData: (number | null)[] = [null]
  localData.historical.forEach(value => {
    commentsData.push(value.metrics.comments_total_lines)
    commentsRatioData.push(value.metrics.comments_lines_ratio)
    filesData.push(value.metrics.files_with_comments)
    filesRatioData.push(value.metrics.files_with_comments_ratio)
  })
  commentsData.push(null)
  commentsRatioData.push(null)
  filesData.push(null)
  filesRatioData.push(null)

  return {
    labels,
    datasets: [
      {
        label: "Comments Lines",
        borderColor: COLORS.green,
        backgroundColor: transparentize(COLORS.green),
        data: commentsData,
        yAxisID: 'y'
      },
      {
        label: "Comments Lines Ratio",
        borderColor: COLORS.orange,
        backgroundColor: transparentize(COLORS.orange),
        data: commentsRatioData,
        yAxisID: 'y1'
      },
      {
        label: "File With Comments",
        borderColor: COLORS.yellow,
        backgroundColor: transparentize(COLORS.yellow),
        data: filesData,
        yAxisID: 'y'
      },
      {
        label: "File With Comments Ratio",
        borderColor: COLORS.purple,
        backgroundColor: transparentize(COLORS.purple),
        data: filesRatioData,
        yAxisID: 'y1'
      }
    ]
  }
}
</script>

<template>
  <div class="card card-primary card-outline">
    <div class="card-header">
      <h3 class="card-title">Comments</h3>
      <div class="text-right">
        <div class="btn btn-sm btn-default action" title="View Metrics Chart"
             data-toggle="modal" data-target="#modal-comments-metrics-chart" v-on:click="showChart">
          <i class="fa-solid fa-chart-line"></i>
        </div>
      </div>
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
          <div class="float-right">
            {{ props.data.total_lines.toLocaleString('en-us', {}) }}
            <i v-if="props.previous != undefined"
               :class="resolveIcon(props.data.total_lines, props.previous.comments_total_lines)"
               :style="resolveStyle(props.data.total_lines, props.previous.comments_total_lines)"
               :title="resolveHint(props.data.total_lines, props.previous.comments_total_lines)"
            ></i>
          </div>
          <div class="ratio">{{ props.data.ratio.toLocaleString('en-us', {}) }}% of Total Lines of Code</div>
        </li>
        <li class="list-group-item">
          <b>Files with comments</b>
          <div class="float-right">
            {{ props.data.files_with_comments.toLocaleString('en-us', {}) }}
            <i v-if="props.previous != undefined"
               :class="resolveIcon(props.data.files_with_comments, props.previous.files_with_comments)"
               :style="resolveStyle(props.data.files_with_comments, props.previous.files_with_comments)"
               :title="resolveHint(props.data.files_with_comments, props.previous.files_with_comments)"
            ></i>
          </div>
          <div class="ratio">{{ props.data.files_with_comments_ratio.toLocaleString('en-us', {}) }}% of Total Files</div>
        </li>
      </ul>
    </div>
  </div>

  <div class="modal fade" id="modal-comments-metrics-chart" style="display: none;" aria-hidden="true">
    <div class="modal-dialog modal-lg">
      <div class="modal-content">
        <div class="modal-header">
          <h4 class="modal-title">Comments Metrics Chart</h4>
          <button type="button" class="close" data-dismiss="modal" aria-label="Close">
            <span aria-hidden="true">×</span>
          </button>
        </div>
        <div class="modal-body">
          <HistoricChart :data="localData.chartData" :show-ratio-scale="true" />
        </div>
        <div class="modal-footer justify-content-between">
          <button type="button" class="btn btn-default" data-dismiss="modal">Cancel</button>
        </div>
      </div>
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
