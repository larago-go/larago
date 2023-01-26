import '../css/app.css';

import { createApp, h } from 'vue';

import App from './App.vue';

import router from './router';

import AdminIndexJS from './Admin/index';

const app = createApp({

  render: () => h(App),

});

app.use(router);

app.use(AdminIndexJS);

app.mount('#app');
