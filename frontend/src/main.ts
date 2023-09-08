import {createApp} from 'vue'
import App from './App.vue'
import 'jquery'
import { router } from './router.js'
// @ts-ignore
import VueHighlightJS from 'vue3-highlightjs'

createApp(App)
    .use(router)
    .use(VueHighlightJS)
    .mount('#app')
