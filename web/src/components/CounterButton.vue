<script setup lang="ts">
import { useCounterStore } from '../stores/counter'

const counter = useCounterStore()
</script>

<template>
  <button @click="counter.increment">Increment</button>
  <button @click="counter.decrement">Decrement</button>
</template>
