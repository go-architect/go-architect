<template>
  <table class="instability table table-bordered">
    <thead>
    <tr>
      <th class="package-id">#</th>
      <th class="package-name">Package</th>
      <th class="metric">Afferent Coupling</th>
      <th class="metric">Efferent Coupling</th>
      <th class="metric">Abstractness</th>
      <th class="metric">Instability</th>
      <th class="metric">Distance from Diagonal</th>
    </tr>
    </thead>
    <tbody>
    <tr v-for="(pkg, idx) in packages">
      <td class="package-id">{{ idx+1 }}</td>
      <td>{{ pkg.package_name }}</td>
      <td class="metric">{{ pkg.afferent_coupling }}</td>
      <td class="metric">{{ pkg.efferent_coupling }}</td>
      <td class="metric">{{ pkg.abstractness }}</td>
      <td class="metric">{{ pkg.instability }}</td>
      <td class="metric">{{ pkg.distance }}</td>
    </tr>
    </tbody>
  </table>
</template>

<script lang="ts">
export default {
  name: "DependencyGraph",
  props: {
    packages: {
      type: Array,
      required: true
    }
  }
}
</script>

<style scoped>
.instability {
  margin: 0 auto;
  width: unset;
  font-size: 10px;
}
.package-id {
  width: 30px;
  font-weight: bold;
  text-align: center;
}
.package-name{
  min-width: 300px;
}
.metric {
  width: 150px;
  text-align: center;
}
</style>
