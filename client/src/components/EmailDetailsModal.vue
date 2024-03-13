<script setup lang="ts">
import { useEmailsStore } from '@/stores/emails'
const emailsStore = useEmailsStore()
import { HighlighWord, SeparateEmailsByCommas, ConvertDateFormat } from '@/utils/emails.utilities'

function closeModal() {
  emailsStore.modalOpen = false
}
</script>

<template>
  <Transition>
    <div v-if="emailsStore.modalOpen" class="fixed fade-in inset-0 z-50 overflow-hidden">
      <div
        class="absolute inset-0 bg-gray-900 bg-opacity-50 transition-opacity"
        @click="closeModal()"
      ></div>
      <!-- Sidebar Content -->
      <section class="absolute inset-y-0 right-0 pl-10 flex">
        <div class="rounded-l-2xl">
          <div
            class="h-full flex flex-col lg:w-[33rem] sm:w-[80vw] md:w-[70vw] modal py-6 bg-white border dark:border-slate-800 dark:shadow-xl dark:bg-gray-900 shadow-xl rounded-2xl"
          >
            <!-- Sidebar Header -->
            <div class="flex items-center justify-between px-4">
              <div class="px-2 flex items-end mr-auto space-x-4">
                <div class="flex items-center space-x-2">
                  <button
                    @click="emailsStore.asigneSelectedContent(emailsStore.selectedEmailIndex - 1)"
                    :class="[
                      'bg-gray-200  p-1.5 rounded-lg',
                      emailsStore.selectedEmailIndex != 0
                        ? 'text-gray-100 bg-royal_purple dark:text-slate-100 dark:bg-royal_blue hover:scale-105'
                        : 'text-gray-400 dark:bg-gray-600  disabled pointer-events-none'
                    ]"
                    title="Previous Email"
                  >
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      class="h-7 w-7"
                      viewBox="0 0 20 20"
                      fill="currentColor"
                    >
                      <path
                        fill-rule="evenodd"
                        d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z"
                        clip-rule="evenodd"
                      ></path>
                    </svg>
                  </button>
                  <button
                    @click="emailsStore.asigneSelectedContent(emailsStore.selectedEmailIndex + 1)"
                    :class="[
                      'bg-gray-200  p-1.5 rounded-lg',
                      emailsStore.selectedEmailIndex != 6
                        ? 'text-gray-100 bg-royal_purple dark:text-slate-100 dark:bg-royal_blue hover:scale-105'
                        : 'text-gray-400 dark:bg-gray-700  disabled pointer-events-none'
                    ]"
                    title="Nex Email"
                  >
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      class="h-7 w-7"
                      viewBox="0 0 20 20"
                      fill="currentColor"
                    >
                      <path
                        fill-rule="evenodd"
                        d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z"
                        clip-rule="evenodd"
                      ></path>
                    </svg>
                  </button>
                </div>
              </div>
              <button @click="closeModal()" class="text-gray-600 scale-105 hover:text-gray-700">
                <span class="sr-only">CloseModal</span>
                <svg
                  class="h-6 w-6"
                  x-description="Heroicon name: x"
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                  aria-hidden="true"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M6 18L18 6M6 6l12 12"
                  ></path>
                </svg>
              </button>
            </div>

            <!-- Sidebar Content -->
            <header>
              <nav class="flex items-start flex-col mt-6 gap-y-3">
                <div class="ml-4 mb-3s font-semibold dark:text-slate-200">
                  From:
                  <div
                    class="inline-flex items-center gap-1 rounded-full bg-violet-50 dark:bg-vivid_blue px-2 py-1 text-xs font-semibold text-slate-600 dark:text-slate-200"
                  >
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      id="Layer_1"
                      class="m-1 fill-[#414141] dark:fill-slate-400"
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
                    <span
                      v-if="emailsStore.selectedEmail"
                      v-html="
                        HighlighWord(emailsStore.selectedEmail.from, emailsStore.currentSearchTerm)
                      "
                    ></span>
                  </div>
                </div>
                <div class="inline-flex ml-4 font-semibold dark:text-slate-200">
                  To:
                  <div
                    v-if="emailsStore.selectedEmail"
                    class="ml-6 grid grid-cols-2 gap-y-1 gap-x-2"
                  >
                    <span
                      v-for="(email, index) in SeparateEmailsByCommas(emailsStore.selectedEmail.to)"
                      :key="index"
                      class="inline-flex rounded-full bg-violet-50 dark:bg-vivid_blue px-2 py-1 text-xs font-semibold text-gray-600 dark:text-slate-200"
                    >
                      <svg
                        xmlns="http://www.w3.org/2000/svg"
                        xmlns:xlink="http://www.w3.org/1999/xlink"
                        version="1.1"
                        id="Capa_1"
                        x="0px"
                        y="0px"
                        viewBox="0 0 512 512"
                        xml:space="preserve"
                        class="m-1 fill-[#414141] dark:fill-slate-400"
                        width="12"
                        height="12"
                      >
                        <g>
                          <circle cx="256" cy="128" r="128" />
                          <path
                            d="M256,298.667c-105.99,0.118-191.882,86.01-192,192C64,502.449,73.551,512,85.333,512h341.333   c11.782,0,21.333-9.551,21.333-21.333C447.882,384.677,361.99,298.784,256,298.667z"
                          />
                        </g>
                      </svg>
                      {{ email }}
                    </span>
                  </div>
                </div>
              </nav>
              <hr class="border-t border-slate-300 dark:border-slate-700 mt-4 shadow-xl" />
            </header>

            <main class="mt-4 px-10 overflow-x-hidden overflow-y-auto custom-scrollbar">
              <h2
                v-if="emailsStore.selectedEmail"
                class="mt-1 mb-4 mr-8 text-right rounded-full text-md font-bold uppercase text-slate-800 dark:text-slate-200"
              >
                {{ ConvertDateFormat(emailsStore.selectedEmail.date) }}
              </h2>
              <h2
                v-if="emailsStore.selectedEmail"
                class="font-bold text-2xl dark:text-slate-300"
                v-html="
                  HighlighWord(emailsStore.selectedEmail.subject, emailsStore.currentSearchTerm)
                "
              ></h2>
              <p
                v-if="emailsStore.selectedEmail"
                class="mt-2 text-gray-900 scroll-smooth font-md text-md dark:text-slate-200 my-14"
                v-html="
                  HighlighWord(emailsStore.selectedEmail.content, emailsStore.currentSearchTerm)
                "
              ></p>
            </main>
            <hr class="border-t border-slate-300 mt-4 dark:border-slate-700 shadow-md" />
          </div>
        </div>
      </section>
    </div>
  </Transition>
</template>
