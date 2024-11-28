<script setup lang="ts">
import { useCounterStore } from '../stores/counter'

const counter = useCounterStore()
</script>

<template>
  <button class="btn btn-primary" @click="counter.increment">Increment</button>
  <button class="btn btn-primary" @click="counter.decrement">Decrement</button>
</template>
