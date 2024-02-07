<script setup lang="ts">
import { RouterLink } from 'vue-router';

</script>

<template>
 
  <div class="antialiased bg-gray-900 shadow-gray-800 shadow-lg">
  <div class="w-full  text-gray-200 bg-gray-900">
    <div class="flex flex-col max-w-screen-xl px-4 mx-auto md:items-center md:justify-between md:flex-row md:px-6 lg:px-8">
      <div class="flex flex-row items-center justify-between p-4">
        <a href="#" class="text-lg font-semibold tracking-widest  uppercase rounded-lg text-white focus:outline-none focus:shadow-outline">Final Challenge</a>
   
      </div>

      <RouterLink  to="/login" class="px-4 py-2 mt-2 text-sm bg-blue-800 font-semibold bg-transparent rounded-lg md:mt-0 md:ml-4 hover:bg-blue-900 focus:bg-blue-900 focus:outline-none focus:shadow-outline">
            Sign In
      </RouterLink>
    </div>
  </div>
</div>

</template>