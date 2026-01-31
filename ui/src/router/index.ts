import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'

import BasicLayout from '../components/BasicLayout.vue'
import CardDetailPage from '../views/CardDetailPage.vue'
import CardsPage from '../views/CardsPage.vue'
import UserLoginPage from '../views/UserLoginPage.vue'
import UserPostCardPage from '../views/UserPostCardPage.vue'
import NotFoundPage from '../views/NotFoundPage.vue'
import CardStatPage from '../views/CardStatPage.vue'

const routes: Array<RouteRecordRaw> = [
  // Login page - no layout wrapper for full screen
  {
    path: '/card/login',
    name: 'user-login',
    component: UserLoginPage
  },
  {
    path: '/',
    name: '',
    component: BasicLayout,
    children: [
      {
        path: '/card/:id',
        name: 'card-detail',
        component: CardDetailPage
      },
      {
        path: '/card/page',
        name: 'card-page',
        component: CardsPage
      },
      {
        path: '/card/stat',
        name: 'card-stat',
        component: CardStatPage
      },
      {
        path: '/card/post/:id',
        name: 'user-card-post',
        component: UserPostCardPage
      },
      {
        path: '/card/error/404',
        name: 'not-found',
        component: NotFoundPage
      }
    ]
  },
  {
    path: '/:catchAll(.*)',
    redirect: '/card/error/404'
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach(async (to) => {
  // Protected routes that require login
  if (to.name === 'user-card-post') {
    let token = localStorage.getItem('accessToken')
    if (!token) {
      return { name: 'user-login' }
    }
  }
})

export default router
