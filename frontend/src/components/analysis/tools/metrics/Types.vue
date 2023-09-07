<script lang="ts" setup>
import {reactive} from "vue";
import {storage} from "../../../../../wailsjs/go/models";
import HistoricalMetrics = storage.HistoricalMetrics;
import Metrics = storage.Metrics;
import HistoricChart from "./HistoryChart.vue";
import {COLORS, DataSetEntry, transparentize} from "./HistoryChartTypes";
import {resolveIcon, resolveStyle, resolveHint} from "./HistoryChartTypes";

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
      source_files: props.data.source_files,
      structs: props.data.structs,
      interfaces: props.data.interfaces,
      functions: props.data.functions,
      methods: props.data.methods,
      variables: props.data.variables,
      constants: props.data.constants,
      num_packages: props.data.packages
    })
  });
}

const resolveChartData = () => {
  const labels = ['']
  localData.historical.forEach(value => labels.push(value.date))
  labels.push('')

  const packagesData: (number | null)[] = [null]
  const sourceFilesData: (number | null)[] = [null]
  const structsData: (number | null)[] = [null]
  const interfacesData: (number | null)[] = [null]
  const functionsData: (number | null)[] = [null]
  const methodsData: (number | null)[] = [null]
  const variablesData: (number | null)[] = [null]
  const constantsData: (number | null)[] = [null]
  localData.historical.forEach(value => {
    packagesData.push(value.metrics.num_packages)
    sourceFilesData.push(value.metrics.source_files)
    structsData.push(value.metrics.structs)
    interfacesData.push(value.metrics.interfaces)
    functionsData.push(value.metrics.functions)
    methodsData.push(value.metrics.methods)
    variablesData.push(value.metrics.variables)
    constantsData.push(value.metrics.constants)
  })
  packagesData.push(null)
  sourceFilesData.push(null)
  structsData.push(null)
  interfacesData.push(null)
  functionsData.push(null)
  methodsData.push(null)
  variablesData.push(null)
  constantsData.push(null)

  return {
    labels,
    datasets: [
      {label: "Packages", borderColor: COLORS.red, backgroundColor: transparentize(COLORS.red), data: packagesData},
      {
        label: "Source Files",
        borderColor: COLORS.blue,
        backgroundColor: transparentize(COLORS.blue),
        data: sourceFilesData
      },
      {label: "Structs", borderColor: COLORS.green, backgroundColor: transparentize(COLORS.green), data: structsData},
      {
        label: "Interfaces",
        borderColor: COLORS.purple,
        backgroundColor: transparentize(COLORS.purple),
        data: interfacesData
      },
      {
        label: "Functions",
        borderColor: COLORS.orange,
        backgroundColor: transparentize(COLORS.orange),
        data: functionsData
      },
      {
        label: "Methods",
        borderColor: COLORS.yellow,
        backgroundColor: transparentize(COLORS.yellow),
        data: methodsData
      },
      {
        label: "Variables",
        borderColor: COLORS.waterblue,
        backgroundColor: transparentize(COLORS.waterblue),
        data: variablesData
      },
      {
        label: "Constants",
        borderColor: COLORS.pink,
        backgroundColor: transparentize(COLORS.pink),
        data: constantsData
      }
    ]
  }
}
</script>

<template>
  <div class="card card-primary card-outline">
    <div class="card-header">
      <h3 class="card-title">Types</h3>
      <div class="text-right">
        <div class="btn btn-sm btn-default action" title="View Metrics Chart"
             data-toggle="modal" data-target="#modal-types-metrics-chart" v-on:click="showChart">
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
          <b>Packages</b>
          <div class="float-right">
            {{ props.data.packages.toLocaleString('en-us', {}) }}
            <i v-if="props.previous != undefined"
              :class="resolveIcon(props.data.packages, props.previous.num_packages)"
              :style="resolveStyle(props.data.packages, props.previous.num_packages)"
              :title="resolveHint(props.data.packages, props.previous.num_packages)"
            ></i>
          </div>
        </li>
        <li class="list-group-item">
          <b>Source Files</b>
          <div class="float-right">
            {{ props.data.source_files.toLocaleString('en-us', {}) }}
            <i v-if="props.previous != undefined"
              :class="resolveIcon(props.data.source_files, props.previous.source_files)"
              :style="resolveStyle(props.data.source_files, props.previous.source_files)"
              :title="resolveHint(props.data.source_files, props.previous.source_files)"
            ></i>
          </div>
        </li>
        <li class="list-group-item">
          <b>Structs</b>
          <div class="float-right">
            {{ props.data.structs.toLocaleString('en-us', {}) }}
            <i v-if="props.previous != undefined"
              :class="resolveIcon(props.data.structs, props.previous.structs)"
              :style="resolveStyle(props.data.structs, props.previous.structs)"
              :title="resolveHint(props.data.structs, props.previous.structs)"
            ></i>
          </div>
        </li>
        <li class="list-group-item">
          <b>Interfaces</b>
          <div class="float-right">
            {{ props.data.interfaces.toLocaleString('en-us', {}) }}
            <i v-if="props.previous != undefined"
              :class="resolveIcon(props.data.interfaces, props.previous.interfaces)"
              :style="resolveStyle(props.data.interfaces, props.previous.interfaces)"
              :title="resolveHint(props.data.interfaces, props.previous.interfaces)"
            ></i>
          </div>
        </li>
        <li class="list-group-item">
          <b>Functions</b>
          <div class="float-right">
            {{ props.data.functions.toLocaleString('en-us', {}) }}
            <i v-if="props.previous != undefined"
              :class="resolveIcon(props.data.functions, props.previous.functions)"
              :style="resolveStyle(props.data.functions, props.previous.functions)"
              :title="resolveHint(props.data.functions, props.previous.functions)"
            ></i>
          </div>
        </li>
        <li class="list-group-item">
          <b>Methods</b>
          <div class="float-right">
            {{ props.data.methods.toLocaleString('en-us', {}) }}
            <i v-if="props.previous != undefined"
              :class="resolveIcon(props.data.methods, props.previous.methods)"
              :style="resolveStyle(props.data.methods, props.previous.methods)"
              :title="resolveHint(props.data.methods, props.previous.methods)"
            ></i>
          </div>
        </li>
        <li class="list-group-item">
          <b>Variables</b>
          <div class="float-right">
            {{ props.data.variables.toLocaleString('en-us', {}) }}
            <i v-if="props.previous != undefined"
              :class="resolveIcon(props.data.variables, props.previous.variables)"
              :style="resolveStyle(props.data.variables, props.previous.variables)"
              :title="resolveHint(props.data.variables, props.previous.variables)"
            ></i>
          </div>
        </li>
        <li class="list-group-item">
          <b>Constants</b>
          <div class="float-right">
            {{ props.data.constants.toLocaleString('en-us', {}) }}
            <i v-if="props.previous != undefined"
              :class="resolveIcon(props.data.constants, props.previous.constants)"
              :style="resolveStyle(props.data.constants, props.previous.constants)"
              :title="resolveHint(props.data.constants, props.previous.constants)"
            ></i>
          </div>
        </li>
      </ul>
    </div>
  </div>

  <div class="modal fade" id="modal-types-metrics-chart" style="display: none;" aria-hidden="true">
    <div class="modal-dialog modal-lg">
      <div class="modal-content">
        <div class="modal-header">
          <h4 class="modal-title">Type Metrics Chart</h4>
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
</style>
