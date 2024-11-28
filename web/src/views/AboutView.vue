<script setup lang="ts">
import { useCounterStore } from '../stores/counter'

const counter = useCounterStore()
</script>

<template>
  <h1>This is an about page</h1>
  <div>Double Count: {{ counter.doubleCount }}</div>
</template>

<style></style>
