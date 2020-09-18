import Vue from 'vue';
import Vuetify from 'vuetify/lib';

Vue.use(Vuetify);

export default new Vuetify({
  icons: {
    iconfont: 'mdi',
  },
  theme: {
    dark: {
      background: '#030a27',
      primary: '#030a27',
      secondary: '#b0bec5',
      accent: '#8c9eff',
      error: '#b71c1c',
    },
    dark: true,
  },
});
