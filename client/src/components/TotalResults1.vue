<script setup lang="ts">
import { defineProps, computed } from 'vue'
interface Props {
  totalResults: number
  isLoading: boolean
  currentSearchTerm: string
  currentPage: number
  sugerencias: string[]
  getData: Function
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
  <div class="flex flex-row align-middle justify-center mt-8 ">
    <div v-if="props.totalResults==0 && !isLoading" class="text-slate-700 text-xs dark:text-slate-200">
      <div class="flex bg-yellow-100 rounded-lg p-4 mb-4 text-sm text-yellow-700" role="alert">
        <svg class="w-5 h-5 inline mr-3" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z" clip-rule="evenodd"></path></svg>
        <div>
            <span class="font-medium">No emails matched the term "{{ currentSearchTerm }}" </span>
            <span> Did you mean 
              <span @click="props.getData(1,props.sugerencias[0] )" class="underline font-semibold cursor-pointer">{{ props.sugerencias[0] }}</span>
              ?</span>
        </div>
      </div>
    </div>
    <p v-else class="text-slate-800 text-xs dark:text-slate-200  mt-4">
      Showing <span class="font-semibold text-vivid_blue dark:text-slate-50">{{ ini +1 }}</span>  to <span class="font-semibold text-vivid_blue dark:text-slate-50">{{ fin }}</span>  of <span class="font-semibold text-vivid_blue dark:text-slate-50">{{ props.totalResults }}</span> emails
  </p>
  </div>
</template>