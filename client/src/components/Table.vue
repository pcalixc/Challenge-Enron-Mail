eslint-disable vue/multi-word-component-names
<script setup lang="ts">
import { ref,defineProps } from 'vue'
import { LottieAnimation } from 'lottie-web-vue'
import animationData from '../assets/Animation.json'
import { ConvertDateFormat } from '../utils/functions'
import { highlighWord } from '../utils/functions'

interface Props {
  getData: Function
  emails : []
  asigneSelectedContent : Function
  isLoading: boolean
  totalResults: number
  totalPages: number
}

const props = defineProps<Props>()
const currentPage = ref(1)
const searchTerm = ref('')

</script>

<template>
  <div class="overflow-x-auto flex items-center justify-center font-sans overflow-hidden">
    <div class="w-full lg:w-5/6">
      <div class="flex flex-col mt-8 h-[67vh]">
        <div class="-mx-4 -my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
          <div class="inline-block min-w-full py-2 align-middle md:px-6 lg:px-8">
            <div class="overflow-hidden dark:border-gray-700 md:rounded-tl-lg">
              <table
                v-if="props.totalResults > 0"
                class="min-w-max w-[85%] rounded-t-xl m-auto mt-4 bg-white dark:bg-slate-600 dark:divide-gray-700 table-auto shadow-md"
              >
                <thead>
                  <tr
                    class="bg-slate-300 dark:bg-black dark:text-slate-100 text-gray-600 uppercase text-sm leading-normal rounded-t-xl"
                  >
                    <th class="py-3 px-6 text-left rounded-tl-xl">
                      <div class="gap-1 flex">
                        <svg
                          xmlns="http://www.w3.org/2000/svg"
                          class="mt-1 dark:fill-white fill-[#414141]"
                          id="Filled"
                          viewBox="0 0 24 24"
                          width="14"
                          height="14"
                        >
                          <path
                            d="M23.954,5.542,15.536,13.96a5.007,5.007,0,0,1-7.072,0L.046,5.542C.032,5.7,0,5.843,0,6V18a5.006,5.006,0,0,0,5,5H19a5.006,5.006,0,0,0,5-5V6C24,5.843,23.968,5.7,23.954,5.542Z"
                          />
                          <path
                            d="M14.122,12.546l9.134-9.135A4.986,4.986,0,0,0,19,1H5A4.986,4.986,0,0,0,.744,3.411l9.134,9.135A3.007,3.007,0,0,0,14.122,12.546Z"
                          />
                        </svg>
                        Subject
                      </div>
                    </th>
                    <th class="py-3 px-6 text-left">
                      <div class="gap-1 flex">
                        <svg
                          xmlns="http://www.w3.org/2000/svg"
                          id="Layer_1"
                          class="mt-1 dark:fill-white fill-[#414141]"
                          data-name="Layer 1"
                          viewBox="0 0 24 24"
                          width="14"
                          height="14"
                        >
                          <path
                            d="m20,0H4C1.794,0,0,1.794,0,4v12c0,2.206,1.794,4,4,4h2.923l3.748,3.156c.382.34.862.509,1.338.509.467,0,.931-.163,1.292-.485l3.847-3.18h2.852c2.206,0,4-1.794,4-4V4c0-2.206-1.794-4-4-4Zm-3.293,9.826c-.391.391-1.023.391-1.414,0l-2.293-2.293v7.467c0,.553-.448,1-1,1s-1-.447-1-1v-7.367l-2.293,2.293c-.195.195-.451.293-.707.293s-.512-.098-.707-.293c-.391-.391-.391-1.023,0-1.414l2.636-2.637c1.17-1.17,3.072-1.17,4.242,0l2.536,2.536c.391.391.391,1.023,0,1.414Z"
                          />
                        </svg>

                        From
                      </div>
                    </th>
                    <th class="py-3 px-6 text-left flex gap-1">
                      <div class="gap-1 flex">
                        <svg
                          xmlns="http://www.w3.org/2000/svg"
                          id="Layer_1"
                          class="mt-1 dark:fill-white fill-[#414141]"
                          data-name="Layer 1"
                          viewBox="0 0 24 24"
                          width="14"
                          height="14"
                        >
                          <path
                            d="m21,0H3C1.346,0,0,1.346,0,3v17h6.923l3.749,3.157c.382.339.861.507,1.337.507.468,0,.931-.163,1.292-.484l3.848-3.18h6.852V3c0-1.654-1.346-3-3-3Zm-7.594,14.419c-.388.387-.897.581-1.406.581s-1.018-.193-1.405-.58l-3.3-3.3,1.414-1.414,2.291,2.291v-6.997h2v7.008l2.291-2.302,1.414,1.414-3.299,3.299Z"
                          />
                        </svg>
                        To
                      </div>
                    </th>

                    <th class="py-3 px-6 text-left">
                      <div class="gap-1 flex">
                        <svg
                          xmlns="http://www.w3.org/2000/svg"
                          id="Filled"
                          class="mt-1 dark:fill-white fill-[#414141]"
                          viewBox="0 0 24 24"
                          width="14"
                          height="14"
                        >
                          <path
                            d="M0,19a5.006,5.006,0,0,0,5,5H19a5.006,5.006,0,0,0,5-5V10H0Zm17-4.5A1.5,1.5,0,1,1,15.5,16,1.5,1.5,0,0,1,17,14.5Zm-5,0A1.5,1.5,0,1,1,10.5,16,1.5,1.5,0,0,1,12,14.5Zm-5,0A1.5,1.5,0,1,1,5.5,16,1.5,1.5,0,0,1,7,14.5Z"
                          />
                          <path
                            d="M19,2H18V1a1,1,0,0,0-2,0V2H8V1A1,1,0,0,0,6,1V2H5A5.006,5.006,0,0,0,0,7V8H24V7A5.006,5.006,0,0,0,19,2Z"
                          />
                        </svg>
                        Date
                      </div>
                    </th>
                    <th class="py-3 px-6 text-left rounded-tr-xl"></th>
                  </tr>
                </thead>
                <tbody class="text-gray-600 text-sm font-light shadow-inner">
                  <tr
                    v-for="(data, index) in props.emails"
                    :key="index"
                    :class="{
                      ' bg-white dark:bg-[#1a1a40] dark:hover:bg-slate-900  ': index % 2 === 0,
                      'bg-zinc-50 dark:bg-[#082032] dark:hover:bg-slate-900': index % 2 !== 0
                    }"
                    class="border-b border-gray-200 dark:border-slate-900 hover:bg-gray-100 dark:text-slate-200"
                  >
                    <td class="py-3 pl-6 text-left whitespace-nowrap">
                      <span
                        :title="data._source.subject"
                        class="font-xs font-semibold block truncate text-ellipsis w-[12rem]"
                        v-html="highlighWord(data._source.subject, searchTerm)"
                      ></span>
                    </td>
                    <td class="py-3 px-6 text-left">
                      <span
                        :title="data._source.from"
                        v-html="highlighWord(data._source.from, searchTerm)"
                      ></span>
                    </td>
                    <td class="py-3 px-6 text-left">
                      <span
                        :title="data._source.to"
                        class="font-xs font-normal block truncate text-ellipsis w-[15rem]"
                        v-html="highlighWord(data._source.to, searchTerm)"
                      ></span>
                    </td>
                    <td class="py-3 px-6 text-left">
                      <span
                        class="inline-flex items-center gap-1 dark:bg-[#0101040e] rounded-xl bg-gray-50 px-2 py-1 text-xs font-semibold"
                      >
                        {{ ConvertDateFormat(data._source.date) }}
                      </span>
                    </td>
                    <td class="py-3 px-6 text-left">
                      <button
                        @click="
                          props.asigneSelectedContent(index)
                        "
                        class="px-2 py-1 font-normal dark:bg-slate-900 dark:text-slate-300 border dark:border-slate-700 dark:bg-[#28287552] border-purple-500 text-violet-600 bg-violet-50 rounded-xl transition duration-300 hover:scale-105 focus:outline-none"
                      >
                        View Details
                      </button>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
        <div
          v-if="props.totalResults == 0 && props.isLoading == false"
          class="h-full flex justify-center align-middle mt-24"
        >
          <lottie-animation :animation-data="animationData" :loop="true" class="w-72 h-72" />
        </div>
        <div v-if="props.isLoading == true" class="h-full flex justify-center align-middle mt-24">
          <svg
            role="status"
            class="inline h-12 w-12 animate-spin mt-24 text-gray-200 dark:text-gray-600 fill-[#9292FF]"
            viewBox="0 0 100 101"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z"
              fill="currentColor"
            />
            <path
              d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z"
              fill="currentFill"
            />
          </svg>
        </div>
      </div>

      <div
        id="pagination"
        v-if="props.totalResults > 0"
        class="flex items-center justify-center gap-2 mt-6 mx-24"
      >
        <!-- Botón de página "Previous" -->
        <button
          @click="props.getData(currentPage - 1, searchTerm)"
          :disabled="currentPage <= 1"
          :class="[
            'flex items-center px-5 py-2 text-sm capitalize transition-colors duration-200   rounded-xl gap-x-2',
            currentPage == 1
              ? ' bg-gray-50 text-gray-300 dark:bg-[#082032] dark:text-slate-500'
              : 'hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-[#1a1a40]   dark:text-slate-200  dark:hover:bg-gray-900'
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

        <div v-if="currentPage > 0 && currentPage < 3" class="items-center hidden md:flex gap-x-3">
          <button
            v-for="pageNumber in Math.min(3, props.totalPages)"
            :key="pageNumber"
            @click="props.getData(pageNumber)"
            :class="[
              'px-2 py-1 text-sm rounded-xl ',
              currentPage === pageNumber
                ? 'text-blue-500  bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 '
                : ' hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-[#1a1a40]   dark:text-slate-200  dark:hover:bg-gray-900'
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
              currentPage === props.totalPages - 1
                ? 'text-blue-500 bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900  '
                : ' hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-[#1a1a40]   dark:text-slate-200  dark:hover:bg-gray-900'
            ]"
          >
            {{ props.totalPages }}
          </button>
        </div>
        <!-- Botones de página -->
        <div v-if="currentPage > 2">
          <section v-if="props.totalPages <= 5" class="items-center hidden md:flex gap-x-3">
            <!-- Mostrar todos los botones de página si hay 5 páginas o menos -->
            <button
              v-for="pageNumber in props.totalPages"
              :key="pageNumber"
              @click="props.getData(pageNumber)"
              class="px-2 py-1 text-sm rounded-xl text-blue-500 bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900"
              :class="{
                'text-blue-500 bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 ':
                  currentPage === pageNumber
              }"
            >
              {{ pageNumber }}
            </button>
          </section>
          <section v-else class="items-center hidden md:flex gap-x-3">
            <!-- Mostrar botones de página con paginación "compacta" -->
            <button
              @click="props.getData(1)"
              class="px-2 py-1 text-sm rounded-xl hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-[#1a1a40] dark:text-slate-200 dark:hover:bg-gray-900"
              :class="{
                'text-blue-500  bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 ':
                  currentPage === 1
              }"
            >
              1
            </button>
            <span v-if="currentPage > 3" class="dark:text-slate-100">...</span>
            <template v-if="currentPage === 1">
              <button
                @click="props.getData(2)"
                class="px-2 py-1 text-sm rounded-xl hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-[#1a1a40] dark:text-slate-200 dark:hover:bg-gray-900"
                :class="{
                  'text-blue-500 bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 ':
                    currentPage === 2
                }"
              >
                2
              </button>
            </template>
            <template v-else>
              <button
                @click="props.getData(currentPage - 1)"
                class="px-2 py-1 text-sm rounded-xl hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-[#1a1a40] dark:text-slate-200 dark:hover:bg-gray-900"
                :class="{
                  'text-blue-500 bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 ':
                    currentPage === currentPage - 1
                }"
              >
                {{ currentPage - 1 }}
              </button>
            </template>
            <button
              @click="props.getData(currentPage)"
              class="px-2 py-1 text-sm rounded-xl text-blue-500 bg-blue-100/60 hover:bg-gray-300 bg-slate-200 shadow-xs dark:bg-[#1a1a40] dark:text-slate-200 dark:hover:bg-gray-900"
              :class="{
                'text-blue-500 bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 ':
                  currentPage === currentPage
              }"
            >
              {{ currentPage }}
            </button>
            <button
              @click="props.getData(currentPage + 1)"
              class="px-2 py-1 text-sm rounded-xl hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-[#1a1a40] dark:text-slate-200 dark:hover:bg-gray-900"
              :class="{
                'text-blue-500 bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 ':
                  currentPage === currentPage + 1
              }"
            >
              {{ currentPage + 1 }}
            </button>
            <span v-if="currentPage + 2 < props.totalPages" class="dark:text-slate-100">...</span>
            <button
              @click="props.getData(props.totalPages)"
              class="px-2 py-1 text-sm rounded-xl hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-[#1a1a40] dark:text-slate-200 dark:hover:bg-gray-900"
              :class="{
                'text-blue-500 bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 ':
                  currentPage === props.totalPages
              }"
            >
              {{ props.totalPages }}
            </button>
          </section>
        </div>
        <!-- Botón de página "Next" -->
        <button
          @click="props.getData(currentPage + 1, searchTerm)"
          :disabled="currentPage == props.totalPages"
          :class="[
            'flex items-center px-5 py-2 text-sm capitalize transition-colors duration-200  rounded-xl gap-x-2 ',
            currentPage == totalPages
              ? ' text-blue-500 bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900'
              : 'hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-[#1a1a40]   dark:text-slate-200   dark:hover:bg-gray-900'
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
    </div>
  </div>
  <div id="results" class="flex flex-row align-middle justify-center mt-5">
    <p class="text-slate-700 text-xs dark:text-slate-200 underline">
      {{ props.totalResults }} emails founded
    </p>
  </div>
</template>
