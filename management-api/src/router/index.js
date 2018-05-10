import Vue from 'vue'
import Router from 'vue-router'
import Login from '@/components/Login'
import DashBoard from '@/components/DashBoard'
import Management from '@/components/Management'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Login',
      component: Login
    },
    {
      path: '/management',
      name: 'Management',
      component: Management,
      children: [
        { path: '', name: 'dashboard', component: DashBoard }
      ]
    }
  ]
})
