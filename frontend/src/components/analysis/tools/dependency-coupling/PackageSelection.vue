<template>
  <div class="form-group col-md-4">
    <label for="idx">Dependency to Analyze</label>
    <div class="input-group input-group-sm">
      <input type="text" class="form-control" v-bind:value="dependency" v-on:input="input" />
      <span class="input-group-append">
        <button type="button" class="btn btn-info btn-flat" v-on:click="() => analyzeDependency(this)">Analyze!</button>
      </span>
    </div>
  </div>
</template>

<script lang="ts">
import {defineComponent} from "vue";

export default defineComponent({
  name: "DependencyGraph",
  emits: ["run"],
  data() {
    return {
      dependency: "",
    }
  },
  methods: {
    input($event: any) {
      this.dependency = $event.target.value
    },
    analyzeDependency(vm: any) {
      this.dependency = this.dependency.toLowerCase()
      vm.$emit("run", { dependency: this.dependency })
    }
  }
})
</script>

<style scoped>
</style>
