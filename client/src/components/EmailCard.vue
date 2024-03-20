<script setup lang="ts">
import { useEmailsStore } from '@/stores/emails'
const emailsStore = useEmailsStore()
import { ConvertDateFormat, HighlighWord } from '../utils/emails.utilities'
</script>

<template>
  <div
    v-for="(email, index) in emailsStore.hits"
    :key="index"
    class="relative h-[14rem]  group overflow-hidden p-4 rounded-xl bg-white border border-gray-200 dark:border-gray-800 dark:bg-gray-900 dark:border-white/15 before:rounded-[7px] before:absolute dark:before:border-white/20 before:bg-gradient-to-b dark:before:from-white/10 dark:before:to-transparent before:shadow dark:before:shadow-gray-950"
  >
    <div
      aria-hidden="true"
      class="inset-0 absolute aspect-video border rounded-full -translate-y-1/2 group-hover:-translate-y-1/4 duration-300 bg-gradient-to-b from-royal_purple to-white dark:from-white dark:to-white blur-2xl opacity-35 dark:opacity-5 dark:group-hover:opacity-10"
    ></div>
    <div class="relative">
      <div class="text-xs text-right ml-auto dark:text-gray-50">
        {{ ConvertDateFormat(email._source.date) }}
      </div>
      <div class="flex items-center justify-between">
        <div class="flex items-center">
          <div class="flex flex-col ml-1">
            <div class="flex gap-1">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="mt-2 dark:fill-white fill-[#414141]"
                id="Filled"
                viewBox="0 0 24 24"
                width="11"
                height="12"
              >
                <path
                  d="M23.954,5.542,15.536,13.96a5.007,5.007,0,0,1-7.072,0L.046,5.542C.032,5.7,0,5.843,0,6V18a5.006,5.006,0,0,0,5,5H19a5.006,5.006,0,0,0,5-5V6C24,5.843,23.968,5.7,23.954,5.542Z"
                />
                <path
                  d="M14.122,12.546l9.134-9.135A4.986,4.986,0,0,0,19,1H5A4.986,4.986,0,0,0,.744,3.411l9.134,9.135A3.007,3.007,0,0,0,14.122,12.546Z"
                />
              </svg>
              <h3
                :title="email._source.subject"
                v-html="HighlighWord(email._source.subject, emailsStore.currentSearchTerm)"
                class="text-base font-semibold leading-7 tracking-tight h-[2rem] dark:text-white text-gray-900 truncate text-ellipsis w-[15rem]"
              ></h3>
            </div>
            <div class="flex">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                id="Layer_1"
                class="mr-1 mt-1 fill-[#414141] dark:fill-slate-400"
                data-name="Layer 1"
                viewBox="0 0 24 24"
                width="12"
                height="12"
              >
                <path
                  d="M16.043,14H7.957A4.963,4.963,0,0,0,3,18.957V24H21V18.957A4.963,4.963,0,0,0,16.043,14Z"
                />
                <circle cx="12" cy="6" r="6" />
              </svg>
              <h3
                :title="email._source.from"
                v-html="HighlighWord(email._source.from, emailsStore.currentSearchTerm)"
                class="text-sm font-semibold leading-6 line-clamp-1 text-gray-700 dark:text-gray-300  text-ellipsis w-[14rem]"
              ></h3>
            </div>
          </div>
        </div>
      </div>

      <div class="mt-3 pb-2 mx-1 h-[4.5rem] rounded-b-[--card-border-radius]">
        <p
          class="text-gray-700 font-xs text-sm dark:text-gray-300 line-clamp-3"
          :title="email._source.content"
          v-html="HighlighWord(email._source.content, emailsStore.currentSearchTerm)"
        ></p>
      </div>

      <div class="flex gap-3 -mb-8 py-2 border-t border-gray-200 dark:border-gray-800">
        <button
          data-ripple-light="true"
          class="mx-2 center rounded-lg bg-royal_purple text-white py-2 px-4 font-sans text-xs font-bold duration-300 hover:scale-105 uppercase active:opacity-[0.85] active:shadow-none disabled:pointer-events-none disabled:opacity-50 disabled:shadow-none dark:text-slate-100 dark:bg-royal_blue"
          @click="emailsStore.asigneSelectedContent(index)"
        >
          Open Email
        </button>
      </div>
    </div>
  </div>
</template>
