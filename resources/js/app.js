/**
 * First we will load all of this project's JavaScript dependencies which
 * includes Vue and other libraries. It is a great starting point when
 * building robust, powerful web applications using Vue and Laravel.
 */

import "../css/app.css";

require('./bootstrap');

import { createApp, h } from 'vue'

import App from './App'

import router from './router'

import admin_index_js from './Admin/index'

/**
 * The following block of code may be used to automatically register your
 * Vue components. It will recursively scan this directory for the Vue
 * components and automatically register them with their "basename".
 *
 * Eg. ./components/ExampleComponent.vue -> <example-component></example-component>

// const files = require.context('./', true, /\.vue$/i);
// files.keys().map(key => Vue.component(key.split('/').pop().split('.')[0], files(key).default));
 */

const app = createApp({

  render: ()=>h(App)

})

app.use(router)

app.use(admin_index_js)

app.mount('#app')

