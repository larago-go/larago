<script setup>

import DataTable from 'datatables.net-vue3';

import DataTablesLib from 'datatables.net';

import { ref, onMounted } from 'vue';

import { useRouter } from 'vue-router';

import SidebarAdmin from '../../SidebarAdmin.vue';

import NavbarAdmin from '../../NavbarAdmin.vue';

import FooterAdmin from '../FooterAdmin.vue';

import Connect from '../../../confconnect';

let dt;

DataTable.use(DataTablesLib);

const router = useRouter();

const datavw = ref({
  lists: [],
  session_id: '',
  session_name: '',
  error: '',
  csrf: '',
});

const table = ref();

const created = () => {
  try {
    Connect.get('/users/api/list')
      .then((response) => {
        if (response.data.error != null) {
          datavw.value.error = response.data.error;
        } else if (response.data.csrf === 'redirect_auth_login') {
          router.push({ name: 'login' });
        } else {
          datavw.value.csrf = response.data.csrf;
          datavw.value.session_id = response.data.session_id;
          datavw.value.session_name = response.data.session_name;
          datavw.value.lists = response.data.list;
        }
      });
  } catch (error) {
    console.log(error);
  }
};

created();

const deleteId = (id) => {
  Connect.defaults.headers.delete['X-CSRF-Token'] = datavw.value.csrf;
  Connect.delete("/users/api/list/" + id + "/delete");
  window.location.href = '/users/list';
};

onMounted(function () {
  dt = table.value.dt;
});

const columns = [
  { data: 'Name', title: 'Username' },
  { data: 'Email', title: 'Email' },
  { data: 'Role', title: 'Role' },
];

function add() {
  router.push({ name: 'users_add' });
}

function edit() {
  dt.rows({ selected: true }).every(function () {
    const idx = datavw.value.lists.indexOf(this.data());
    const itemId = dt.row(idx).data().ID;
    router.push({ name: 'users_list_prev', params: { id: itemId } });
  });
}

function remove() {
  dt.rows({ selected: true }).every(function () {
    const idx = datavw.value.lists.indexOf(this.data());
    const itemId = dt.row(idx).data().ID;
    deleteId(itemId);
  });
}

</script>
<style>
@import 'datatables.net-dt';
@import 'datatables.net-searchbuilder-dt';
</style>
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

            <div v-else class="p-8 mt-6 lg:mt-0 rounded shadow bg-white">
              <button @click="add" type="button" class="focus:outline-none text-white bg-green-700 hover:bg-green-800 focus:ring-4 focus:ring-green-300 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-green-600 dark:hover:bg-green-700 dark:focus:ring-green-800">Add</button>
              <button @click="edit" type="button" class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800">Edit</button>
              <button  @click="remove" type="button" class="focus:outline-none text-white bg-red-700 hover:bg-red-800 focus:ring-4 focus:ring-red-300 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-red-600 dark:hover:bg-red-700 dark:focus:ring-red-900">Delete</button>
               <DataTable
               :columns="columns"
               :data="datavw.lists"
               class="stripe hover"
               width="100%"
               :options="{ select: true,
                           dom: 'Qlfrtip'
                         }"
               ref="table"
               />
        </div>
      </div>
    </div>
        
    <FooterAdmin />

      </div>
    </div>
  </div>
</template>

<style>
.checkbox:checked + .check-icon {
    display: flex;
}
	/*Overrides for Tailwind CSS */

/*Form fields*/
.dataTables_wrapper select,
.dataTables_wrapper .dataTables_filter input {
color: #4a5568;
/*text-gray-700*/
padding-left: 1rem;
/*pl-4*/
padding-right: 1rem;
/*pl-4*/
padding-top: .5rem;
/*pl-2*/
padding-bottom: .5rem;
/*pl-2*/
line-height: 1.25;
/*leading-tight*/
border-width: 2px;
/*border-2*/
border-radius: .25rem;
border-color: #edf2f7;
/*border-gray-200*/
background-color: #edf2f7;
/*bg-gray-200*/
}

/*Row Hover*/
table.dataTable.hover tbody tr:hover,
table.dataTable.display tbody tr:hover {
background-color: #ebf4ff;
/*bg-indigo-100*/
}

/*Pagination Buttons*/
.dataTables_wrapper .dataTables_paginate .paginate_button {
font-weight: 700;
/*font-bold*/
border-radius: .25rem;
/*rounded*/
border: 1px solid transparent;
/*border border-transparent*/
}

/*Pagination Buttons - Current selected */
.dataTables_wrapper .dataTables_paginate .paginate_button.current {
color: #fff !important;
/*text-white*/
box-shadow: 0 1px 3px 0 rgba(0, 0, 0, .1), 0 1px 2px 0 rgba(0, 0, 0, .06);
/*shadow*/
font-weight: 700;
/*font-bold*/
border-radius: .25rem;
/*rounded*/
background: #667eea !important;
/*bg-indigo-500*/
border: 1px solid transparent;
/*border border-transparent*/
}

/*Pagination Buttons - Hover */
.dataTables_wrapper .dataTables_paginate .paginate_button:hover {
color: #fff !important;
/*text-white*/
box-shadow: 0 1px 3px 0 rgba(0, 0, 0, .1), 0 1px 2px 0 rgba(0, 0, 0, .06);
/*shadow*/
font-weight: 700;
/*font-bold*/
border-radius: .25rem;
/*rounded*/
background: #667eea !important;
/*bg-indigo-500*/
border: 1px solid transparent;
/*border border-transparent*/
}

/*Add padding to bottom border */
table.dataTable.no-footer {
border-bottom: 1px solid #e2e8f0;
/*border-b-1 border-gray-300*/
margin-top: 0.75em;
margin-bottom: 0.75em;
}

/*Change colour of responsive icon*/
table.dataTable.dtr-inline.collapsed>tbody>tr>td:first-child:before,
table.dataTable.dtr-inline.collapsed>tbody>tr>th:first-child:before {
background-color: #667eea !important;
/*bg-indigo-500*/
}
</style>  
