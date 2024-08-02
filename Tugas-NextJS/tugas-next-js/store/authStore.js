import create from 'zustand'
import { persist } from 'zustand/middleware'

const useAuthStore = create(persist(
  (set, get) => ({
    user: null,
    token: null,
    login: (user, token) => set({ user, token }),
    logout: () => set({ user: null, token: null })
  }),
  {
    name: 'auth-storage', // nama unik untuk penyimpanan di localStorage
    getStorage: () => localStorage // menentukan tempat penyimpanan
  }
));

export default useAuthStore;
