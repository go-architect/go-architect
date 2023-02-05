<template>
  <div class="form-group col-md-4">
    <label for="idx">Select Number of Months for VCS Analysis</label>
    <div class="input-group input-group-sm">
      <input type="number" min="1" class="form-control" v-bind:value="range" v-on:input="input" />
      <span class="input-group-append">
        <button type="button" class="btn btn-info btn-flat" v-on:click="() => analyzeDependency(this)">Analyze!</button>
      </span>
    </div>
  </div>
</template>

<script lang="ts">
import {defineComponent} from "vue";

export default defineComponent({
  name: "DateRangeSelection",
  emits: ["run"],
  data() {
    return {
      range: 6,
    }
  },
  methods: {
    input($event: any) {
      this.range = parseInt($event.target.value)
    },
    analyzeDependency(vm: any) {
      vm.$emit("run", { range: this.range })
    }
  }
})
</script>

<style scoped>
</style>
