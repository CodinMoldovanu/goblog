import Vue from 'vue';
import VueRouter from 'vue-router';
import Home from './components/Home';
import Login from './components/Login';
import Post from './components/Post';
import Backend from './components/Backend';
import EditPost from './components/EditPost';
import NewPost from './components/NewPost';
import store from './store.js';

Vue.use(VueRouter);

const router = new VueRouter({
  mode: 'history',
  routes: [
    { path: '/login', component: Login, meta: { layout: 'default' } },
    { path: '/', component: Home, meta: { layout: 'default' } },
    { path: '/posts/:post_URL', name: 'post', component: Post, meta: { layout: 'default' } },
    { path: '/backend', component: Backend, meta: { layout: 'admin' } },
    { path: '/backend/post/:post_URL', name: 'editPost', component: EditPost, meta: { layout: 'admin' } },
    { path: '/backend/post/new', name: 'newPost', component: NewPost, meta: { layout: 'admin' } },
    { path: '/*', redirect: '/login', meta: { layout: 'default' } }

  ]
});

router.beforeEach((to, from, next) => {
  store.dispatch('fetchAccessToken');
  if (to.fullPath === '/backend') {
    if (!store.state.accessToken) {
      next('/login');
    }
  }
  if (to.fullPath === '/login') {
    if (store.state.accessToken) {
      next('/backend');
    }
  }
  next();
});
export default router