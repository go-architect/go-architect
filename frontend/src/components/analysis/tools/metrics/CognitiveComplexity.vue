<script lang="ts" setup>
import { goToUrl } from "../../../../utils/goto";
import {reactive} from "vue";
import {resolveIcon, resolveStyle, resolveHint, COLORS, transparentize} from "./HistoryChartTypes";
import {DataSetEntry} from "./HistoryChartTypes";
import HistoricChart from "./HistoryChart.vue";
import {storage} from "../../../../../wailsjs/go/models";
import HistoricalMetrics = storage.HistoricalMetrics;
import Metrics = storage.Metrics;

const props = defineProps(['data', 'error', 'previous', 'historical'])

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
      cognitive_complexity_avg: props.data.AverageComplexity,
    })
  });
}

const resolveChartData = () => {
  const labels = ['']
  localData.historical.forEach(value => labels.push(value.date))
  labels.push('')

  const complexityData: (number | null)[] = [null]
  localData.historical.forEach(value => {
    complexityData.push(value.metrics.cognitive_complexity_avg)
  })
  complexityData.push(null)

  return {
    labels,
    datasets: [
      {
        label: "Cognitive Complexity",
        borderColor: COLORS.blue,
        backgroundColor: transparentize(COLORS.blue),
        data: complexityData
      }
    ]
  }
}
</script>

<template>
  <div class="card card-primary card-outline">
    <div class="card-header">
      <h3 class="card-title">Functions Cognitive Complexity</h3>
      <div class="text-right">
        <div class="btn btn-sm btn-default action" title="View Metrics Chart"
             data-toggle="modal" data-target="#modal-cognitive-complexity-metrics-chart" v-on:click="showChart">
          <i class="fa-solid fa-chart-line"></i>
        </div>
        <div class="btn btn-sm btn-default action" title="About Cognitive Complexity"
             v-on:click="goToUrl('https://github.com/uudashr/gocognit')">
          <i class="fa fa-question-circle" />
        </div>
      </div>
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
          <div class="float-right">
            {{ props.data.average_complexity }}
            <i v-if="props.previous != undefined"
               :class="resolveIcon(props.data.average_complexity, props.previous.cognitive_complexity_avg)"
               :style="resolveStyle(props.data.average_complexity, props.previous.cognitive_complexity_avg)"
               :title="resolveHint(props.data.average_complexity, props.previous.cognitive_complexity_avg)"
            ></i>
          </div>
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


  <div class="modal fade" id="modal-cognitive-complexity-metrics-chart" style="display: none;" aria-hidden="true">
    <div class="modal-dialog modal-lg">
      <div class="modal-content">
        <div class="modal-header">
          <h4 class="modal-title">Cognitive Complexity Chart</h4>
          <button type="button" class="close" data-dismiss="modal" aria-label="Close">
            <span aria-hidden="true">×</span>
          </button>
        </div>
        <div class="modal-body">
          <HistoricChart :data="localData.chartData" />
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
.function-details {
  font-size: 12px;
}
</style>
