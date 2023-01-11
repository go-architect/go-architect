<template>
  <div class="card card-primary card-outline">
    <div class="card-header">
      <h3 class="card-title">Coupling Analysis Result</h3>
    </div>
    <div class="card-body">
      <ul class="list-group list-group-unbordered mb-3">
        <li class="list-group-item">
          <b>Package</b>
          <div class="float-right">{{ result?.dependency }}</div>
        </li>
        <li class="list-group-item">
          <b>Overall Coupling Level</b>
          <div class="float-right">{{ result?.coupling_level.toLocaleString('en-us', {}) }}</div>
        </li>
        <li class="list-group-item">
          <div class="inner-details">
            <ul class="list-group list-group-unbordered mb-3">
              <li v-for="pkgCoupling in result?.details" class="list-group-item">
                <b>{{ pkgCoupling.package }}</b>
                <div class="float-right">Coupling Level: {{ pkgCoupling.coupling_level.toLocaleString('en-us', {}) }}</div>
                <div class="col-md-8">
                <div class="inner-details">
                  <ul class="list-group list-group-unbordered mb-3">
                    <li v-for="fileCoupling in pkgCoupling.details" class="list-group-item">
                        <b>{{ fileCoupling.file }}</b>
                        <div class="float-right">Coupling Level: {{ fileCoupling.coupling_level.toLocaleString('en-us', {}) }}</div>
                        <div class="inner-details col-md-8">
                          <table class="coupling table table-bordered">
                            <thead>
                            <tr>
                              <th style="width:80px;">Line</th>
                              <th>Details</th>
                            </tr>
                            </thead>
                            <tbody>
                            <tr v-for="d in fileCoupling.details">
                              <td style="text-align: center;">{{ d.line}}</td>
                              <td>{{  d.details }}</td>
                            </tr>
                            </tbody>
                          </table>
                        </div>
                    </li>
                  </ul>
                </div>
                </div>
              </li>
            </ul>
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>

<script lang="ts">
import {PropType} from "vue";
import {coupling} from "../../../../../wailsjs/go/models";

export default {
  name: "DependencyCouplingDetails",
  props: {
    result: {
      required: true,
      type: Object as PropType<coupling.DependencyCoupling>
    }
  }
}
</script>

<style scoped>
.inner-details {
  margin-left: 50px;
  margin-right: 50px;
}
.coupling th {
  text-align: center;
}
.coupling td, .coupling th {
  padding: 0;
}
</style>
