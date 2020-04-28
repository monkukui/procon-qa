import Vue from 'vue';
import Router from 'vue-router';
import Home from './views/Home.vue';
import Question from './views/Question.vue';
import Login from './views/Login.vue';
import SignUp from './views/SignUp.vue';
import QuestionForm from './views/QuestionForm.vue';
import UserPage from './views/UserPage.vue';
import TagSearch from './views/TagSearch.vue';
import About from './views/About.vue';
import Completed from './views/Completed.vue';
import UserNotification from './views/UserNotification.vue';

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
      path: '/userpage',
      name: 'userpage',
      component: UserPage,
    },
    {
      path: '/tagsearch',
      name: 'tagsearch',
      component: TagSearch,
    },
    {
      path: '/about',
      name: 'about',
      component: About,
    },
    {
      path: '/completed',
      name: 'completed',
      component: Completed,
    },
    {
      path: '/notification',
      name: 'notification',
      component: UserNotification,
    },
  ],
  scrollBehavior() {
    return {x: 0, y: 0};
  },
});
