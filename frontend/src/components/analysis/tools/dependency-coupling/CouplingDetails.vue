<template>
  <div class="card card-primary card-outline">
    <div class="card-header">
      <h3 class="card-title">Coupling Analysis Result</h3>
    </div>
    <div class="card-body">
      <ul class="list-group list-group-unbordered mb-3">
        <li class="list-group-item">
          <b>Dependency</b>
          <div class="float-right">{{ result?.dependency }}</div>
        </li>
        <li class="list-group-item">
          <b>Overall Coupling Level</b>
          <div class="float-right">
            <div class="badge badge-warning" style="padding: 6px 6px 0 6px;">
              <h6>{{ result?.coupling_level.toLocaleString('en-us', {}) }}</h6>
            </div>
          </div>
        </li>
        <li class="list-group-item">
          <div class="inner-details">

            <div v-for="pkgCoupling in result?.details" class="card card-outline card-default collapsed-card">
              <div class="card-header">
                <h3 class="card-title bold">
                  <small class="badge badge-warning" style="margin-right: 10px;">
                    Coupling Level: {{ pkgCoupling.coupling_level.toLocaleString('en-us', {}) }}
                  </small>
                  {{ pkgCoupling.package }}
                </h3>
                <div class="card-tools">
                  <button type="button" class="btn btn-tool" data-card-widget="collapse" title="Open Details">
                    <i class="fas fa-plus"></i>
                  </button>
                </div>
              </div>
              <div class="card-body" style="display: none;">
                <div class="inner-details">
                  <div v-for="fileCoupling in pkgCoupling.details" class="card card-outline card-default collapsed-card">
                    <div class="card-header">
                      <h3 class="card-title bold">
                        <small class="badge badge-warning" style="margin-right: 10px;">
                          Coupling Level: {{ fileCoupling.coupling_level.toLocaleString('en-us', {}) }}
                        </small>
                        {{ fileCoupling.file }}
                      </h3>
                      <div class="card-tools">
                        <button type="button" class="btn btn-tool" data-card-widget="collapse" title="Open Details">
                          <i class="fas fa-plus"></i>
                        </button>
                      </div>
                    </div>
                    <div class="card-body" style="display: none;">
                      <div class="inner-details">
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
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>

<script lang="ts">
import {PropType} from "vue";
import {coupling} from "../../../../../wailsjs/go/models";
import {BrowserOpenURL} from "../../../../../wailsjs/runtime";

export default {
  name: "DependencyCouplingDetails",
  props: {
    result: {
      required: true,
      type: Object as PropType<coupling.DependencyCoupling>
    }
  },
  methods: {
    openDetails(pkg: string) {
      console.log("Open Details for: " + pkg)
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
  padding: 0 0 0 10px;
}
</style>
