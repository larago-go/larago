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

            <div v-else class="relative overflow-x-auto shadow-md sm:rounded-lg">
               <table class="w-full text-sm text-left text-gray-500 dark:text-gray-400">
                  <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
                    <tr>
                      <th scope="col" class="px-6 py-3">
                        Username
                      </th> 
                      <th scope="col" class="px-6 py-3">
                        Email
                      </th>
                      <th scope="col" class="px-6 py-3">
                        Role
                      </th>
                      <th scope="col" class="px-6 py-3">
                      <span class="sr-only">Edit</span>
                      </th>
                      <th scope="col" class="px-6 py-3">
                      <span class="sr-only">Delete</span>
                      </th>
                    </tr>
                </thead>
                <tbody>
                  <tr v-for="list in lists" :key="list" class="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600">
                    <th scope="row" class="px-6 py-4 font-medium text-gray-900 dark:text-white whitespace-nowrap">
                      {{ list.Name }}
                    </th> 
                    <td class="px-6 py-4">
                      {{ list.Email }}
                    </td>
                    <td class="px-6 py-4">
                      {{ list.Role }}
                    </td>
                    <td class="px-6 py-4 text-right">
                      <router-link class="inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500" :to="{ name: 'users_list_prev', params: { id: list.ID }}">Edit</router-link>
                    </td>
                    <td class="px-6 py-4 text-right">

              <button type="submit" class="inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500" @click="deleteId(list.ID)">Delete</button>
                    </td>
                </tr>
              </tbody>
          </table>
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

            lists: null,
            session_id: '',
            session_name: '',
            error: '',
        }
    },

async created() { 

  try {
  
    await Connect.get("/users/api/list")
            
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
                this.lists = response.data.list;
                    
              }
             
             }      
      })

  } catch (error) {
      
    console.log(error);
    
  }
    
},

  methods: {

    async deleteId(id) {
          
          Connect.defaults.headers.post['X-CSRF-Token'] = this.csrf;
          await Connect.get("/users/api/list/" + id + "/delete");
          window.location.href = '/users/list';
      },
    
  }


}
</script>
<style>
.checkbox:checked + .check-icon {
    display: flex;
}
</style>  
