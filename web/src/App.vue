<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import HelloWorld from './components/HelloWorld.vue'
import { useCounterStore } from './stores/counter'

const counter = useCounterStore()
</script>

<template>
  <header>
    <img alt="Vue logo" class="logo" src="@/assets/logo.svg" width="125" height="125" />

    <HelloWorld msg="You did it!" />

    <div>Current Count: {{ counter.count }}</div>

    <nav>
      <RouterLink to="/">Home</RouterLink>
      <RouterLink to="/about">About</RouterLink>
      <RouterLink to="/counter">Counter</RouterLink>
    </nav>
  </header>

  <RouterView />
</template>

<style scoped></style>
