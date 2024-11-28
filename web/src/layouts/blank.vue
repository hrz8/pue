<script  setup lang="ts">
import { onMounted } from 'vue'
import LangSwitcher from '../components/LangSwitcher.vue'

onMounted(() => {
  console.log('layouts/blank')
})
</script>

<template>
  <div class="container my-5">
    <RouterLink to="/" class="btn btn-primary">Home</RouterLink>
    <LangSwitcher/>
  </div>
  <RouterView />
</template>
