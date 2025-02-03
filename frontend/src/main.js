import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import { createPinia } from 'pinia'
import { router } from './route'

import { library } from '@fortawesome/fontawesome-svg-core';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { fas } from '@fortawesome/free-solid-svg-icons';
import { fab } from '@fortawesome/free-brands-svg-icons';

import { QuillEditor } from '@vueup/vue-quill'
import '@vueup/vue-quill/dist/vue-quill.snow.css';


library.add(fas, fab);
const pinia = createPinia()

const app = createApp(App)
app.use(router)
app.use(pinia)
app.component('QuillEditor', QuillEditor)
app.component('font-awesome-icon', FontAwesomeIcon)
app.mount('#app')
