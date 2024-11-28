import { ref } from 'vue'
import { defineStore } from 'pinia'
import { ServiceFactory } from '@/services'
import type { Room } from '@/services/catalogue'

const service = ServiceFactory.createCatalogueService("http://localhost:8080")

export const useCatalogueStore = defineStore('catalogue', () => {
  const rooms = ref<Room[]>([])
  const loading = ref(false)
  const error = ref<null | string>(null)

  async function fetchRoomLists(hotelId: string) {
    loading.value = true
    error.value = null

    try {
      const response = await service.listRooms(hotelId)
      rooms.value = response.rooms
    } catch (err) {
      if (err instanceof Error) {
        error.value = err.message || 'An error occurred while fetching user data.'
      }
    } finally {
      loading.value = false
    }
  }

  return { rooms, loading, error, fetchRoomLists }
})
