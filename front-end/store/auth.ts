import { defineStore } from 'pinia'

export const useAuth = defineStore('auth', {
  state: () => ({
    user: null,


  }),
  getters: {

    getUser(state) {
      return state.user
    }

  },
  actions: {

    login(body) {

      return Auth.login(body).then(res => {

        this.user = res.user;
        localStorage.setItem("token", res.token);


      })

    },
  },
})