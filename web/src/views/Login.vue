<script setup lang="ts">
import {onMounted, ref} from 'vue'
import {Auth} from '@/repositories/auth'
import router from '@/router'
import LoginLayout from '@/layouts/LoginLayout.vue'
import {ElNotification} from 'element-plus'
import 'element-plus/es/components/notification/style/css'
import {getTestUsers, loginUser} from '@/repositories/user'

onMounted(() => {
  if (Auth.isLoggedIn()) {
    router.push({name: 'trainingsList'})
  }
})

interface FormModel {
  login: string
  password: string
}

const form = ref(<FormModel>{})
const showLoader = ref(false)

const submit = (): void => {
  showLoader.value = true
  loginUser(form.value.login, form.value.password).then(() => {
    ElNotification({
      message: 'Hey buddy!',
      type: 'success',
    })
    router.push({name: 'trainingsList'})
  }).catch((error) => {
    ElNotification({
      message: 'Failed to log in',
      type: 'error',
    })
    console.error(error)
    showLoader.value = false
  })
}

const loadCredentials = (event: Event) => {
  if (!(event.target instanceof HTMLElement)) {
    return
  }
  const target = event.target as HTMLElement
  form.value.login = target.getAttribute('data-login') ?? ''
  form.value.password = target.getAttribute('data-password') ?? ''
}
</script>

<template>
  <login-layout>
    <form class="form-signin" @submit.prevent="submit">
      <span style="font-size: 60px;">&#128023;</span>
      <h1 class="h3 mb-3 font-weight-normal">Please sign in</h1>
      <div class="alert alert-primary" role="alert">
        <p v-for="user in getTestUsers()" :key="user.login">
          {{ user.role.charAt(0).toUpperCase() + user.role.slice(1) }} credentials
          <a href="#" :data-login="user.login" :data-password="user.password"
             @click="loadCredentials" :key="user.login">{{ user.login }}:{{ user.password }}</a>
        </p>
      </div>

      <label for="inputEmail" class="visually-hidden">Email address</label>
      <input type="email" id="inputEmail" class="form-control" v-model="form.login" placeholder="Email address"
             required autofocus>
      <label for="inputPassword" class="visually-hidden">Password</label>
      <input type="password" id="inputPassword" class="form-control" v-model="form.password" placeholder="Password"
             required>
      <div class="checkbox mb-3">
        <label>
          <input type="checkbox" value="remember-me"> Remember me
        </label>
      </div>
      <button class="btn btn-lg btn-primary btn-block" type="submit">
        Sign in
        <span v-if="showLoader" class="spinner-grow spinner-grow-sm" style="width: 1.3rem; height: 1.3rem;"
              role="status" aria-hidden="true"></span>
      </button>
      <p class="mt-5 mb-3 text-muted">&copy; 2017-2020</p>
    </form>
  </login-layout>
</template>

<style scoped>
.form-signin {
  width: 100%;
  max-width: 330px;
  padding: 15px;
  margin: 0 auto;
}

.form-signin .checkbox {
  font-weight: 400;
}

.form-signin .form-control {
  position: relative;
  box-sizing: border-box;
  height: auto;
  padding: 10px;
  font-size: 16px;
}

.form-signin .form-control:focus {
  z-index: 2;
}

.form-signin input[type="email"] {
  margin-bottom: -1px;
  border-bottom-right-radius: 0;
  border-bottom-left-radius: 0;
}

.form-signin input[type="password"] {
  margin-bottom: 10px;
  border-top-left-radius: 0;
  border-top-right-radius: 0;
}
</style>