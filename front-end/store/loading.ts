import { defineStore } from 'pinia'

export const useLoading = defineStore('loading', {
    state: () => ({
        loading: {},


    }),
    getters: {

        getLoading(state) {
            return state.loading
        }

    },

})