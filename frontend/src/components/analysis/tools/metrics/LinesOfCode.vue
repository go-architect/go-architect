<script lang="ts" setup>
import {reactive, watch} from "vue";
import {loc, storage} from "../../../../../wailsjs/go/models";
import {COLORS, DataSetEntry, transparentize} from "./HistoryChartTypes";
import {resolveIcon, resolveStyle, resolveHint} from "./HistoryChartTypes";
import HistoricChart from "./HistoryChart.vue";
import HistoricalMetrics = storage.HistoricalMetrics;
import Metrics = storage.Metrics;

const props = defineProps(['data', 'previous', 'historical'])
const data = reactive({
  longestPackage: undefined as (loc.PackageLOC | undefined),
  longestFile: undefined as (loc.FileLOC | undefined),
  historical: [] as HistoricalMetrics[],
  chartData: {
    labels: [] as string[],
    datasets: [] as DataSetEntry[]
  }
})

watch(() => props.data, async (newVal: loc.ProjectLOC) => {
  data.longestPackage = newVal.packages.reduce(function(prev, current) {
    return (prev.loc > current.loc) ? prev : current
  })
  data.longestFile = newVal.files.reduce(function(prev, current) {
    return (prev.loc > current.loc) ? prev : current
  })
})

const showChart = () => {
  if(props.historical !== undefined) {
    data.historical = [...props.historical]
    if(data.historical.length > 10){
      data.historical = data.historical.slice(props.historical.length - 10)
    }
  }
  data.historical.push(createTodayData())
  data.chartData = resolveChartData();
}

const createTodayData = () => {
  return HistoricalMetrics.createFrom({
    date: 'Today',
    commit: '',
    metrics: Metrics.createFrom({
      total_loc: props.data.total,
    })
  });
}

const resolveChartData = () => {
  const labels = ['']
  data.historical.forEach(value => labels.push(value.date))
  labels.push('')

  const locData: (number | null)[] = [null]
  data.historical.forEach(value => {
    locData.push(value.metrics.total_loc)
  })
  locData.push(null)

  return {
    labels,
    datasets: [
      {
        label: "Lines of Code",
        borderColor: COLORS.blue,
        backgroundColor: transparentize(COLORS.blue),
        data: locData
      }
    ]
  }
}
</script>

<template>
  <div class="card card-primary card-outline">
    <div class="card-header">
      <h3 class="card-title">Lines of Code</h3>
      <div class="text-right">
        <div class="btn btn-sm btn-default action" title="View Metrics Chart"
             data-toggle="modal" data-target="#modal-loc-metrics-chart" v-on:click="showChart">
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
          <b>Total Lines of Code</b>
          <div class="float-right">
            {{ props.data.total.toLocaleString('en-us', {}) }}
            <i v-if="props.previous != undefined"
               :class="resolveIcon(props.data.total, props.previous.total_loc)"
               :style="resolveStyle(props.data.total, props.previous.total_loc)"
               :title="resolveHint(props.data.total, props.previous.total_loc)"
            ></i>
          </div>
        </li>
        <li class="list-group-item">
          <b>Longest Package</b>
          <div class="float-right">{{ data.longestPackage?.loc.toLocaleString('en-us', {}) }}</div>
          <div>
            <ul class="list-unstyled">
              <li class="details">
                <b>Package: </b>{{ data.longestPackage?.package }}
              </li>
            </ul>
          </div>
        </li>
        <li class="list-group-item">
          <b>Longest File</b>
          <div class="float-right">{{ data.longestFile?.loc.toLocaleString('en-us', {}) }}</div>
          <div>
            <ul class="list-unstyled">
              <li class="details">
                <b>Package: </b>{{ data.longestFile?.package }}<br/>
                <b>File: </b>{{ data.longestFile?.file }}
              </li>
            </ul>
          </div>
        </li>
      </ul>
    </div>
  </div>

  <div class="modal fade" id="modal-loc-metrics-chart" style="display: none;" aria-hidden="true">
    <div class="modal-dialog modal-lg">
      <div class="modal-content">
        <div class="modal-header">
          <h4 class="modal-title">Lines of Code Metrics Chart</h4>
          <button type="button" class="close" data-dismiss="modal" aria-label="Close">
            <span aria-hidden="true">×</span>
          </button>
        </div>
        <div class="modal-body">
          <HistoricChart :data="data.chartData" />
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
.details {
  font-size: 12px;
}
</style>
