<script setup lang="ts">
import { onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { RouterLink } from 'vue-router'

import { useCatalogueStore } from '../../../stores/catalogue'

const route = useRoute()
const catalogue = useCatalogueStore()

onMounted(async () => {
  catalogue.setHotel(route.params.id as string)
})
</script>

<template>
  <div class="container">
    <h3>See rooms</h3>
    <RouterLink :to="`/hotels/${route.params.id}/rooms`">Hotels</RouterLink>
  </div>
</template>

<route lang="yaml">
meta:
  layout: blank
</route>
