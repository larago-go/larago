<template>
  <div class="min-h-full flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
    <div class="max-w-md w-full space-y-8">
      <div>
        <img class="mx-auto h-12 w-auto" src="https://github.com/larago-go/larago/raw/master/larago-logo_git.png" alt="Workflow" />
        <h2 class="mt-6 text-center text-3xl font-extrabold text-gray-900">Sign in to your account</h2>
      </div>
            <!-- alert  -->
            <div v-if="error" class="flex p-4 mb-4 text-sm text-red-700 bg-red-100 rounded-lg dark:bg-red-200 dark:text-red-800" role="alert">
               <svg class="inline flex-shrink-0 mr-3 w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z" clip-rule="evenodd"></path></svg>
               <div>
                 <span class="font-medium" >{{ error }}</span>
               </div>
            </div>
      <form class="mt-8 space-y-6" @submit.prevent="submit">
        <div class="rounded-md shadow-sm -space-y-px">
          <div>
            <label for="email-address" class="sr-only">Email address</label>
            <input id="email-address" name="email" v-model="form.email" type="email" autocomplete="email" required="" class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm" placeholder="Email address" />
          </div>
          <div>
            <label for="password" class="sr-only">Password</label>
            <input id="password" name="password" v-model="form.password" type="password" autocomplete="current-password" required="" class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-b-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm" placeholder="Password" />
          </div>
        </div>


        <div>
          <button type="submit" class="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
            <span class="absolute left-0 inset-y-0 flex items-center pl-3">
              <LockClosedIcon class="h-5 w-5 text-indigo-500 group-hover:text-indigo-400" aria-hidden="true" />
            </span>
            Sign in
          </button>
                               <router-link :to="'/login/forgot_password'">Reset password</router-link>
        </div>
      </form>
    </div>
  </div>
</template>

<script>

import { LockClosedIcon } from '@heroicons/vue/solid'

import Connect from '@/config_conn'

export default {
  
  components: {
    
    LockClosedIcon,
    
  },

    data() {
        return {

          csrf: '',
          error: '', 
          form: {

                email: '',
                password: '',

            }
        }
    },

async created() { 

  try {
  
    await Connect.get("/auth/api/login")
            
      .then(response => { 
                // JSON responses are automatically parsed.
              if(response.data.csrf == "redirect_home"){  
                
                this.$router.push({ name: 'welcome' }); 
                  
              } else {
                    
                this.csrf = response.data.csrf;
                    
              }
            
      })  
    
  } catch (error) {
      
    console.log(error);
    
  }
    
}, 
    
  methods: {
       
    async submit() {
     
      try {
      
        Connect.defaults.headers.post['X-CSRF-Token'] = this.csrf;
       
         await Connect.post("/auth/signin", this.form)
                 
        .then(response => { 
                  
          this.email = '';      
          this.password = '';
               
        })

      window.location.href = '/home';       
    
      } catch (error) {
      
       this.error = error;
    
      }

    },
    

  }

}
</script>
