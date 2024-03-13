<!-- eslint-disable vue/multi-word-component-names -->
<script setup lang="ts">
import { useEmailsStore } from '@/stores/emails'
const emailsStore = useEmailsStore()

const ascendingPageNumbers = [1, 2, 3, 4, 5, 6]
const ascendingPageNumbersReverse = [6, 5, 4, 3, 2, 1]
const descendingPageNumbers = [3, 2, 1]
const descendingPageNumbersReverse = [1, 2, 3]
</script>

<template>
  <div
    id="pagination"
    v-if="
      emailsStore.totalResults > 0 &&
      !emailsStore.isLoading &&
      !emailsStore.ServerErrorResponse.errorStatus
    "
    class="flex items-center justify-center gap-2 mt-0 mx-24 my-22"
  >
    <button
      @click="emailsStore.getData(emailsStore.currentPage - 1, emailsStore.currentSearchTerm)"
      :disabled="emailsStore.currentPage <= 1"
      :class="[
        'flex items-center px-5 py-2 font-semibold text-xs duration-300 hover:scale-105 uppercase rounded-xl gap-x-2',
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

    <section class="hideInSmallScreens">
      <div v-if="emailsStore.totalPages < 10 && emailsStore.currentPage > 0">
        <button
          v-for="pageNumber in emailsStore.totalPages"
          :key="pageNumber"
          @click="emailsStore.getData(pageNumber, emailsStore.currentSearchTerm)"
          :class="[
            'px-2 py-1 text-sm rounded-xl mx-1',
            emailsStore.currentPage === pageNumber
              ? 'text-royal_purple font-bold  bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 '
              : ' hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue   dark:text-slate-200  dark:hover:bg-gray-900'
          ]"
        >
          {{ pageNumber }}
        </button>
      </div>

      <div
        v-if="
          emailsStore.totalPages >= 10 && emailsStore.currentPage > 0 && emailsStore.currentPage < 6
        "
      >
        <button
          v-for="pageNumber in ascendingPageNumbers"
          :key="pageNumber"
          @click="emailsStore.getData(pageNumber, emailsStore.currentSearchTerm)"
          :class="[
            'px-2 py-1 text-sm rounded-xl mx-1',
            emailsStore.currentPage === pageNumber
              ? 'text-royal_purple font-bold  bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 '
              : ' hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue   dark:text-slate-200  dark:hover:bg-gray-900'
          ]"
        >
          {{ pageNumber }}
        </button>

        <button
          v-if="emailsStore.totalPages > 5"
          class="px-2 py-1 text-sm mx-1 text-gray-500 rounded-xl hover:bg-gray-100 dark:text-slate-100"
          disabled
        >
          ...
        </button>

        <button
          v-for="pageNumber in descendingPageNumbers"
          :key="pageNumber"
          @click="
            emailsStore.getData(
              emailsStore.totalPages - pageNumber + 1,
              emailsStore.currentSearchTerm
            )
          "
          :class="[
            'px-2 py-1 text-sm rounded-xl mx-1',
            emailsStore.currentPage === emailsStore.totalPages - pageNumber + 1
              ? 'text-royal_purple font-bold  bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 '
              : ' hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue   dark:text-slate-200  dark:hover:bg-gray-900'
          ]"
        >
          {{ emailsStore.totalPages - pageNumber + 1 }}
        </button>
      </div>

      <div
        v-if="emailsStore.currentPage >= 6 && emailsStore.currentPage <= emailsStore.totalPages - 5"
      >
        <button
          v-for="pageNumber in descendingPageNumbersReverse"
          :key="pageNumber"
          @click="emailsStore.getData(pageNumber, emailsStore.currentSearchTerm)"
          :class="[
            'px-2 py-1 text-sm rounded-xl mx-1',
            emailsStore.currentPage === pageNumber
              ? 'text-royal_purple font-bold  bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 '
              : ' hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue   dark:text-slate-200  dark:hover:bg-gray-900'
          ]"
        >
          {{ pageNumber }}
        </button>

        <button
          v-if="emailsStore.totalPages > 5"
          class="px-2 py-1 text-sm mx-1 text-gray-500 rounded-xl hover:bg-gray-100 dark:text-slate-100"
          disabled
        >
          ...
        </button>

        <button
          @click="emailsStore.getData(emailsStore.currentPage - 2, emailsStore.currentSearchTerm)"
          class="px-2 py-1 mx-1 text-sm rounded-xl hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue dark:text-slate-200 dark:hover:bg-gray-900"
          :class="{
            'text-royal_purple font-bold bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 ':
              emailsStore.currentPage === emailsStore.currentPage - 2
          }"
        >
          {{ emailsStore.currentPage - 2 }}
        </button>
        <button
          @click="emailsStore.getData(emailsStore.currentPage - 1, emailsStore.currentSearchTerm)"
          class="px-2 py-1 text-sm mx-1 rounded-xl hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue dark:text-slate-200 dark:hover:bg-gray-900"
          :class="{
            'text-royal_purple font-bold bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 ':
              emailsStore.currentPage === emailsStore.currentPage - 1
          }"
        >
          {{ emailsStore.currentPage - 1 }}
        </button>
        <button
          @click="emailsStore.getData(emailsStore.currentPage, emailsStore.currentSearchTerm)"
          class="px-2 py-1 text-sm mx-1 rounded-xl text-royal_purple font-bold bg-blue-100/60 hover:bg-gray-300 bg-slate-200 shadow-xs dark:bg-ligth_blue dark:text-slate-200 dark:hover:bg-gray-900"
          :class="{
            'text-royal_purple font-bold bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 ':
              emailsStore.currentPage === emailsStore.currentPage
          }"
        >
          {{ emailsStore.currentPage }}
        </button>
        <button
          @click="emailsStore.getData(emailsStore.currentPage + 1, emailsStore.currentSearchTerm)"
          class="px-2 py-1 text-sm mx-1 rounded-xl hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue dark:text-slate-200 dark:hover:bg-gray-900"
          :class="{
            'text-royal_purple font-bold bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 ':
              emailsStore.currentPage === emailsStore.currentPage + 1
          }"
        >
          {{ emailsStore.currentPage + 1 }}
        </button>
        <button
          @click="emailsStore.getData(emailsStore.currentPage + 2, emailsStore.currentSearchTerm)"
          class="px-2 py-1 text-sm mx-1 rounded-xl hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue dark:text-slate-200 dark:hover:bg-gray-900"
          :class="{
            'text-royal_purple font-bold bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 ':
              emailsStore.currentPage === emailsStore.currentPage + 2
          }"
        >
          {{ emailsStore.currentPage + 2 }}
        </button>
        <button
          v-if="emailsStore.totalPages > 5"
          class="px-2 py-1 text-sm mx-1 text-gray-500 rounded-xl hover:bg-gray-100 dark:text-slate-100"
          disabled
        >
          ...
        </button>

        <button
          v-for="pageNumber in descendingPageNumbers"
          :key="pageNumber"
          @click="
            emailsStore.getData(
              emailsStore.totalPages - pageNumber + 1,
              emailsStore.currentSearchTerm
            )
          "
          :class="[
            'px-2 py-1 text-sm rounded-xl mx-1',
            emailsStore.currentPage === emailsStore.totalPages - pageNumber + 1
              ? 'text-royal_purple font-bold  bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 '
              : ' hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue   dark:text-slate-200  dark:hover:bg-gray-900'
          ]"
        >
          {{ emailsStore.totalPages - pageNumber + 1 }}
        </button>
      </div>

      <div
        v-if="
          emailsStore.totalPages >= 10 &&
          emailsStore.currentPage > emailsStore.totalPages - 5 &&
          emailsStore.currentPage <= emailsStore.totalPages
        "
      >
        <button
          v-for="pageNumber in descendingPageNumbersReverse"
          :key="pageNumber"
          @click="emailsStore.getData(pageNumber, emailsStore.currentSearchTerm)"
          :class="[
            'px-2 py-1 text-sm rounded-xl mx-1 ',
            emailsStore.currentPage === pageNumber
              ? 'text-royal_purple font-bold  bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 '
              : ' hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue   dark:text-slate-200  dark:hover:bg-gray-900'
          ]"
        >
          {{ pageNumber }}
        </button>

        <button
          v-if="emailsStore.totalPages > 5"
          class="px-2 py-1 text-sm mx-1 text-gray-500 rounded-xl hover:bg-gray-100 dark:text-slate-100"
          disabled
        >
          ...
        </button>

        <button
          v-for="pageNumber in ascendingPageNumbersReverse"
          :key="pageNumber"
          @click="
            emailsStore.getData(
              emailsStore.totalPages - pageNumber + 1,
              emailsStore.currentSearchTerm
            )
          "
          :class="[
            'px-2 py-1 text-sm rounded-xl mx-1',
            emailsStore.currentPage === emailsStore.totalPages - pageNumber + 1
              ? 'text-royal_purple font-bold  bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 '
              : ' hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue   dark:text-slate-200  dark:hover:bg-gray-900'
          ]"
        >
          {{ emailsStore.totalPages - pageNumber + 1 }}
        </button>
      </div>
    </section>

    <button
      @click="emailsStore.getData(emailsStore.currentPage + 1, emailsStore.currentSearchTerm)"
      :disabled="emailsStore.currentPage == emailsStore.totalPages"
      :class="[
        'flex items-center px-5 py-2 font-semibold text-xs duration-300 hover:scale-105 uppercase rounded-xl gap-x-2',
        emailsStore.currentPage == emailsStore.totalPages
          ? ' text-gray-400  bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900'
          : 'hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue   dark:text-slate-200 dark:hover:bg-[#3A3AFF]'
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

<style scoped>
@media (max-width: 1024px) {
  .hideInSmallScreens {
    display: none;
  }
}
</style>
