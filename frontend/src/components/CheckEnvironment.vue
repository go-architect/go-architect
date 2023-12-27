<script lang="ts">
import {BrowserOpenURL} from "../../wailsjs/runtime";

export default {
  methods: {
    goTo(url: string) {
      BrowserOpenURL(url)
    }
  }
}
</script>

<template>
  <div class="card">
    <div class="card-body text-center">
      <h2 class="card-title centered">Warning</h2>
      <p class="card-text">
        It seems the GO executable directory is not set.
        Please check <a class="card-link" v-on:click="goTo('https://go-architect.github.io/docs/install#go-directory-is-not-set')">Go-Architect's Documentation</a> to set the content of the
        <i>~/.goarchitect/environment</i> file and restart the application.
      </p>
    </div>
  </div>
</template>

<style scoped>
.card-title {
  float: none;
  text-align: center;
  font-size: 2.4rem;
  padding-bottom: 10px;
}
</style>
