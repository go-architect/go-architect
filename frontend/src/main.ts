import {createApp} from 'vue'
import VNetworkGraph from 'v-network-graph'
import App from './App.vue'
import 'jquery'
import "v-network-graph/lib/style.css"
import { router } from './router.js'

createApp(App)
    .use(router)
    .use(VNetworkGraph)
    .mount('#app')
