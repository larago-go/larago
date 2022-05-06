<template>
  <div class="leading-normal tracking-normal" id="main-body">
    <div class="flex flex-wrap">

      <SidebarAdmin />

      <div class="w-full bg-gray-100 pl-0 lg:pl-64 min-h-screen" :class="sideBarOpen ? 'overlay' : ''" id="main-content">

        <NavbarAdmin />
        <div class="p-6 bg-gray-100 mb-20">
          <div class="py-12">   

            <!-- alert  -->
            <div v-if="error" class="flex p-4 mb-4 text-sm text-red-700 bg-red-100 rounded-lg dark:bg-red-200 dark:text-red-800" role="alert">
               <svg class="inline flex-shrink-0 mr-3 w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z" clip-rule="evenodd"></path></svg>
               <div>
                 <span class="font-medium" >{{ error }}</span>
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
                  <label for="first-name" class="block text-sm font-medium text-gray-700">Name</label>
                  <input type="text" name="first-name" id="first-name"  v-model="form.name" required autocomplete="name" class="mt-1 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md" />
                </div>

                <div class="col-span-6 sm:col-span-3">
                  <label for="last-name" class="block text-sm font-medium text-gray-700">Role</label>
                  <input type="text" name="last-name" id="last-name" v-model="form.role"  autocomplete="role" class="mt-1 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md" />
                </div>

                <div class="col-span-6 sm:col-span-4">
                  <label for="email-address" class="block text-sm font-medium text-gray-700">Email address</label>
                  <input type="text" name="email-address" id="email-address" v-model="form.email"  required autocomplete="email" class="mt-1 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md" />
                </div>

                <div class="col-span-6 sm:col-span-4">
                  <label for="last-name" class="block text-sm font-medium text-gray-700">Password</label>
                  <input type="password" name="last-name" id="last-name" v-model="form.password" autocomplete="password" class="mt-1 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md" />
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
        
    <Footer />

      </div>
    </div>
  </div>
</template>

<script>
import { mapState }  from 'vuex'
import SidebarAdmin from '@/Admin/Sidebaradmin.vue'
import NavbarAdmin from '@/Admin/Navbaradmin.vue'
import Footer from '@/Admin/Views/Footer.vue'
import Connect from '@/config_conn'
export default {
  computed: {
    ...mapState(['sideBarOpen'])
  },
  components: {
    SidebarAdmin,
    NavbarAdmin,
    Footer
  },

    data() {
        return {

            session_id: '',
            session_name: '',
            id: '',
            error: '',

            form:  
              {
                name: '',
                email: '',
                role: '',
                password: '', 
              }
        }
    },

async created() { 

  try {
  
    await Connect.get("/users/api/list/" + this.$route.params.id)
            
      .then(response => { 
                // JSON responses are automatically parsed.
             if(response.data.error){

                this.error = response.data.error;
             
             } else {

              if(response.data.csrf == "redirect_auth_login"){  
                
                this.$router.push({ name: 'login' }); 
                  
              } else {
                    
                this.csrf = response.data.csrf;
                this.session_id = response.data.session_id;
                this.session_name = response.data.session_name;
                this.id = response.data.id;
                this.form.name = response.data.name;
                this.form.email = response.data.email;
                this.form.role = response.data.role;
                    
              }

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
       
         await Connect.post("/users/list/" + this.$route.params.id + "/edit", this.form)
                 
           .then(response => { 
                  
                  this.name = '';
                  this.email = '';
                  this.role = '';
                  this.password = '';
               
                 })
         
         this.$router.push({ name: 'users_list' });
       
         } catch (error) {
      
          this.error = error;
    
        }

       },
    }
}
</script>
