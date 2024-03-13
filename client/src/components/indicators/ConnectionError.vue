<script setup lang="ts">
import { LottieAnimation } from 'lottie-web-vue'
import animationConnError from '@/assets/ConnError.json'
import { useEmailsStore } from '@/stores/emails'
const emailsStore = useEmailsStore()

let animationData: any = null 

import('@/assets/ConnError.json').then(module => {
  animationData = module.default 
    console.log(animationData)
  }).catch(error => {
    console.error('Failed to load animationData:', error)
  })

</script>

<template>
  <div
    v-if="emailsStore.ServerErrorResponse.errorStatus"
    class="lg:px-24 md:px-44 items-center flex h-[90vh] overflow-hidden justify-center flex-col-reverse lg:flex-row md:gap-28 gap-5"
  >
    <div class="xl:pt-0 w-44 xl:w-1/3  relative pb-12 lg:pb-0">
      <h1 class="font-mono ml-1 text-gray-800 font-bold text-2xl dark:text-slate-100">
        Whoops, something went wrong.
      </h1>

      <div class="flex mt-3 w-52 bg-red-100 rounded-lg p-4 mb-4 text-sm text-red-700" role="alert">
        <svg
          class="w-5 h-5 inline mr-3"
          fill="currentColor"
          viewBox="0 0 20 20"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            fill-rule="evenodd"
            d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z"
            clip-rule="evenodd"
          ></path>
        </svg>
        <div>
          <span class="font-medium">{{ emailsStore.ServerErrorResponse.errorMessage }}</span>
        </div>
      </div>
    </div>
    <lottie-animation
      :animation-data="animationConnError"
      :loop="true"
      class="w-[30rem] h-[30rem]"
    />
  </div>
</template>
