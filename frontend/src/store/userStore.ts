import { defineStore } from 'pinia';

export const useUserStore = defineStore('user', {
  state: () => ({
    user: null as { name: string; email: string } | null,
    token: null as string | null,
  }),

  actions: {
    setUser(userData: { name: string; email: string }) {
      this.user = userData;
    },

    setToken(token: string) {
      this.token = token;
    },

    clearUser() {
      this.user = null;
      this.token = null;
    },
  },

  // Enable persistence for this store
  persist: true,
});
