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
      <img src="/images/logo.png" alt="" class="logo" />
      <h2 class="card-title centered">Welcome to Go Architect</h2>
      <p class="card-text">
        An Architecture Analysis Tool for Golang Projects.
      </p>
      <a class="card-link" v-on:click="goTo('https://go-architect.github.io/')">Browse Documentation</a>
      <a class="card-link" v-on:click="goTo('https://www.github.com/go-architect/go-architect')">Github Repository</a>
      <a class="card-link" v-on:click="goTo('https://www.github.com/go-architect/go-architect/issues')">Report an Issue</a>
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

.logo {
  width: 200px;
  margin-bottom: 10px;
  margin-left: -30px;
}
</style>
