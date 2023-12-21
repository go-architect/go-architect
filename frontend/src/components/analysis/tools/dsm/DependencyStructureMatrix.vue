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
      <div class="row">
        <div class="col-4">
          <div class="form-group">
            <div class="custom-control custom-switch">
              <input type="checkbox" class="custom-control-input" id="weighted-dsm" v-on:change="switchShowWeighted">
              <label class="custom-control-label" for="weighted-dsm">Show Weighted DSM</label>
            </div>
          </div>
          <div class="form-group">
            <div class="custom-control custom-switch">
              <input type="checkbox" class="custom-control-input" id="aggregate-values" v-on:change="switchShowAggregated">
              <label class="custom-control-label" for="aggregate-values">Aggregate by Package Layer</label>
            </div>
          </div>
        </div>
        <div class="col-8">
          <table class="table table-bordered">
            <thead>
            <tr>
              <th colspan="8">Package Layers</th>
            </tr>
            </thead>
            <tbody>
            <tr>
              <td>Internal Packages</td>
              <td>Organization Packages</td>
              <td>External Packages</td>
              <td>Standard Packages</td>
            </tr>
            <tr>
              <td class="legend internal-package"></td>
              <td class="legend organization-package"></td>
              <td class="legend external-package"></td>
              <td class="legend standard-package"></td>
            </tr>
            </tbody>
          </table>
        </div>
      </div>
      <div v-if="dsm.length === 0" class="overlay-wrapper" style="height: 300px;">
        <div class="overlay">
          <i class="fas fa-3x fa-sync-alt fa-spin"></i>
          <div class="text-bold pt-2">Loading Dependency Structure Matrix...</div>
        </div>
      </div>
      <div v-else class="dsm">
        <table :class="resolveDSMTableClasses()">
          <thead>
          <tr>
            <th class="package-id">#</th>
            <th>Package</th>
            <th v-for="(p, idx) in dsm" class="package-id" :title=p[1]>{{ idx+1 }}</th>
          </tr>
          </thead>
          <tbody>
          <tr v-for="(p, idx) in dsm">
            <td class="package-id" :title=p[1]>{{ idx+1 }}</td>
            <td class="package-name" :title=p[1] :class="resolveEntryClass(p)">{{ wrapPackageName(p[1], mainPackage) }}</td>
            <td v-for="idx2 in p.length-2" :class="resolveClasses(p, idx, idx2)">
              {{(idx===idx2-1) ? "": displayValue(p[idx2+1])}}
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
import {resolvePackageLayer, resolvePackageLayerIndex} from "../../../../utils/packages";

export default defineComponent({
  name: "DependencyGraph",
  props: {

  },
  data() {
    return {
      mainPackage: "",
      weightedDSM: {packages: [] as string[]} as dsm.DependencyStructureMatrix,
      dsm: [] as string[][],
      orgPackages: [] as string[],
      showWeighted: false,
      showAggregated: false,
    }
  },
  methods: {
    switchShowWeighted($event: any) {
      this.showWeighted = !this.showWeighted
      this.showDSM()
    },
    switchShowAggregated($event: any) {
      this.dsm = []
      this.showAggregated = !this.showAggregated
      this.showDSM()
    },
    async showDSM() {
      const project = await getSelectedProject()
      this.mainPackage = project?.package!
      this.dsm = this.mapToViewModel(this.weightedDSM)
    },
    reduceToBoolean(matrix: dsm.DependencyStructureMatrix) {
      if(this.showWeighted) return matrix
      return {
        module: matrix.module,
        packages: matrix.packages,
        dependencies: matrix.dependencies.map((value: number[]) => {
          return value.map((v: number) => v > 0 ? 1 : 0)
        })
      }
    },
    aggregateValues(matrix: dsm.DependencyStructureMatrix) {
      if(!this.showAggregated) return matrix

      const aggregatedDependencies = [[0,0,0,0],[0,0,0,0],[0,0,0,0],[0,0,0,0]]
      matrix.dependencies.forEach((values: number[], index: number) => {
        const destinationLayer = resolvePackageLayerIndex(
            resolvePackageLayer(matrix.packages[index], this.mainPackage, this.orgPackages))
        values.forEach((value: number, idx: number) => {
          const originLayer = resolvePackageLayerIndex(
              resolvePackageLayer(matrix.packages[idx], this.mainPackage, this.orgPackages))
         aggregatedDependencies[destinationLayer][originLayer] += value
        })
      })


      const aggregatedMatrix: dsm.DependencyStructureMatrix = {
        module: matrix.module,
        packages: ["Internal Packages", "Organization Packages", "External Packages", "Standard Packages"],
        dependencies: aggregatedDependencies
      }

      return aggregatedMatrix
    },
    mapToViewModel(matrix: dsm.DependencyStructureMatrix) {
      const processed = this.reduceToBoolean(this.aggregateValues(matrix))
      return processed.packages.map((entryName: string, index: number) => {
        const packageEntry = [resolvePackageLayer(entryName, this.mainPackage, this.orgPackages), entryName]
        packageEntry.push(...processed.dependencies[index].map((d: number) => ""+d))
        return packageEntry
      })
    },
    displayValue(value: string) {
      if(value!=="0") return value;
      return "";
    },
    isOrganizationPackage(pkg: string) {
      return this.orgPackages.some(op => pkg.startsWith(op))
    },
    resolveEntryClass(p: any) {
      const classes = [];
      classes.push(p[0]+"-package")
      return classes
    },
    resolveClasses(p: any, idx1: number, idx2: number) {
      const classes = [];
      if(p[0]==="internal" && this.dsm[idx2-1][0] === "internal") {
        classes.push("internal-package")
      } else if(p[0]==="organization" && idx2<=idx1) {
        classes.push("organization-package")
      } else if(this.dsm[idx2-1][0]==="organization" && idx1<idx2) {
        classes.push("organization-package")
      } else if(p[0]==="standard" || this.dsm[idx2-1][0]==="standard") {
        classes.push("standard-package")
      } else if(p[0]==="external" || this.dsm[idx2-1][0]==="external") {
        classes.push("external-package")
      }
      classes.push("dsm-cell")
      if(this.showAggregated){
        classes.push("aggregated-cell")
      }
      if(idx1===idx2-1){
        classes.push("same-package")
      } else if(p[idx2+1]>0){
        classes.push("dependency")
      }
      return classes.join(" ")
    },
    resolveDSMTableClasses() {
      const classes = ["table", "table-responsive", "table-bordered"]
      if(this.showAggregated) classes.push("aggregated-dsm")
      return classes.join(" ")
    },
    wrapPackageName(p: string, mainPackage: string) {
      if (p===mainPackage) return p

      return p.replace(mainPackage, "{{MainPackage}}")
    }
  },
  async mounted() {
    const project = await getSelectedProject()
    this.weightedDSM = await GetDSM(project!)
    this.mainPackage = project?.package!
    this.orgPackages = project?.organization_packages ? project?.organization_packages! : []
    this.dsm = this.mapToViewModel(this.weightedDSM)
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
  background-color: #dddddd !important;
}
.dependency {
  background-color: #BEC1DB !important;
  font-weight: bolder;
}
.internal-package {
  background-color: #EDF9FF;
}
.organization-package {
  background-color: #E7FEE2;
}
.external-package {
  background-color: #FEF4F4;
}
.standard-package {
  background-color: #FFFBDB;
}
.card-body {
  font-size: 10px;
}
.legend {
  width: 25%;
  height: 20px;
}
.custom-control-label {
  padding-top: 5px;
}

.aggregated-dsm {
  font-size: 20px;
}
.aggregated-dsm td {
  vertical-align: middle;
}
.aggregated-dsm .package-id {
  width: 30px;
}
.aggregated-cell {
  width: 75px;
  height: 75px;
  vertical-align: middle;
}
</style>
