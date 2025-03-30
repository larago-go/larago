<script setup>

import { ref } from 'vue';

import { useRouter } from 'vue-router';

import { LockClosedIcon } from '@heroicons/vue/solid';

import Connect from '../../confconnect';

const router = useRouter();

const datavw = ref({
  error: '',
  form: {
    name: '',
    email: '',
    password: '',

  },
});

const created = () => {
  try {
    Connect.get('/auth/api/register')
      .then((response) => {
        if (response.data.error != null) {
          datavw.value.error = response.data.error;
        }
      });
  } catch (error) {
    console.log(error);
  }
};

created();

const submit = () => {
  try {
    Connect.post('/auth/signup', datavw.value.form)
      .then((response) => {
        localStorage.setItem('token', response.data.token);
        localStorage.setItem('user_name', JSON.stringify(response.data.user_name));
        localStorage.setItem('user_id', JSON.stringify(response.data.user_id));
        localStorage.setItem('user_email', JSON.stringify(response.data.user_email));
        router.push({ name: 'login' });
      });
  } catch (error) {
    datavw.value.error = error;
  }
};
</script>
<template>
  <div class="min-h-full flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
    <div class="max-w-md w-full space-y-8">
      <div>
        <img class="mx-auto h-12 w-auto" src="https://github.com/larago-go/larago/raw/master/larago-logo_git.png" alt="Workflow" />
        <h2 class="mt-6 text-center text-3xl font-extrabold text-gray-900">Sign up your account</h2>
      </div>
            <!-- alert  -->
            <div v-if="datavw.error" class="flex p-4 mb-4 text-sm text-red-700 bg-red-100 rounded-lg dark:bg-red-200 dark:text-red-800" role="alert">
               <svg class="inline flex-shrink-0 mr-3 w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z" clip-rule="evenodd"></path></svg>
               <div>
                 <span class="font-medium" >{{ datavw.error }}</span>
               </div>
            </div>
      <form class="mt-8 space-y-6" @submit.prevent="submit">
        <div class="rounded-md shadow-sm -space-y-px">
          <div>
            <label for="name" class="sr-only">Name</label>
            <input id="name" name="name" v-model="datavw.form.name" type="text" autocomplete="name" required="" class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm" placeholder="Name" />
          </div>
          <div>
            <label for="email-address" class="sr-only">Email address</label>
            <input id="email-address" name="email" v-model="datavw.form.email" type="email" autocomplete="email" required="" class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm" placeholder="Email address" />
          </div>
          <div>
            <label for="password" class="sr-only">Password</label>
            <input id="password" name="password" v-model="datavw.form.password" type="password" autocomplete="current-password" required="" class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-b-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm" placeholder="Password" />
          </div>
        </div>


        <div>
          <button type="submit" class="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
            <span class="absolute left-0 inset-y-0 flex items-center pl-3">
              <LockClosedIcon class="h-5 w-5 text-indigo-500 group-hover:text-indigo-400" aria-hidden="true" />
            </span>
            Sign up
          </button>
        </div>
      </form>
    </div>
  </div>
</template>
