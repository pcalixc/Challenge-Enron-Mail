<script setup lang="ts">
import { useEmailsStore } from '@/stores/emails';
import { computed } from 'vue'

const emailsStore = useEmailsStore()

const emailsByPage=8
const ini = computed(() => (emailsStore.currentPage  * emailsByPage) - emailsByPage)
const fin = computed(() => {
  let endValue = ini.value + emailsByPage
  if (endValue > emailsStore.totalResults) {
    endValue = emailsStore.totalResults
  }
  return endValue
})

</script>

<template>
  <div v-if="!emailsStore.ServerErrorResponse.errorStatus && emailsStore.totalResults !=0" class="flex flex-row align-middle justify-center">
    <p class="text-slate-800 text-xs dark:text-slate-200  mt-4">
      Showing 
      <span class="font-semibold text-vivid_blue dark:text-slate-50">{{ ini +1 }}</span>  to 
      <span class="font-semibold text-vivid_blue dark:text-slate-50">{{ fin }}</span>  of 
      <span class="font-semibold text-vivid_blue dark:text-slate-50">{{ emailsStore.totalResults }}</span> emails
  </p>
  </div>
</template>