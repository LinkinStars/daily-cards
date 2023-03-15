import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'

import BasicLayout from '../components/BasicLayout.vue'
import CardDetailPage from '../views/CardDetailPage.vue'
import CardsPage from '../views/CardsPage.vue'
import UserCardDetail from '../views/UserCardDetailPage.vue'
import UserCardsPage from '../views/UserCardsPage.vue'
import UserLoginPage from '../views/UserLoginPage.vue'
import UserPostCardPage from '../views/UserPostCardPage.vue'

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
      }
    ]
  }
  
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
