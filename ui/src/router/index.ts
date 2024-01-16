import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'

import BasicLayout from '../components/BasicLayout.vue'
import CardDetailPage from '../views/CardDetailPage.vue'
import CardsPage from '../views/CardsPage.vue'
import UserCardDetail from '../views/UserCardDetailPage.vue'
import UserCardsPage from '../views/UserCardsPage.vue'
import UserLoginPage from '../views/UserLoginPage.vue'
import UserPostCardPage from '../views/UserPostCardPage.vue'
import NotFoundPage from '../views/NotFoundPage.vue'
import CardStatPage from '../views/CardStatPage.vue'

const routes: Array<RouteRecordRaw> = [
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
        path: '/card/uc/:id',
        name: 'user-card-detail',
        component: UserCardDetail
      },
      {
        path: '/card/uc/page',
        name: 'user-card-page',
        component: UserCardsPage
      },
      {
        path: '/card/uc/login',
        name: 'user-login',
        component: UserLoginPage
      },
      {
        path: '/card/uc/post/:id',
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

router.beforeEach(async (to, from) => {
  if (to.name === 'user-card-detail' || to.name === 'user-card-page' || to.name === 'user-card-post') {
    let token = localStorage.getItem('accessToken')
    if (!token) {
     return { name: 'user-login' }
    }
  }
 })

export default router
