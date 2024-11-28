<script setup lang="ts">
import { onMounted } from 'vue'
import { useRoute } from 'vue-router'

import { useCatalogueStore } from '../../../stores/catalogue'

const route = useRoute()
const catalogue = useCatalogueStore()

onMounted(async () => {
  catalogue.setHotel(route.params.id as string)
  await catalogue.fetchRoomLists()
})
</script>

<template>
  <div class="container">
    <div class="row row-cols-1 row-cols-md-2 g-4">
      <div v-for="room in catalogue.rooms" :key="room.id" class="col">
        <div class="card">
          <img v-if="room.images?.length > 0" :src="room.images[0]" class="card-img-top" :alt="room.hotel.slug">
          <div class="card-body">
            <h5 class="card-title">{{ room.roomType }}</h5>
            <p class="card-text">{{ room.description }}</p>
            <a href="#" class="btn btn-primary">Details & Book</a>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<route lang="yaml">
meta:
  layout: blank
</route>
