<script setup lang="ts">
import WebMenu from '@/components/WebMenu.vue'
import WebFooter from '@/components/WebFooter.vue'
import {Auth} from '@/repositories/auth'
import router from '@/router'

document.body.className = 'app'

if (!Auth.isLoggedIn()) {
  router.push({name: 'login'})
}
</script>

<template>
  <div>
    <web-menu/>
    <div class="container">
      <slot></slot>
      <web-footer/>
    </div>
  </div>
</template>

<style scoped>

</style>