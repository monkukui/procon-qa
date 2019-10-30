import Vue from 'vue';
import Router from 'vue-router';
import Home from './views/Home.vue';
import Question from './views/Question.vue';
import Login from './views/Login.vue';
import SignUp from './views/SignUp.vue';
import QuestionForm from './views/QuestionForm.vue';
import MyPage from './views/MyPage.vue';
import TagSearch from './views/TagSearch.vue';

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home,
    },
    {
      path: '/question',
      name: 'question',
      component: Question,
    },
    {
      path: '/login',
      name: 'login',
      component: Login,
    },
    {
      path: '/signup',
      name: 'signup',
      component: SignUp,
    },
    {
      path: '/questionform',
      name: 'questionform',
      component: QuestionForm,
    },
    {
      path: '/mypage',
      name: 'mypage',
      component: MyPage,
    },
    {
      path: '/tagsearch',
      name: 'tagsearch',
      component: TagSearch,
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import(/* webpackChunkName: "about" */ './views/About.vue'),
    },
  ],
});
