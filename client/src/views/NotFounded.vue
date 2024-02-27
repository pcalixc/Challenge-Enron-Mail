<script setup lang="ts">
import { LottieAnimation } from 'lottie-web-vue'

import animationConnError from '@/assets/404.json'
</script>

<template>
  <div
    class="lg:px-24 md:px-44 items-center flex justify-center h-[100vh] lg:flex-row md:gap-28 gap-5"
  >
    <lottie-animation
      :animation-data="animationConnError"
      :loop="true"
      class="w-[20rem] h-[20rem]"
    />
  </div>
</template>