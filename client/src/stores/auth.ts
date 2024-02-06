import { defineStore } from 'pinia';

interface AuthState {
    authToken: string | null;
  }


  export const UseAuthStore = defineStore({
    id: 'auth',
    state: (): AuthState => ({
      authToken: localStorage.getItem('jwt-token') || null,
    }),
    actions: {
      setAuthToken(token: string) {
        this.authToken = token;
        localStorage.setItem('jwt-token', token);
      },
      removeAuthToken() {
        this.authToken = null;
        localStorage.removeItem('jwt-token');
      },
    },
    getters: {
      //isAuthenticated: () => !!this.authToken,
    },
  });
