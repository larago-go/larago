
<template>
  <div class="leading-normal tracking-normal" id="main-body">
    <div class="flex flex-wrap">

      <SidebarAdmin />

      <div class="w-full bg-gray-100 pl-0 lg:pl-64 min-h-screen" :class="sideBarOpen ? 'overlay' : ''" id="main-content">

        <NavbarAdmin />

        <div class="p-6 bg-gray-100 mb-20">
          <div class="py-12">    
            <div class="max-w-7xl mx-auto sm:px-6 lg:px-8">
                <div class="bg-white overflow-hidden shadow-sm sm:rounded-lg">
                    <div class="p-6 bg-white border-b border-gray-200">
                        You're logged in!               
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
            session_name: ''

        }
    },
async created() { 

  try {
  
    await Connect.get("/home/api")
            
      .then(response => { 
                // JSON responses are automatically parsed.
              if(response.data.csrf == "redirect_auth_login"){  
                
                this.$router.push({ name: 'login' }); 
                  
              } else {
                    
                this.csrf = response.data.csrf;
                this.session_id = response.data.session_id;
                this.session_name = response.data.session_name;
                    
              }
            
      })  
  } catch (error) {
      
    console.log(error);
    
  }
    
},
  
}
</script>
