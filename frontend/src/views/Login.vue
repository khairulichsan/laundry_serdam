<template>
  <div class="container">
    <div class="row justify-content-center">
      <div class="col-md-6">
        <div class="card">
          <div class="card-header">
            <h3>Login</h3>
          </div>
          <div class="card-body">
            <form @submit.prevent="loginUser">
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

              <button type="submit" class="btn btn-primary w-100">Login</button>
            </form>
          </div>
          <div class="card-footer text-center">
            <p>
              Don't have an account? <router-link to="/register">Register here</router-link>
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
// Import the api.js instance for centralized axios configuration
import api from "@/api";

export default {
  data() {
    return {
      email: "",
      password: "",
    };
  },
  methods: {
    async loginUser() {
      const credentials = {
        email: this.email,
        password: this.password,
      };

      try {
        // Using the api instance from api.js to send the POST request
        const response = await api.post("/api/v1/login", credentials);

        // Store the JWT token in localStorage
        localStorage.setItem("token", response.data.token);
        alert("Login successful!");
        this.$router.push("/dashboard"); // Redirect to the dashboard or homepage after successful login
      } catch (error) {
        console.error("Error during login:", error.response.data);
        alert("Invalid credentials! Please try again.");
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
