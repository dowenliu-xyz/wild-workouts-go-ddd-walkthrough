<script setup lang="ts">
import {ref} from 'vue'
import {Attendee, getUserRole, Trainer} from '@/repositories/user'
import {Auth} from '@/repositories/auth'
import router from '@/router'

const userType = ref(getUserRole())

const signOut = () => {
  Auth.logout().finally(() => {
    router.push({name: 'login'})
  })
}
</script>

<template>
  <div class="d-flex flex-column flex-md-row align-items-center p-3 px-md-4 mb-3 bg-white border-bottom shadow-sm">
    <h5 class="my-0 me-md-auto font-weight-normal">Wild Workouts &#128023;</h5>
    <nav class="my-2 my-md-0 me-md-3">
      <router-link class="p-2 text-dark" :to="{ name: 'scheduleTraining' }" v-if="userType === Attendee">
        Schedule new training
      </router-link>
      <router-link class="p-2 text-dark" :to="{ name: 'trainingsList' }">Trainings</router-link>
      <router-link class="p-2 text-dark" :to="{ name: 'calendar' }">Calendar</router-link>
      <router-link class="p-2 text-dark" :to="{ name: 'setSchedule' }" v-if="userType === Trainer">Set
        schedule
      </router-link>
    </nav>
    <a class="btn btn-outline-primary" @click="signOut" href="/login">Logout</a>
  </div>
</template>

<style scoped>

</style>