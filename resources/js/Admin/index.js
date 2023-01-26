import Vuex from 'vuex';

export default new Vuex.Store({
  state: {
    sideBarOpen: false,
  },
  getters: {
    sideBarOpen: state => {
      return state.sideBarOpen;
    },
  },
  mutations: {
    toggleSidebar (state) {
      state.sideBarOpen = !state.sideBarOpen;
    },
  },
  actions: {
    toggleSidebar(context) {
      context.commit('toggleSidebar');
    },
  },
});
