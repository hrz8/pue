import { ref } from 'vue'
import { defineStore } from 'pinia'
import { supportedLangs } from '@/constants/lang'

export const useLangStore = defineStore('lang', () => {
  const _lang = localStorage.getItem('lang') ?? 'en-US'
  const lang = ref(_lang)

  function setLang(newLang: string) {
    if (!supportedLangs.includes(newLang)) {
      return
    }
    lang.value = newLang
    localStorage.setItem('lang', newLang)
  }

  return { lang, setLang }
})
