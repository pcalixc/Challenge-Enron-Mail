<script setup lang="ts">
import { defineProps, computed } from 'vue'
interface Props {
  totalResults: number
  isLoading: boolean
  currentSearchTerm: string
  currentPage: number
}

const emailsByPage=8
const ini = computed(() => (props.currentPage  * emailsByPage) - emailsByPage)
const fin = computed(() => {
  let endValue = ini.value + emailsByPage
  if (endValue > props.totalResults) {
    endValue = props.totalResults
  }
  return endValue
})


const props = defineProps<Props>()
</script>

<template>
  <div class="flex flex-row align-middle justify-center">
    <p class="text-slate-800 text-xs dark:text-slate-200  mt-4">
      Showing 
      <span class="font-semibold text-vivid_blue dark:text-slate-50">{{ ini +1 }}</span>  to 
      <span class="font-semibold text-vivid_blue dark:text-slate-50">{{ fin }}</span>  of 
      <span class="font-semibold text-vivid_blue dark:text-slate-50">{{ props.totalResults }}</span> emails
  </p>
  </div>
</template>