<script setup lang="ts">
import { LottieAnimation } from 'lottie-web-vue'
import animationNotData from '@/assets/se2.json'
</script>

<template>
  <div class="h-full flex justify-center align-middle mt-24">
    <lottie-animation :animation-data="animationNotData" :loop="true" class="w-72 h-72" />
  </div>
</template>
