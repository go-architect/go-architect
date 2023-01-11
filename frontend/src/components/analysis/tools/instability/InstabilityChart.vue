<template>
  <div class="chart-canvas">
    <svg id="chart-svg" width="100%" height="100%" viewBox="0,0,700,770" xmlns="http://www.w3.org/2000/svg">
      <defs>
        <marker id='head' orient="auto" markerWidth='4' markerHeight='4' refX='0.1' refY='2'>
          <path d='M0,0 V4 L2,2 Z' fill="black" />
        </marker>
      </defs>

      <rect width="600" height="600" x="20" y="100" fill="#e0ffcd" />
      <polygon points="20,220 20,700 500,700" style="fill:#ffebbb;" />
      <polygon points="140,100 620,100 620,580" style="fill:#ffebbb;" />
      <polygon points="20,340 20,700 380,700" style="fill:#ffcab0;" />
      <polygon points="260,100 620,100 620,460" style="fill:#ffcab0;" />

      <path
          id='abstractness-axis'
          marker-end='url(#head)'
          stroke-width='3'
          fill='none' stroke='black'
          d='M20,700, 20 60,120'
      />
      <path
          id='instability-axis'
          marker-end='url(#head)'
          stroke-width='3'
          fill='none' stroke='black'
          d='M20,700, 660 700,120'
      />
      <path
          id='diagonal'
          stroke-width='2'
          stroke-dasharray="5,5"
          fill='none' stroke='gray'
          d='M20,100, 620 700,120'
      />

      <text x="30" y="70" class="small">Abstractness</text>
      <text x="600" y="720" class="small">Instability</text>

      <circle v-for="(pkg, idx) in packages" :cx=calculateX(pkg.instability) :cy=calculateY(pkg.abstractness) r="5">
        <title>Package: {{ pkg }}</title>
      </circle>
    </svg>
  </div>
</template>

<script lang="ts">
export default {
  name: "DependencyGraph",
  props: {
    packages: {
      type: Array,
      required: true
    }
  },
  methods: {
    calculateX: (instability: number) => {
      return 20 + (instability * 600)
    },
    calculateY: (abstractness: number) => {
      return 700 - (abstractness * 600)
    }
  }
}

</script>

<style scoped>
.chart-canvas {
  margin-top: -40px;
  text-align: center;
}
</style>
