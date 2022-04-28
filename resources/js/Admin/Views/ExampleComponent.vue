<template>
  <div class="leading-normal tracking-normal" id="main-body">
    <div class="flex flex-wrap">

      <SidebarAdmin />

      <div class="w-full bg-gray-100 pl-0 lg:pl-64 min-h-screen" :class="sideBarOpen ? 'overlay' : ''" id="main-content">

        <NavbarAdmin />

        <div class="p-6 bg-gray-100 mb-20">
                    <div class="card-header">Example Component</div>

                    <div class="card-body">
                        I'm an example component.
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
export default {
  computed: {
    ...mapState(['sideBarOpen'])
  },
  components: {
    SidebarAdmin,
    NavbarAdmin,
    Footer
  }
}
</script>
