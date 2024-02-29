<script setup lang="ts">
import { defineProps } from 'vue'
import ConnectionError from '@/components/indicators/ConnectionError.vue'
import NoData from '@/components/indicators/NoData.vue'
import Loading from '@/components/indicators/Loading.vue'
import Pagination from './Pagination.vue'
import { ConvertDateFormat, HighlighWord } from '../utils/emails.utilities'
import type { Hit } from '@/types'


// This function checks if a current SearchTerm is present 'only' in the content of the email. 
const isWordInContent = (
  subject: string,
  from: string,
  to: string,
  content: string,
  term: string
) => {
  const regex = new RegExp(term, 'gi')
  let a = regex.test(subject) || regex.test(from) || regex.test(to)
  if (a) {
    return !a
  }
  return regex.test(content)
}

interface Props {
  getData: Function
  emails: Hit[]
  asigneSelectedContent: Function
  currentSearchTerm: string
  isLoading: boolean
  totalResults: number
  conectionError: boolean
  totalPages: number
  currentPage: number
  suggestions: string[]
}

const props = defineProps<Props>()
</script>

<template>
  <div class="flex items-center justify-center font-sans overflow-hidden">
    <div class="w-full lg:w-5/6">
      <div class="flex flex-col h-[70vh] overflow-hidden mt-2 mb-0">
        <div class="overflow-x-auto px-9 md:overflow-auto sm:-mx-6 lg:-mx-8">
          <div class="min-w-[90%] py-0 align-middle md:px-6 lg:px-8">
            <div
              class="dark:border-gray-700 md:rounded-tl-lg rounded flex align-middle justify-center"
            >
              <table
                v-if="props.totalResults > 0 && !props.isLoading"
                class="w-[95%] rounded-t-xl bg-white shadow-md dark:bg-slate-600 mt-10 dark:divide-gray-700 min-w-max border-collapse block md:table"
              >
                <thead class="block md:table-header-group">
                  <tr
                    class="bg-slate-300 dark:bg-black dark:text-slate-100 text-gray-600 uppercase text-sm block md:table-row absolute -top-full md:top-auto -left-full md:left-auto md:relative"
                  >
                    <th class="py-3 px-6 rounded-tl-xl text-left block md:table-cell">
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
                    <th class="py-3 px-6 text-left block md:table-cell">
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
                    <th class="py-3 px-6 text-left block md:table-cell">
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
                    <th class="py-3 px-6 text-left block md:table-cell">
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
                    <th class="py-3 px-6 rounded-tr-xl text-left block md:table-cell"></th>
                  </tr>
                </thead>
                <tbody class="block md:table-row-group font-light text-sm text-gray-600">
                  <tr
                    v-for="(data, index) in props.emails"
                    :key="index"
                    :class="{
                      ' bg-white dark:bg-ligth_blue dark:hover:bg-slate-900  ': index % 2 === 0,
                      'bg-zinc-50 dark:bg-deep_blue dark:hover:bg-slate-900': index % 2 !== 0
                    }"
                    class="border-b border-gray-200 dark:border-slate-900 hover:bg-gray-100 dark:text-slate-200 block md:table-row"
                  >
                    <td class="p-2 text-left block md:table-cell">
                      <span class="inline-block w-1/3 md:hidden font-bold">
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
                      </span>
                      <span
                        :title="data._source.subject"
                        class="font-xs ml-3 font-semibold md:block truncate text-ellipsis w-[15rem]"
                        v-html="HighlighWord(data._source.subject, props.currentSearchTerm)"
                      ></span>
                    </td>
                    <td class="p-2 text-left block md:table-cell">
                      <span class="inline-block w-1/3 md:hidden font-bold">
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
                      </span>
                      <span
                        class="ml-4"
                        :title="data._source.from"
                        v-html="HighlighWord(data._source.from, currentSearchTerm)"
                      ></span>
                    </td>
                    <td class="p-2 text-left block md:table-cell">
                      <span class="inline-block w-1/3 md:hidden font-bold">
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
                      </span>
                      <span
                        :title="data._source.to"
                        class="font-xs ml-4 font-normal md:block truncate text-ellipsis w-[15rem] overflow-x-auto"
                        v-html="HighlighWord(data._source.to, currentSearchTerm)"
                      ></span>
                    </td>
                    <td class="p-2 text-left block md:table-cell">
                      <span class="inline-block w-1/3 md:hidden font-bold">
                        <div class="gap-1 flex">
                          <svg
                            xmlns="http://www.w3.org/2000/svg"
                            id="Filled"
                            class="mt-1 dark:fill-white fill-[#414141]"
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
                      </span>
                      <span
                        class="inline-flex ml-4 items-center gap-1 dark:bg-transparent_blue rounded-xl bg-gray-50 px-2 py-1 text-xs font-semibold"
                      >
                        {{ ConvertDateFormat(data._source.date) }}
                      </span>
                    </td>
                    <td class="p-2 text-left block md:table-cell">
                      <span class="inline-block w-1/3 md:hidden font-bold"></span>
                      <button
                        :class="[
                          'bg-gray-200  p-1.5 rounded-lg mx-3',
                          isWordInContent(
                            data._source.subject,
                            data._source.from,
                            data._source.to,
                            data._source.content,
                            currentSearchTerm
                          )
                            ? 'px-2 py-1 dark:bg-royal_purple dark:text-slate-950 border font-semibold dark:border-yellow-700 dark:bg-transparent_blue border-yellow-500 text-yellow-600 bg-yellow-50  rounded-lg transition duration-300 hover:scale-105 focus:outline-none'
                            : 'px-2 py-1 font-semibold dark:bg-slate-900 dark:text-slate-300 border dark:border-slate-700 dark:bg-transparent_blue border-purple-500 text-violet-600 bg-violet-50 rounded-lg transition duration-300 hover:scale-105 focus:outline-none'
                        ]"
                        @click="asigneSelectedContent(index)"
                      >
                        Open Email
                      </button>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>

        <!-- indicators -->

        <NoData
          v-if="
            props.totalResults == 0 && props.isLoading == false && props.conectionError == false
          "
          :currentSearchTerm="currentSearchTerm"
          :suggestions="suggestions"
          :getData="getData"
        />

        <ConnectionError v-if="props.conectionError" />

        <Loading v-if="props.isLoading" />
      </div>

      <Pagination
        v-if="!conectionError"
        :currentPage="props.currentPage"
        :totalResults="props.totalResults"
        :currentSearchTerm="currentSearchTerm"
        :totalPages="totalPages"
        :getData="getData"
      />
    </div>
  </div>
</template>
