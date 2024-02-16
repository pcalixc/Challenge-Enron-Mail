<!-- eslint-disable vue/multi-word-component-names -->
<script setup lang="ts">
import { defineProps, ref } from 'vue'
import {useDark, useToggle} from '@vueuse/core'
const isDark= useDark();
const togleDark= useToggle(isDark)

const props = defineProps(['getData'])

const searchTerm = ref('')

</script>

<template>
<div
    class="container mx-auto dark:bg-primary_dark bg-white rounded-lg p-4 dark:shadow-none shadow-violet-100 shadow-md"
  >
    <div
      class="flex items-center relative justify-between max-w-screen-xl px-4 mx-auto md:px-6 lg:px-8"
    >
      <div class="flex flex-row items-center justify-center mb-5 mr-3 lg:mb-0">
      
        <div class="flex items-center justify-center">
          <img
            src="../../public/envelope.png"
            class="w-12 h-13 transform hover:scale-125 transition duration-700"
            alt=""
          />
        </div>

      
        <div class="flex flex-col ml-3">
         
          <div class="mb-0 flex flex-col">
            <a
              href="#"
              class="text-lg dark:text-slate-100 font-semibold tracking-wider uppercase text-black focus:outline-none focus:shadow-outline"
              >Mail Searcher</a
            >
            <span class="text-slate-600 dark:text-slate-300 text-[0.70rem] font-xs"
              >Find Emails Effortlessly!</span
            >
          </div>
        </div>
      </div>
      <div class="flex items-center justify-center max-w-md mr-8">
        <form
          @submit.prevent="props.getData(1, searchTerm)"
          class="relative w-full text-slate-800 dark:text-slate-50"
        >
          <input
            type="search"
            placeholder="Enter a keyword..."
            v-model="searchTerm"
      
            class="peer cursor-pointer relative shadow-inner z-10 h-12 rounded-xl border-2 border-[#2525ff52] dark:bg-[#28287549] bg-transparent pl-16 pr-4 outline-none focus:w-full focus:cursor-text focus:border-[#4a4aff5c] focus:border-3"
          />
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="absolute inset-y-0 my-auto h-8 w-12 border-r border-transparent stroke-gray-500 px-3.5 dark:peer-focus:border-slate-600 peer-focus:border-slate-300 dark:peer-focus:stroke-slate-200 peer-focus:stroke-slate-500"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
            stroke-width="2"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
            />
          </svg>
        </form>
      </div>

      <div class="flex justify-center items-center">
        <span class="">
          <svg class="h-6 w-6 dark:text-gray-500 text-[#9292FF]" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z"
            />
          </svg>
        </span>
        <button
        title="Toggle Theme"
        @click="togleDark()"
        class="w-14 h-6 flex items-center  mx-3 px-1 rounded-full p-1 bg-gray-200 dark:bg-[#2525ff52]   relative transition-colors duration-500 ease-in focus:outline-none focus:ring-2 focus:ring-[#9292FF] dark:focus:ring-[#2525ff52] focus:border-transparent"
      >
        <div
          id="toggle"
          class="rounded-full w-5 h-5 bg-[#9292FF] dark:bg-gray-200 relative ml-0 dark:ml-6 pointer-events-none transition-all duration-300 ease-out"
        ></div>
      </button>

        <span class="">
          <svg class="h-6 w-6 text-gray-400 dark:text-[#2525ff52]" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z"
            />
          </svg>
        </span>
      </div>
      
    </div>
  </div>
</template>
