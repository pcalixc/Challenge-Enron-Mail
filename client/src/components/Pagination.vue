<!-- eslint-disable vue/multi-word-component-names -->
<script setup lang="ts">

import {defineProps } from 'vue'

interface Props {
currentPage: number
totalResults: number
currentSearchTerm: string
totalPages: number
getData: Function
}
const props = defineProps<Props>()

</script>


<template>
        <div
        id="pagination"
        v-if="props.totalResults > 0"
        class="flex items-center justify-center gap-2 mt-3 mx-24"
      >
        <button
          @click="getData(props.currentPage - 1, props.currentSearchTerm)"
          :disabled="props.currentPage <= 1"
          :class="[
            'flex items-center px-5 py-2 text-sm capitalize transition-colors duration-200   rounded-xl gap-x-2',
            currentPage == 1
              ? ' bg-gray-50 text-gray-300 dark:bg-deep_blue dark:text-slate-500'
              : 'hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue   dark:text-slate-200  dark:hover:bg-gray-900'
          ]"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke-width="1.5"
            stroke="currentColor"
            class="w-5 h-5 rtl:-scale-x-100"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              d="M6.75 15.75L3 12m0 0l3.75-3.75M3 12h18"
            />
          </svg>
          <span> Previous </span>
        </button>
        <div v-if="props.currentPage > 0 && props.currentPage < 3" class="items-center hidden md:flex gap-x-3">
          <button
            v-for="pageNumber in Math.min(3, props.totalPages)"
            :key="pageNumber"
            @click="props.getData(pageNumber)"
            :class="[
              'px-2 py-1 text-sm rounded-xl ',
              props.currentPage === pageNumber
                ? 'text-blue-500  bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 '
                : ' hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue   dark:text-slate-200  dark:hover:bg-gray-900'
            ]"
          >
            {{ pageNumber }}
          </button>

          <button
            v-if="props.totalPages > 5"
            class="px-2 py-1 text-sm text-gray-500 rounded-xl hover:bg-gray-100 dark:text-slate-100"
            disabled
          >
            ...
          </button>

          <button
            @click="props.getData(props.totalPages)"
            :class="[
              'px-2 py-1 text-sm rounded-xl ',
              props.currentPage === props.totalPages - 1
                ? 'text-blue-500 bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900  '
                : ' hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue   dark:text-slate-200  dark:hover:bg-gray-900'
            ]"
          >
            {{ props.totalPages }}
          </button>
        </div>
     
        <div v-if="props.currentPage > 2">
          <section v-if="props.totalPages <= 5" class="items-center hidden md:flex gap-x-3">
   
            <button
              v-for="pageNumber in props.totalPages"
              :key="pageNumber"
              @click="getData(pageNumber)"
              class="px-2 py-1 text-sm rounded-xl text-blue-500 bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900"
              :class="{
                'text-blue-500 bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 ':
                props.currentPage === pageNumber
              }"
            >
              {{ pageNumber }}
            </button>
          </section>
          <section v-else class="items-center hidden md:flex gap-x-3">
        
            <button
              @click="props.getData(1)"
              class="px-2 py-1 text-sm rounded-xl hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue dark:text-slate-200 dark:hover:bg-gray-900"
              :class="{
                'text-blue-500  bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 ':
                props.currentPage === 1
              }"
            >
              1
            </button>
            <span v-if="props.currentPage > 3" class="dark:text-slate-100">...</span>
            <template v-if="props.currentPage === 1">
              <button
                @click="props.getData(2)"
                class="px-2 py-1 text-sm rounded-xl hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue dark:text-slate-200 dark:hover:bg-gray-900"
               
              >
                2
              </button>
            </template>
            <template v-else>
              <button
                @click="props.getData(props.currentPage - 1)"
                class="px-2 py-1 text-sm rounded-xl hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue dark:text-slate-200 dark:hover:bg-gray-900"
                :class="{
                  'text-blue-500 bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 ':
                  props.currentPage === props.currentPage - 1
                }"
              >
                {{ props.currentPage - 1 }}
              </button>
            </template>
            <button
              @click="props.getData(props.currentPage)"
              class="px-2 py-1 text-sm rounded-xl text-blue-500 bg-blue-100/60 hover:bg-gray-300 bg-slate-200 shadow-xs dark:bg-ligth_blue dark:text-slate-200 dark:hover:bg-gray-900"
              :class="{
                'text-blue-500 bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 ':
                props.currentPage === props.currentPage
              }"
            >
              {{ props.currentPage }}
            </button>
            <button
              @click="props.getData(props.currentPage + 1)"
              class="px-2 py-1 text-sm rounded-xl hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue dark:text-slate-200 dark:hover:bg-gray-900"
              :class="{
                'text-blue-500 bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 ':
                props.currentPage === props.currentPage + 1
              }"
            >
              {{ props.currentPage + 1 }}
            </button>
            <span v-if="props.currentPage + 2 < props.totalPages" class="dark:text-slate-100">...</span>
            <button
              @click="props.getData(props.totalPages)"
              class="px-2 py-1 text-sm rounded-xl hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue dark:text-slate-200 dark:hover:bg-gray-900"
              :class="{
                'text-blue-500 bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 ':
                props.currentPage === props.totalPages
              }"
            >
              {{ props.totalPages }}
            </button>
          </section>
        </div>
   
        <button
          @click="props.getData(props.currentPage + 1, props.currentSearchTerm)"
          :disabled="props.currentPage == props.totalPages"
          :class="[
            'flex items-center px-5 py-2 text-sm capitalize transition-colors duration-200  rounded-xl gap-x-2 ',
            props.currentPage == props.totalPages
              ? ' text-blue-500 bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900'
              : 'hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue   dark:text-slate-200   dark:hover:bg-gray-900'
          ]"
        >
          <span> Next </span>
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke-width="1.5"
            stroke="currentColor"
            class="w-5 h-5 rtl:-scale-x-100"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              d="M17.25 8.25L21 12m0 0l-3.75 3.75M21 12H3"
            />
          </svg>
        </button>
      </div>

</template>