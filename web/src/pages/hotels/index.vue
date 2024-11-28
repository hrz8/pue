<script setup lang="ts">
const hotels = [
  {
    id: '2YwwGCoTcmxbmVIa9rkP8R',
    name: 'Hotel Neo+ Green Savana',
  },
  {
    id: '4fFZlerhRzE0qnchBn9gpF',
    name: 'ASTON Sentul Resort & Conference Center',
  },
  {
    id: '6BTh2iaRRd6OyOouhR8c4h',
    name: 'ASTON Bellevue Radio Dalam',
  },
  {
    id: '6hTky1vmHTvX8Muqo0UdKO',
    name: 'favehotel Gatot Subroto',
  },
  {
    id: '7Hc1jLJEsldJhez23YgkH4',
    name: 'ASTON Kartika Grogol',
  }
]
</script>

<template>
  <div class="container">
    <div class="row row-cols-1 row-cols-md-2 g-4">
      <div v-for="hotel in hotels" :key="hotel.id" class="col">
        <div class="card">
          <div class="card-body">
            <h5 class="card-title">{{ hotel.name }}</h5>
            <RouterLink :to="`/hotels/${hotel.id}`" class="btn btn-primary">Details</RouterLink>
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
