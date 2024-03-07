<script setup lang="ts">
import { LottieAnimation } from 'lottie-web-vue'

import animationConnError from '@/assets/ConnectionError.json'
</script>

<template>
  <div
    class="lg:px-24 md:px-44 items-center flex justify-center flex-col-reverse lg:flex-row md:gap-28 gap-5"
  >
    <div class="xl:pt-0 w-44 xl:w-1/3 relative pb-12 lg:pb-0">
      <h1 class="font-mono text-gray-800 font-bold text-2xl dark:text-slate-100">
        Sorry, we couldn't complete the request at this time.
      </h1>
    
    </div>
    <lottie-animation
      :animation-data="animationConnError"
      :loop="true"
      class="w-[35rem] h-[35rem]"
    />
  </div>
</template>
