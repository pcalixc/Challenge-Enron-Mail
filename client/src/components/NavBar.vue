<script setup lang="ts">
import { defineProps, defineEmits,ref } from 'vue';

const props = defineProps(['getData']);

const searchTerm= ref('')

interface Emits{
    (event: 'Search', value: String): void
}

 const emit = defineEmits<Emits>()

 function Search(value: string) {
  console.log("aaaaaaaaa", value)
  searchTerm.value = value;
  emit('Search', value);
}


</script>

<template>
  <div class="container mx-auto  bg-white rounded-lg p-4 shadow-violet-100 shadow-md">
  <div class="flex items-center relative justify-between max-w-screen-xl px-4 mx-auto md:px-6 lg:px-8">


    <div class="flex flex-row items-center justify-center mb-5 mr-3 lg:mb-0">
  <!-- Primera columna -->
  <div class="flex items-center justify-center">
    <img src="../../public/envelope.png" class="w-12 h-13  transform hover:scale-125 transition duration-700 " alt="" />
  </div>

  <!-- Segunda columna -->
  <div class="flex flex-col ml-3">
    <!-- Primera fila -->
    <div class="mb-0 flex flex-col">
      <a href="#" class="text-lg font-semibold tracking-wider uppercase text-black focus:outline-none focus:shadow-outline">Mail Searcher</a>
      <span class="text-slate-600 text-[0.70rem] font-xs">Find Emails Effortlessly!</span>
    </div>
  </div>
</div>
    <div class="flex items-center justify-center max-w-md mr-8">
      <form action="" class="relative w-full">
        <input type="search" 
        placeholder="Enter a keyword..."
        v-model="searchTerm"
        @keyup.enter="Search(searchTerm)"
        @input="Search(searchTerm)"
              class="peer cursor-pointer relative shadow-inner z-10 h-12 rounded-xl  border-2 border-[#4a4aff2e] bg-transparent pl-16 pr-4 outline-none focus:w-full focus:cursor-text focus:border-[#4a4aff5c]  focus:border-3" />
        <svg xmlns="http://www.w3.org/2000/svg" class="absolute inset-y-0 my-auto h-8 w-12 border-r border-transparent stroke-gray-500 px-3.5  peer-focus:border-slate-300 peer-focus:stroke-slate-500" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
          <path stroke-linecap="round" stroke-linejoin="round" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
        </svg>
      </form>
    </div>

    <div class="flex justify-center items-center">
    <span class="">
        <svg class="h-6 w-6 text-gray-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z" />
        </svg>
    </span>
    <!-- Switch Container -->
    <div :class="{ 'bg-cyan-700': toggleActive}" class="w-14 h-7 flex items-center bg-gray-300 rounded-full mx-3 px-1" @click="handleToggleActive">
        <!-- Switch -->
        <div class="bg-white w-5 h-5 rounded-full shadow-md transform" :class="{ 'translate-x-7': toggleActive}"></div>
    </div>
    <span class="">
        <svg class="h-6 w-6 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z" />
        </svg>
    </span>
</div>

    
  </div> 
</div>

</template>
