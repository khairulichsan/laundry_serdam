<template>
  <div class="container">
    <div class="row justify-content-center">
      <div class="col-md-6">
        <div class="card">
          <div class="card-header">
            <h3>Register</h3>
          </div>
          <div class="card-body">
            <form @submit.prevent="registerUser">
              <div class="form-group">
                <label for="name">Name</label>
                <input
                  type="text"
                  class="form-control"
                  id="name"
                  v-model="name"
                  placeholder="Enter your name"
                  required
                />
              </div>


              <div class="form-group">
                <label for="email">Email</label>
                <input
                  type="email"
                  class="form-control"
                  id="email"
                  v-model="email"
                  placeholder="Enter your email"
                  required
                />
              </div>

              <div class="form-group">
                <label for="password">Password</label>
                <input
                  type="password"
                  class="form-control"
                  id="password"
                  v-model="password"
                  placeholder="Enter your password"
                  required
                />
              </div>

              <button type="submit" class="btn btn-primary w-100">Register</button>
            </form>
          </div>
          <div class="card-footer text-center">
            <p>
              Already have an account? <router-link to="/login">Login here</router-link>
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import api from "@/api";
import axios from "axios";
export default {
  data() {
    return {
      email: "",
      password: "",
    };
  },
  methods: {
    async registerUser() {
      const userData = {
        email: this.email,
        password: this.password,
      };

      try {
        const response = await api.post("/api/v1/register", userData);
        alert("Registration successful!");
        this.$router.push("/login"); // Redirect to login page after successful registration
      } catch (error) {
        console.error("Error during registration:", error.response.data);
        alert("Registration failed! Please try again.");
      }
    },
  },
};
</script>

<style scoped>
.container {
  margin-top: 50px;
}
.card {
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}
</style>
