<script setup lang="ts">
import { onMounted } from 'vue'

import WelcomeItem from '../components/WelcomeItem.vue'
import DocumentationIcon from '../components/icons/IconDocumentation.vue'
import SupportIcon from '../components/icons/IconSupport.vue'

onMounted(() => {
  console.log('Halo')
})
</script>

<template>
  <main>
    <WelcomeItem>
      <template #icon>
        <DocumentationIcon />
      </template>
      <template #heading>Documentation</template>
      <div>Content documentation</div>
    </WelcomeItem>

    <WelcomeItem>
      <template #icon>
        <SupportIcon />
      </template>
      <template #heading>Support</template>
      <div>Content support</div>
    </WelcomeItem>
  </main>
</template>
