<template>
  <div class="position-relative tool-box">
    <div class="card bg-light d-flex flex-fill">
      <div class="card-header text-muted border-bottom-0">
        <h5>{{ tool.title }}</h5>
      </div>
      <div class="card-body pt-0">
        <div class="row">
          <div class="col-7">
            <p class="text-muted text-sm tool-description">
              {{ tool.description }}
            </p>
          </div>
          <div class="col-5 text-center">
            <img :src=tool.image alt="user-avatar" class="img-circle img-fluid tool-image">
          </div>
        </div>
      </div>
      <div class="card-footer">
        <div class="text-right">
          <router-link :to="tool.href" class="btn btn-sm btn-primary" :class="tool.soon==true ? 'disabled':''" v-on:click="tool.onclick">Open Tool</router-link>
        </div>
      </div>
    </div>
    <div v-if="tool.beta==true" class="ribbon-wrapper ribbon-sm">
      <div class="ribbon bg-warning text-xs">Beta</div>
    </div>
    <div v-if="tool.soon==true" class="ribbon-wrapper ribbon-md">
      <div class="ribbon bg-warning text-xs">Soon</div>
    </div>
  </div>
</template>

<script lang="ts">
export default {
  name: "AnalysisTool",
  props: {
    tool: {
      type: Object,
      required: true
    },
  }
}
</script>

<style scoped>
.tool-box {
  min-height: 100px;
}
.tool-description{
  height: 100px;
}

.tool-image {
  border: 1px solid #6c757d;
}
</style>
