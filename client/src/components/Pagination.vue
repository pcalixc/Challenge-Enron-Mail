<!-- eslint-disable vue/multi-word-component-names -->
<script setup lang="ts">
import { useEmailsStore } from '@/stores/emails';
const emailsStore = useEmailsStore()

</script>

<template>
  <div
    id="pagination"
    v-if="emailsStore.totalResults > 0"
    class="flex items-center justify-center gap-2 mt-0 mx-24"
  >
    <button
      @click="emailsStore.getData(emailsStore.currentPage - 1, emailsStore.currentSearchTerm)"
      :disabled="emailsStore.currentPage <= 1"
      :class="[
        'flex items-center px-5 py-2 text-sm capitalize transition-colors duration-200   rounded-xl gap-x-2',
        emailsStore.currentPage == 1
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
    <div
      v-if="emailsStore.currentPage > 0 && emailsStore.currentPage < 3"
      class="items-center hidden md:flex gap-x-3"
    >
      <button
        v-for="pageNumber in Math.min(4, emailsStore.totalPages)"
        :key="pageNumber"
        @click="emailsStore.getData(pageNumber, emailsStore.currentSearchTerm)"
        :class="[
          'px-2 py-1 text-sm rounded-xl ',
          emailsStore.currentPage === pageNumber
            ? 'text-royal_purple font-bold  bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 '
            : ' hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue   dark:text-slate-200  dark:hover:bg-gray-900'
        ]"
      >
        {{ pageNumber }}
      </button>

      <button
        v-if="emailsStore.totalPages > 5"
        class="px-2 py-1 text-sm text-gray-500 rounded-xl hover:bg-gray-100 dark:text-slate-100"
        disabled
      >
        ...
      </button>

      <button
        @click="emailsStore.getData(emailsStore.totalPages, emailsStore.currentSearchTerm)"
        :class="[
          'px-2 py-1 text-sm rounded-xl ',
          emailsStore.currentPage === emailsStore.totalPages 
            ? 'text-royal_purple font-bold bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900  '
            : ' hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue   dark:text-slate-200  dark:hover:bg-gray-900'
        ]"
      >
        {{ emailsStore.totalPages }}
      </button>
    </div>

    <div v-if="emailsStore.currentPage > 2 && emailsStore.currentPage != emailsStore.totalPages ">
      <section v-if="emailsStore.totalPages <= 5" class="items-center hidden md:flex gap-x-3">
        <button
          v-for="pageNumber in emailsStore.totalPages"
          :key="pageNumber"
          @click="emailsStore.getData(pageNumber, emailsStore.currentSearchTerm)"
          class="px-2 py-1 text-sm rounded-xl text-royal_purple font-bold bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900"
          :class="{
            'text-royal_purple font-bold bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 ':
              emailsStore.currentPage === pageNumber
          }"
        >
          {{ pageNumber }}
        </button>
      </section>

      <section v-else class="items-center hidden md:flex gap-x-3">
        <button
          @click="emailsStore.getData(1, emailsStore.currentSearchTerm)"
          class="px-2 py-1 text-sm rounded-xl hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue dark:text-slate-200 dark:hover:bg-gray-900"
          :class="{
            'text-royal_purple font-bold  bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 ':
              emailsStore.currentPage === 1
          }"
        >
          1
        </button>
        <span v-if="emailsStore.currentPage > 3" class="dark:text-slate-100">...</span>

        <template v-if="emailsStore.currentPage === 1">
          <button
            @click="emailsStore.getData(2, emailsStore.currentSearchTerm)"
            class="px-2 py-1 text-sm rounded-xl hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue dark:text-slate-200 dark:hover:bg-gray-900"
          >
            2
          </button>
        </template>

        <template v-else>
          <button
            @click="emailsStore.getData(emailsStore.currentPage - 2, emailsStore.currentSearchTerm)"
            class="px-2 py-1 text-sm rounded-xl hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue dark:text-slate-200 dark:hover:bg-gray-900"
            :class="{
              'text-royal_purple font-bold bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 ':
                emailsStore.currentPage === emailsStore.currentPage - 2
            }"
          >
            {{ emailsStore.currentPage - 2 }}
          </button>
          <button
            @click="emailsStore.getData(emailsStore.currentPage - 1, emailsStore.currentSearchTerm)"
            class="px-2 py-1 text-sm rounded-xl hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue dark:text-slate-200 dark:hover:bg-gray-900"
            :class="{
              'text-royal_purple font-bold bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 ':
                emailsStore.currentPage === emailsStore.currentPage - 1
            }"
          >
            {{ emailsStore.currentPage - 1 }}
          </button>
        </template>
        <button
          @click="emailsStore.getData(emailsStore.currentPage, emailsStore.currentSearchTerm)"
          class="px-2 py-1 text-sm rounded-xl text-royal_purple font-bold bg-blue-100/60 hover:bg-gray-300 bg-slate-200 shadow-xs dark:bg-ligth_blue dark:text-slate-200 dark:hover:bg-gray-900"
          :class="{
            'text-royal_purple font-bold bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 ':
              emailsStore.currentPage === emailsStore.currentPage
          }"
        >
          {{ emailsStore.currentPage }}
        </button>
        <button
          @click="emailsStore.getData(emailsStore.currentPage + 1, emailsStore.currentSearchTerm)"
          class="px-2 py-1 text-sm rounded-xl hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue dark:text-slate-200 dark:hover:bg-gray-900"
          :class="{
            'text-royal_purple font-bold bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 ':
              emailsStore.currentPage === emailsStore.currentPage + 1
          }"
        >
          {{ emailsStore.currentPage + 1 }}
        </button>
        <button
          @click="emailsStore.getData(emailsStore.currentPage + 2, emailsStore.currentSearchTerm)"
          class="px-2 py-1 text-sm rounded-xl hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue dark:text-slate-200 dark:hover:bg-gray-900"
          :class="{
            'text-royal_purple font-bold bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 ':
              emailsStore.currentPage === emailsStore.currentPage + 2
          }"
        >
          {{ emailsStore.currentPage + 2 }}
        </button>
        <span v-if="emailsStore.currentPage + 2 < emailsStore.totalPages" class="dark:text-slate-100">...</span>
        <button
          @click="emailsStore.getData(emailsStore.totalPages , emailsStore.currentSearchTerm)"
          class="px-2 py-1 text-sm rounded-xl hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue dark:text-slate-200 dark:hover:bg-gray-900"
          :class="{
            'text-royal_purple font-bold bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 ':
              emailsStore.currentPage === emailsStore.totalPages
          }"
        >
          {{ emailsStore.totalPages  }}
        </button>
      </section>
    </div>

    <div v-if="emailsStore.currentPage == emailsStore.totalPages" class="items-center hidden md:flex gap-x-3">
      <button
        v-for="pageNumber in Math.min(3, emailsStore.totalPages)"
        :key="pageNumber"
        @click="emailsStore.getData(pageNumber, emailsStore.currentSearchTerm)"
        :class="[
          'px-2 py-1 text-sm rounded-xl ',
          emailsStore.currentPage === pageNumber
            ? 'text-royal_purple font-bold  bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 '
            : ' hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue   dark:text-slate-200  dark:hover:bg-gray-900'
        ]"
      >
        {{ pageNumber }}
      </button>

      <button
        v-if="emailsStore.totalPages > 5"
        class="px-2 py-1 text-sm text-gray-500 rounded-xl hover:bg-gray-100 dark:text-slate-100"
        disabled
      >
        ...
      </button>
      <button
        @click="emailsStore.getData(Math.round(emailsStore.totalPages / 2) , emailsStore.currentSearchTerm)"
        :class="[
          'px-2 py-1 text-sm rounded-xl ',
          emailsStore.currentPage === Math.round(emailsStore.totalPages / 2) 
            ? 'text-royal_purple font-bold  bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 '
            : ' hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue   dark:text-slate-200  dark:hover:bg-gray-900'
        ]"
      >
        {{ Math.round(emailsStore.totalPages / 2)  }}
      </button>
      <button
        @click="emailsStore.getData(Math.round(emailsStore.totalPages / 2), emailsStore.currentSearchTerm)"
        :class="[
          'px-2 py-1 text-sm rounded-xl ',
          emailsStore.currentPage === Math.round(emailsStore.totalPages / 2)
            ? 'text-royal_purple font-bold  bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 '
            : ' hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue   dark:text-slate-200  dark:hover:bg-gray-900'
        ]"
      >
        {{ Math.round(emailsStore.totalPages / 2) }}
      </button>
      <button
        @click="emailsStore.getData(Math.round(emailsStore.totalPages / 2) + 1, emailsStore.currentSearchTerm)"
        :class="[
          'px-2 py-1 text-sm rounded-xl ',
          emailsStore.currentPage === Math.round(emailsStore.totalPages / 2 + 1)
            ? 'text-royal_purple font-bold  bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 '
            : ' hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue   dark:text-slate-200  dark:hover:bg-gray-900'
        ]"
      >
        {{ Math.round(emailsStore.totalPages / 2) + 1 }}
      </button>
      <button
        v-if="emailsStore.totalPages > 5"
        class="px-2 py-1 text-sm text-gray-500 rounded-xl hover:bg-gray-100 dark:text-slate-100"
        disabled
      >
        ...
      </button>
      <button
  
        :class="[
          'px-2 py-1 text-sm rounded-xl ',
          emailsStore.currentPage === emailsStore.totalPages -1
            ? 'text-royal_purple font-bold  bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 '
            : ' hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue   dark:text-slate-200  dark:hover:bg-gray-900'
        ]"
      >
        {{ emailsStore.totalPages-1 }}
      </button>

      <button
        @click="emailsStore.getData(emailsStore.totalPages , emailsStore.currentSearchTerm)"
        :class="[
          'px-2 py-1 text-sm rounded-xl ',
          emailsStore.currentPage === emailsStore.totalPages 
            ? 'text-royal_purple font-bold  bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 '
            : ' hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue   dark:text-slate-200  dark:hover:bg-gray-900'
        ]"
      >
        {{ emailsStore.totalPages }}
      </button>
    </div>

    <button
      @click="emailsStore.getData(emailsStore.currentPage + 1, emailsStore.currentSearchTerm)"
      :disabled="emailsStore.currentPage == emailsStore.totalPages "
      :class="[
        'flex items-center px-8 py-2 text-sm capitalize transition-colors duration-200  rounded-xl gap-x-2 ',
        emailsStore.currentPage == emailsStore.totalPages 
          ? ' text-gray-400  bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900'
          : 'hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue   dark:text-slate-200 dark:hover:bg-gray-900'
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
