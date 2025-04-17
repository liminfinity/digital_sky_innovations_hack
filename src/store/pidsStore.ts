import { defineStore } from 'pinia'

export const usePidsStore = defineStore('pids', {
  state: () => ({
    pids: {},
  }),
})
