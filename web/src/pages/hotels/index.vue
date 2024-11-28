<script setup lang="ts">
import { onMounted } from 'vue'
import { useCatalogueStore } from '../../stores/catalogue'

const catalogue = useCatalogueStore()

const hotelId = '6BTh2iaRRd6OyOouhR8c4h'

onMounted(async () => {
  await catalogue.fetchRoomLists(hotelId)
  console.log(catalogue.rooms)
})
</script>

<template>
  <ul>
    <li>ABC</li>
    <li>DEF</li>
    <li>GHI</li>
  </ul>
</template>

<route lang="yaml">
meta:
  layout: blank
</route>
