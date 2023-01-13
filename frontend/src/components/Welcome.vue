<script lang="ts">
import {BrowserOpenURL} from '../../wailsjs/runtime'

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
      <h2 class="card-title centered">Welcome to Go Architect</h2>
      <p class="card-text">
        Some quick example text to build on the card title and make up the bulk of the card's
        content.
      </p>
      <a class="card-link" v-on:click="goTo('https://www.go-architect.org')">Browse Documentation</a>
      <a class="card-link" v-on:click="goTo('https://www.github.com/fdaines/go-architect')">Github Repository</a>
      <a class="card-link" v-on:click="goTo('https://www.github.com/fdaines/go-architect/issues')">Report an Issue</a>
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
