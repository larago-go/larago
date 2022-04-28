<template>
  <div class="min-h-full flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
    <div class="max-w-md w-full space-y-8">
      <div>
        <img class="mx-auto h-12 w-auto" src="https://github.com/larago-go/larago/raw/master/larago-logo_git.png" alt="Workflow" />
        <h2 class="mt-6 text-center text-3xl font-extrabold text-gray-900">Forgot password</h2>
        <p class="mt-6 text-center text-3xl font-extrabold text-gray-900">To reset your password, enter your username or email address below. If your account is in the database, an email will be sent to your email address containing instructions on how to restore access.</p>
      </div>
      <form class="mt-8 space-y-6" @submit.prevent="submit">
        <div class="rounded-md shadow-sm -space-y-px">
          <div>
            <label for="email-address" class="sr-only">Email address</label>
            <input id="email-address" name="email" v-model="form.email" type="email" autocomplete="email" required="" class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm" placeholder="Email address" />
          </div>
        </div>


        <div>
          <button type="submit" class="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
            <span class="absolute left-0 inset-y-0 flex items-center pl-3">
              <LockClosedIcon class="h-5 w-5 text-indigo-500 group-hover:text-indigo-400" aria-hidden="true" />
            </span>
            Send
          </button>
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
            
          form: {
                email: '',
            }
        }
    },

async created() { 

  try {
  
    await Connect.get("/login/api/forgot_password")
            
      .then(response => { 
                // JSON responses are automatically parsed.
                  
                    
                this.csrf = response.data.csrf;
                    
            
      })  
    
  } catch (error) {
      
    console.log(error);
    
  }
    
}, 
    
  methods: {
       
    async submit() {

        Connect.defaults.headers.post['X-CSRF-Token'] = this.csrf;
       
         await Connect.post("/login/post_add", this.form)
                 
        .then(response => { 
                  
          this.email = '';      
               
        })

      window.location.href = '/';       
    
    },
    
  }

}
</script>
