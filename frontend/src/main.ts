import { createApp } from 'vue';
import App from './App.vue';
import { createPinia } from 'pinia';
import piniaPersist from 'pinia-plugin-persistedstate';
import router from './router';
import 'bootstrap/dist/css/bootstrap.min.css';


const app = createApp(App);
const pinia = createPinia();

// Add persisted state plugin to Pinia
pinia.use(piniaPersist);

app.use(router);
app.use(pinia);
app.mount('#app');
