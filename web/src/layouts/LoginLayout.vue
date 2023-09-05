<script setup lang="ts">
document.body.className = 'login'
</script>

<template>
  <div>
    <slot></slot>
  </div>
</template>

<style>
body.login {
  display: flex;
  align-items: center;
  justify-content: center;
  padding-top: 40px;
  padding-bottom: 40px;
  text-align: center;
}

.form-label-group input::placeholder {
  color: transparent;
}

.form-label-group input:not(:placeholder-shown) ~ label {
  font-size: 12px;
  color: #777;
}
</style>