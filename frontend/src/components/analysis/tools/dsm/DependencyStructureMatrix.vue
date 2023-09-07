<template>
  <div class="card">
    <div class="card-header">
      <h3 class="card-title">Dependency Structure Matrix</h3>
      <div class="card-tools">
        <router-link to="/analysis" class="btn btn-sm btn-primary" title="Back to Project Analysis">
          <i class="fa fa-circle-left" /> Back to Project Analysis
        </router-link>
      </div>
    </div>
    <div class="card-body">
      <div v-if="dsm.length === 0" class="overlay-wrapper" style="height: 300px;">
        <div class="overlay">
          <i class="fas fa-3x fa-sync-alt fa-spin"></i>
          <div class="text-bold pt-2">Loading Dependency Structure Matrix...</div>
        </div>
      </div>
      <div v-else class="dsm">
        <table class="table table-responsive table-bordered">
          <thead>
          <tr>
            <th>Package</th>
            <th class="package-id">#</th>
            <th v-for="(p, idx) in dsm" class="package-id" :title=p[0]>{{ idx+1 }}</th>
          </tr>
          </thead>
          <tbody>
          <tr v-for="(p, idx) in dsm">
            <td class="package-name" :title=p[0]>{{ wrapPackageName(p[0], mainPackage) }}</td>
            <td class="package-id" :title=p[0]>{{ idx+1 }}</td>
            <td v-for="idx2 in p.length-1" :class="resolveClasses(p, idx, idx2)">
              {{(idx===idx2-1) ? "": displayValue(p[idx2])}}
            </td>
          </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import {defineComponent} from "vue";
import {getSelectedProject} from "../../../../utils/storage";
import {GetDSM} from "../../../../../wailsjs/go/main/Api";
import {dsm} from "../../../../../wailsjs/go/models";

export default defineComponent({
  name: "DependencyGraph",
  props: {

  },
  data() {
    return {
      mainPackage: "",
      dsm: [] as string[][],
    }
  },
  methods: {
    mapToViewModel(matrix: dsm.DependencyStructureMatrix) {
      return matrix.packages.map((pkg: string, index: number) => {
        const packageEntry = [pkg]
        packageEntry.push(...matrix.dependencies[index].map((d: number) => ""+d))
        return packageEntry
      })
    },
    displayValue(value: string) {
      if(value!=="0") return value;
      return "";
    },
    resolveClasses(p: any, idx1: number, idx2: number) {
      const classes = [];
      if(idx1===idx2-1){
        classes.push("same-package")
      }
      if(p[idx2]>0){
        classes.push("dependency")
      }
      classes.push("dsm-cell")
      return classes.join(" ")
    },
    wrapPackageName(p: string, mainPackage: string) {
      if (p===mainPackage) return p

      return p.replace(mainPackage, "{{MainPackage}}")
    }
  },
  async mounted() {
    const project = await getSelectedProject()
    const dsm = await GetDSM(project!)
    this.dsm = this.mapToViewModel(dsm)
    this.mainPackage = project?.package!
  }
})
</script>

<style scoped>
.table td, .table th {
  padding: .3rem;
  text-align: center;
}

.dsm {
  font-size: 8px;
}
.package-id {
  width: 20px;
  font-weight: bold;
  text-align: center;
}
.package-name {
  text-align: left !important;
}
.dsm-cell {
  width: 20px;
  height: 20px;
}
.same-package {
  background-color: #dddddd;
}
.dependency {
   background-color: #ffebbb;
 }
</style>
