import { ref, watch } from 'vue'
import { defineStore } from 'pinia'
import { ServiceFactory } from '@/services'
import type { Room } from '@/services/catalogue'
import { useLangStore } from './lang'

const service = ServiceFactory.createCatalogueService('http://localhost:8080')

export const useCatalogueStore = defineStore('catalogue', () => {
  const hotel = ref('')
  const rooms = ref<Room[]>([])
  const loading = ref(false)
  const error = ref<null | string>(null)
  const langStore = useLangStore()

  async function fetchRoomLists() {
    loading.value = true
    error.value = null

    try {
      const response = await service.listRooms(hotel.value, langStore.lang)
      rooms.value = response.rooms
    } catch (err) {
      if (err instanceof Error) {
        error.value = err.message || 'An error occurred while fetching room data.'
      }
    } finally {
      loading.value = false
    }
  }

  function setHotel(hotelId: string) {
    hotel.value = hotelId
  }

  watch(
    () => langStore.lang,
    () => {
      fetchRoomLists()
    }
  )

  return { hotel, rooms, loading, error, fetchRoomLists, setHotel };
});
