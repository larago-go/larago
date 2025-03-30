<script setup>
import { ref } from 'vue';

import { useRouter } from 'vue-router';

import SidebarAdmin from '../SidebarAdmin.vue';

import NavbarAdmin from '../NavbarAdmin.vue';

import FooterAdmin from './FooterAdmin.vue';

import Connect from '../../confconnect';

const router = useRouter();

const session = ref({
  session_id: '',
  session_name: '',
});

const created = () => {
  try {
    Connect.get('/home/api')
      .then((response) => {
          session.value.session_id = response.data.session_id;
          session.value.session_name = response.data.session_name;
      });
  } catch (error) {
    console.log(error);
  }
};

created();

</script>
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

        <FooterAdmin />

      </div>
    </div>
  </div>
</template>
