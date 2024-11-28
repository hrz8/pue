<script setup lang="ts">
import { useLangStore } from '../stores/lang'
import { supportedLangs } from '../constants/lang'

const langStore = useLangStore()

function changeLanguage(event: Event) {
  const selectedLang = (event.target as HTMLSelectElement).value
  langStore.setLang(selectedLang)
}
</script>

<template>
  <select
    class="form-select"
    aria-label="Select language"
    @change="changeLanguage"
  >
    <option
      v-for="langOpt in supportedLangs"
      :key="langOpt"
      :selected="langOpt === langStore.lang"
    >
      {{ langOpt }}
    </option>
  </select>
</template>
