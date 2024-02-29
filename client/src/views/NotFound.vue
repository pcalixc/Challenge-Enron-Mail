<script setup lang="ts">
import { LottieAnimation } from 'lottie-web-vue'
import animationConnError from '@/assets/404error.json'
import { useRouter } from 'vue-router'
const router = useRouter()

const goToHome=()=>{
  router.push('/')
}
</script>

<template>
  <div
    class="lg:px-24 items-center -mt-15 flex flex-col justify-center  h-[100vh] lg:flex-col gap-10"
  >
    <lottie-animation
      :animation-data="animationConnError"
      :loop="true"
      class="w-[30rem] h-[30rem]"
    />
    <button
    @click="goToHome"
      class="px-2 py-1 my-24 font-mono -mt-32 font-semibold dark:bg-slate-900 animate-pulse dark:text-slate-300  dark:border-slate-700 dark:bg-transparent_blue  text-violet-600 rounded-lg transition duration-300 focus:outline-none"
      >
      Go to home
  </button>
  </div>

</template>