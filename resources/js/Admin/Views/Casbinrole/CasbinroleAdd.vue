<script setup>
import { ref } from 'vue';

import { useRouter } from 'vue-router';

import SidebarAdmin from '../../SidebarAdmin.vue';

import NavbarAdmin from '../../NavbarAdmin.vue';

import FooterAdmin from '../FooterAdmin.vue';

import Connect from '../../../confconnect';

const router = useRouter();

const datavw = ref({

  session_id: '',
  session_name: '',
  error: '',

  form:
    {
      rolename: '',
      path: '',
      method: '',
    },

});

const created = () => {
  try {
    Connect.get('/role/api/add')
      .then((response) => {
        if (response.data.error != null) {
          datavw.value.error = response.data.error;
        } else {
          datavw.value.session_id = response.data.session_id;
          datavw.value.session_name = response.data.session_name;
        }
      });
  } catch (error) {
    console.log(error);
  }
};

created();

const submit = () => {
  try {
    Connect.post('/role/post_add', datavw.value.form)
      .then((response) => {
        response.rolename = '';
        response.path = '';
        response.method = '';
        router.push({ name: 'role_list' });
      });
  } catch (error) {
    datavw.value.error = error;
  }
};
</script>
<template>
  <div class="leading-normal tracking-normal" id="main-body">
    <div class="flex flex-wrap">

      <SidebarAdmin />

      <div class="w-full bg-gray-100 pl-0 lg:pl-64 min-h-screen" :class="sideBarOpen ? 'overlay' : ''" id="main-content">

        <NavbarAdmin />
        <div class="p-6 bg-gray-100 mb-20">
          <div class="py-12">    

            <!-- alert  -->
            <div v-if="datavw.error" class="flex p-4 mb-4 text-sm text-red-700 bg-red-100 rounded-lg dark:bg-red-200 dark:text-red-800" role="alert">
               <svg class="inline flex-shrink-0 mr-3 w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z" clip-rule="evenodd"></path></svg>
               <div>
                 <span class="font-medium" >{{ datavw.error }}</span>
               </div>
            </div>

     <div v-else class="mt-10 sm:mt-0">
    <div class="md:grid md:grid-cols-3 md:gap-6">
      <div class="mt-5 md:mt-0 md:col-span-2">
        <form @submit.prevent="submit">
          <div class="shadow overflow-hidden sm:rounded-md">
            <div class="px-4 py-5 bg-white sm:p-6">
              <div class="grid grid-cols-6 gap-6">
                <div class="col-span-6 sm:col-span-3">
                  <label for="first-name" class="block text-sm font-medium text-gray-700">Role name</label>
                  <input type="text" name="first-name" id="first-name"  v-model="datavw.form.rolename" required autocomplete="name" class="mt-1 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md" />
                </div>

                <div class="col-span-6 sm:col-span-3">
                  <label for="last-name" class="block text-sm font-medium text-gray-700">Path (Example: /*)</label>
                  <input type="text" name="last-Path" id="last-Path" v-model="datavw.form.path"  autocomplete="Path" class="mt-1 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md" />
                </div>

                <div class="col-span-6 sm:col-span-4">
                  <label for="email-address" class="block text-sm font-medium text-gray-700">Method (GET,POST,PUT,PATCH,DELETE)</label>
                  <input type="text" name="Method" id="Method" v-model="datavw.form.method"  required autocomplete="Method" class="mt-1 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md" />
                </div>


              </div>
            </div>
            <div class="px-4 py-3 bg-gray-50 text-right sm:px-6">
              <button type="submit" class="inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">Save</button>
            </div>
          </div>
        </form>
      </div>
    </div>
  </div>
  </div>
    </div>
        
    <FooterAdmin />

      </div>
    </div>
  </div>
</template>
