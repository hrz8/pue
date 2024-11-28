<script setup lang="ts">
defineProps<{
  msg: string
}>()
</script>

<template>
  <h1>{{ msg }}</h1>
  <h3>Hello, world!</h3>
</template>

<style scoped></style>
