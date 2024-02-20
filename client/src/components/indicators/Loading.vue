<!-- eslint-disable vue/multi-word-component-names -->
<script setup lang="ts">
import { LottieAnimation } from 'lottie-web-vue'
import animationNotData from '@/assets/Loading.json'
</script>

<template>
  <div class="h-full flex justify-center align-middle mt-24">
    <lottie-animation :animation-data="animationNotData" :loop="true" class="w-60 h-60" />
  </div>
</template>