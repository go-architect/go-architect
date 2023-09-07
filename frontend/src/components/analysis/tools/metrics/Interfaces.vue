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
      interface_avg_size: props.data.average_methods,
      interface_max_size: props.data.max_methods.length > 0 ? props.data.max_methods[0].methods : 0,
      interface_min_size: props.data.min_methods.length > 0 ? props.data.min_methods[0].methods : 0,
    })
  });
}

const resolveChartData = () => {
  const labels = ['']
  localData.historical.forEach(value => labels.push(value.date))
  labels.push('')

  const avgData: (number | null)[] = [null]
  const maxData: (number | null)[] = [null]
  const minData: (number | null)[] = [null]
  localData.historical.forEach(value => {
    avgData.push(value.metrics.interface_avg_size)
    maxData.push(value.metrics.interface_max_size)
    minData.push(value.metrics.interface_min_size)
  })
  avgData.push(null)
  maxData.push(null)
  minData.push(null)

  return {
    labels,
    datasets: [
      {
        label: "Avg Size",
        borderColor: COLORS.green,
        backgroundColor: transparentize(COLORS.green),
        data: avgData
      },
      {
        label: "Max Size",
        borderColor: COLORS.orange,
        backgroundColor: transparentize(COLORS.orange),
        data: maxData
      },
      {
        label: "Min Size",
        borderColor: COLORS.yellow,
        backgroundColor: transparentize(COLORS.yellow),
        data: minData
      }
    ]
  }
}
</script>

<template>
  <div class="card card-primary card-outline">
    <div class="card-header">
      <h3 class="card-title">Interfaces</h3>
      <div class="text-right">
        <div class="btn btn-sm btn-default action" title="View Metrics Chart"
             data-toggle="modal" data-target="#modal-interfaces-metrics-chart" v-on:click="showChart">
          <i class="fa-solid fa-chart-line"></i>
        </div>
      </div>
    </div>
    <div class="card-body">
      <div v-if="data == undefined" class="overlay-wrapper">
        <div class="overlay">
          <i class="fas fa-3x fa-sync-alt fa-spin"></i>
          <div class="text-bold pt-2">Loading...</div>
        </div>
        .......
      </div>
      <div v-else-if="data.min_methods?.length == undefined && data.max_methods?.length == undefined" class="overlay-wrapper">
        <b>The Go Project does not contain interfaces.</b>
      </div>
      <ul v-else class="list-group list-group-unbordered mb-3">
        <li class="list-group-item">
          <b>Average Methods</b>
          <div class="float-right">
            {{ data?.average_methods.toLocaleString('en-us', {}) }}
            <i v-if="props.previous != undefined"
               :class="resolveIcon(props.data.average_methods, props.previous.interface_avg_size)"
               :style="resolveStyle(props.data.average_methods, props.previous.interface_avg_size)"
               :title="resolveHint(props.data.average_methods, props.previous.interface_avg_size)"
            ></i>
          </div>
        </li>
        <li v-if="data!=null && data.min_methods?.length > 0" class="list-group-item">
          <b>Interfaces with Min number of methods</b>
          <div class="float-right">
            {{ data.min_methods[0].methods.toLocaleString('en-us', {}) }}
            <i v-if="props.previous != undefined"
               :class="resolveIcon(props.data.min_methods[0].methods, props.previous.interface_min_size)"
               :style="resolveStyle(props.data.min_methods[0].methods, props.previous.interface_min_size)"
               :title="resolveHint(props.data.min_methods[0].methods, props.previous.interface_min_size)"
            ></i>
          </div>
          <div>
            <ul class="list-unstyled">
              <li class="details" v-for="i in data.min_methods">
                <b>Package: </b>{{i.package }}
                <ul>
                  <li>{{ i.name }} in {{ i.file }}</li>
                </ul>
              </li>
            </ul>
          </div>
        </li>
        <li v-if="data.max_methods?.length > 0" class="list-group-item">
          <b>Interfaces with Max number of methods</b>
          <div class="float-right">
            {{ data.max_methods[0].methods.toLocaleString('en-us', {}) }}
            <i v-if="props.previous != undefined"
               :class="resolveIcon(props.data.max_methods[0].methods, props.previous.interface_max_size)"
               :style="resolveStyle(props.data.max_methods[0].methods, props.previous.interface_max_size)"
               :title="resolveHint(props.data.max_methods[0].methods, props.previous.interface_max_size)"
            ></i>
          </div>
          <div>
            <ul class="list-unstyled">
              <li class="details" v-for="i in data.max_methods">
                <b>Package: </b>{{i.package }}
                <ul>
                  <li>{{ i.name }} in {{ i.file }}</li>
                </ul>
              </li>
            </ul>
          </div>
        </li>
      </ul>
    </div>
  </div>

  <div class="modal fade" id="modal-interfaces-metrics-chart" style="display: none;" aria-hidden="true">
    <div class="modal-dialog modal-lg">
      <div class="modal-content">
        <div class="modal-header">
          <h4 class="modal-title">Interfaces Metrics Chart</h4>
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
.details {
  font-size: 12px;
}
</style>
