<!-- eslint-disable vue/multi-word-component-names -->
<script setup lang="ts">
import { ref, onBeforeMount } from 'vue'
// import animation from '@/assets/planets.json'
import Header from '@/components/Header.vue'
import  type {Email, Hit} from '@/types/index'
import ConnectionError from '@/components/indicators/ConnectionError.vue'
import NoData from '@/components/indicators/NoData.vue'
import Loading from '@/components/indicators/Loading.vue'
import EmailModal from '@/components/EmailModal.vue'
import TotalResults from '@/components/TotalResults.vue'
import { ConvertDateFormat, HighlighWord } from '../utils/functions'

const open = ref(false)
const selectedEmail = ref<Email | null>();
const emails = ref<Hit[]>([]);
const selectedEmailIndex = ref(0)
const currentSearchTerm = ref('')
const totalResults = ref(0)
const currentPage = ref(1)
var totalPages = ref(0)
const amountEmailsByPage = ref(7)
const isLoading = ref(false)
const conectionError= ref(false)


const getData = async (pageNumber: number, searchTerm?: string) => {
  searchTerm = typeof searchTerm === 'undefined' ? '' : searchTerm
  let pn = pageNumber.toString()
  currentSearchTerm.value = searchTerm
  isLoading.value = true
  try {
    const response = await fetch(
      `http://${import.meta.env.VITE_API_URL}/emails?page=${pn}&max=${amountEmailsByPage.value}&term=${searchTerm}`
    )

    const data = await response.json()

    emails.value = data.hits.hits
    totalResults.value = data.hits.Total.value
  } catch (error) {
    console.error('Error fetching data:', error, Response)
    conectionError.value=true
  }
  isLoading.value = false
  currentPage.value = pageNumber
  totalPages.value = Math.ceil(totalResults.value / amountEmailsByPage.value)
}

const asigneSelectedContent = (index: number) => {
  selectedEmail.value = emails.value[index]._source
  selectedEmailIndex.value = index
  open.value=true
  console.log(index, selectedEmail.value)
}

onBeforeMount(() => {
  getData(1)
})
</script>

<template>

  <Header :getData="getData" />

  <!-- table component -->
  <div class="overflow-x-auto flex items-center justify-center font-sans overflow-hidden">
    <div class="w-full lg:w-5/6">
      <div class="flex flex-col mt-8 h-[68vh]">
        <div class="-mx-4 -my-2  sm:-mx-6 lg:-mx-8">
          <div class="inline-block min-w-full py-2 align-middle md:px-6 lg:px-8">
            <div class="overflow-hidden dark:border-gray-700 md:rounded-tl-lg">
              <table
                v-if="totalResults > 0"
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
                          width="13"
                          height="13"
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
                          width="13"
                          height="13"
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
                          width="13"
                          height="13"
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
                          class="mt-1 ml-2 dark:fill-white fill-[#414141]"
                          viewBox="0 0 24 24"
                          width="13"
                          height="13"
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
                    v-for="(data, index) in emails"
                    :key="index"
                    :class="{
                      ' bg-white dark:bg-ligth_blue dark:hover:bg-slate-900  ': index % 2 === 0,
                      'bg-zinc-50 dark:bg-deep_blue dark:hover:bg-slate-900': index % 2 !== 0
                    }"
                    class="border-b border-gray-200 dark:border-slate-900 hover:bg-gray-100 dark:text-slate-200"
                  >
                    <td class="py-3 pl-6 text-left whitespace-nowrap">
                      <span
                        :title="data._source.subject"
                        class="font-xs font-semibold block truncate text-ellipsis w-[12rem]"
                        v-html="HighlighWord(data._source.subject, currentSearchTerm)"
                      ></span>
                    </td>
                    <td class="py-3 px-6 text-left">
                      <span
                        :title="data._source.from"
                        v-html="HighlighWord(data._source.from, currentSearchTerm)"
                      ></span>
                    </td>
                    <td class="py-3 px-6 text-left">
                      <span
                        :title="data._source.to"
                        class="font-xs font-normal block truncate text-ellipsis w-[15rem]"
                        v-html="HighlighWord(data._source.to, currentSearchTerm)"
                      ></span>
                    </td>
                    <td class="py-3 px-6 text-left">
                      <span
                        class="inline-flex items-center gap-1 dark:bg-transparent_blue rounded-xl bg-gray-50 px-2 py-1 text-xs font-semibold"
                      >
                        {{ ConvertDateFormat(data._source.date) }}
                      </span>
                    </td>
                    <td class="py-3 px-6 text-left">
                      <button
                        @click="asigneSelectedContent(index)"
                        class="px-2 py-1 font-normal dark:bg-slate-900 dark:text-slate-300 border dark:border-slate-700 dark:bg-transparent_blue border-purple-500 text-violet-600 bg-violet-50 rounded-xl transition duration-300 hover:scale-105 focus:outline-none"
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

        <NoData  v-if="totalResults == 0 && isLoading == false && conectionError==false"  />
        
        <ConnectionError v-if="conectionError"/>

        <Loading v-if="isLoading" />

      </div>
  <!-- <lottie-animation :animation-data="animation" :loop="true" class="w-72 h-72 opacity-10  fixed" /> -->

      <div
        id="pagination"
        v-if="totalResults > 0"
        class="flex items-center justify-center gap-2 mt-6 mx-24"
      >
        <button
          @click="getData(currentPage - 1, currentSearchTerm)"
          :disabled="currentPage <= 1"
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
        <div v-if="currentPage > 0 && currentPage < 3" class="items-center hidden md:flex gap-x-3">
          <button
            v-for="pageNumber in Math.min(3, totalPages)"
            :key="pageNumber"
            @click="getData(pageNumber)"
            :class="[
              'px-2 py-1 text-sm rounded-xl ',
              currentPage === pageNumber
                ? 'text-blue-500  bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 '
                : ' hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue   dark:text-slate-200  dark:hover:bg-gray-900'
            ]"
          >
            {{ pageNumber }}
          </button>

          <button
            v-if="totalPages > 5"
            class="px-2 py-1 text-sm text-gray-500 rounded-xl hover:bg-gray-100 dark:text-slate-100"
            disabled
          >
            ...
          </button>

          <button
            @click="getData(totalPages)"
            :class="[
              'px-2 py-1 text-sm rounded-xl ',
              currentPage === totalPages - 1
                ? 'text-blue-500 bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900  '
                : ' hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue   dark:text-slate-200  dark:hover:bg-gray-900'
            ]"
          >
            {{ totalPages }}
          </button>
        </div>
     
        <div v-if="currentPage > 2">
          <section v-if="totalPages <= 5" class="items-center hidden md:flex gap-x-3">
   
            <button
              v-for="pageNumber in totalPages"
              :key="pageNumber"
              @click="getData(pageNumber)"
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
        
            <button
              @click="getData(1)"
              class="px-2 py-1 text-sm rounded-xl hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue dark:text-slate-200 dark:hover:bg-gray-900"
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
                @click="getData(2)"
                class="px-2 py-1 text-sm rounded-xl hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue dark:text-slate-200 dark:hover:bg-gray-900"
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
                @click="getData(currentPage - 1)"
                class="px-2 py-1 text-sm rounded-xl hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue dark:text-slate-200 dark:hover:bg-gray-900"
                :class="{
                  'text-blue-500 bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 ':
                    currentPage === currentPage - 1
                }"
              >
                {{ currentPage - 1 }}
              </button>
            </template>
            <button
              @click="getData(currentPage)"
              class="px-2 py-1 text-sm rounded-xl text-blue-500 bg-blue-100/60 hover:bg-gray-300 bg-slate-200 shadow-xs dark:bg-ligth_blue dark:text-slate-200 dark:hover:bg-gray-900"
              :class="{
                'text-blue-500 bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 ':
                  currentPage === currentPage
              }"
            >
              {{ currentPage }}
            </button>
            <button
              @click="getData(currentPage + 1)"
              class="px-2 py-1 text-sm rounded-xl hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue dark:text-slate-200 dark:hover:bg-gray-900"
              :class="{
                'text-blue-500 bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 ':
                  currentPage === currentPage + 1
              }"
            >
              {{ currentPage + 1 }}
            </button>
            <span v-if="currentPage + 2 < totalPages" class="dark:text-slate-100">...</span>
            <button
              @click="getData(totalPages)"
              class="px-2 py-1 text-sm rounded-xl hover:bg-gray-300 bg-slate-200 shadow-xs text-gray-800 dark:bg-ligth_blue dark:text-slate-200 dark:hover:bg-gray-900"
              :class="{
                'text-blue-500 bg-blue-100/60 dark:bg-slate-200 dark:text-slate-900 ':
                  currentPage === totalPages
              }"
            >
              {{ totalPages }}
            </button>
          </section>
        </div>
   
        <button
          @click="getData(currentPage + 1, currentSearchTerm)"
          :disabled="currentPage == totalPages"
          :class="[
            'flex items-center px-5 py-2 text-sm capitalize transition-colors duration-200  rounded-xl gap-x-2 ',
            currentPage == totalPages
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

    </div>
  </div>

  <TotalResults v-if="!conectionError"
  :totalResults="totalResults"/>


  <EmailModal v-if="open"
  :selectedEmail="selectedEmail"
  :asigneSelectedContent="asigneSelectedContent"
  :currentSearchTerm="currentSearchTerm"
  :selectedEmailIndex="selectedEmailIndex"
  @close="open=false"
  />

</template>

<style>

::-webkit-scrollbar {
  width: 10px;
  height: 10px;
}


::-webkit-scrollbar-thumb {
  background-color: #9292FF;
  border-radius: 5px;
}

::-webkit-scrollbar-thumb:hover {
  background-color: #555;
}

::-webkit-scrollbar-track {
  background-color: #f1f1f1;
}
</style>
