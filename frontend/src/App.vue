<template>
  <div>
    <nav>
      <router-link to="/manager/dashboard">Dashboard</router-link>
      <router-link to="/manager/services">Manage Services</router-link>
      <router-link to="/manager/promotions">Promotions</router-link>
      <router-link to="/manager/attendance">Employee Attendance</router-link>
    </nav>
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  name: 'Home',
}
</script>
<!--
// src/components/Manager/Dashboard.vue
<template>
  <div>
    <h1>Manager Dashboard</h1>
    <div>
      <p>Total Revenue: {{ revenue }}</p>
      <p>Total Orders: {{ totalOrders }}</p>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Dashboard',
  data() {
    return {
      revenue: 0,
      totalOrders: 0,
    }
  },
  created() {
    // Fetch dashboard data
    this.revenue = 100000 // Example value
    this.totalOrders = 50 // Example value
  },
}
</script> -->

<!-- // src/components/Manager/ManageServices.vue
<template>
  <div>
    <h1>Manage Services</h1>
    <form @submit.prevent="addService">
      <input v-model="serviceName" placeholder="Service Name" required />
      <input v-model.number="price" placeholder="Price Per Kilo" type="number" required />
      <input v-model.number="duration" placeholder="Duration (hours)" type="number" required />
      <button type="submit">Add Service</button>
    </form>
    <ul>
      <li v-for="service in services" :key="service.id">
        {{ service.name }} - {{ service.price }} per kilo, {{ service.duration }} hours
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  name: "ManageServices",
  data() {
    return {
      serviceName: "",
      price: 0,
      duration: 0,
      services: [],
    };
  },
  methods: {
    addService() {
      const newService = {
        id: Date.now(),
        name: this.serviceName,
        price: this.price,
        duration: this.duration,
      };
      this.services.push(newService);
      this.serviceName = "";
      this.price = 0;
      this.duration = 0;
    },
  },
};
</script>

// src/router/index.js
import Vue from "vue";
import VueRouter from "vue-router";
import ManagerView from "../views/ManagerView.vue";
import Dashboard from "../components/Manager/Dashboard.vue";
import ManageServices from "../components/Manager/ManageServices.vue";
import Promotions from "../components/Manager/Promotions.vue";
import Attendance from "../components/Manager/Attendance.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/manager",
    component: ManagerView,
    children: [
      { path: "dashboard", component: Dashboard },
      { path: "services", component: ManageServices },
      { path: "promotions", component: Promotions },
      { path: "attendance", component: Attendance },
    ],
  },
  // Define routes for EmployeeView and CustomerView here
];

const router = new VueRouter({
  mode: "history",
  routes,
});

export default router;

// main.js
import Vue from "vue";
import App from "./App.vue";
import router from "./router";

Vue.config.productionTip = false;

new Vue({
  router,
  render: (h) => h(App),
}).$mount("#app"); -->
